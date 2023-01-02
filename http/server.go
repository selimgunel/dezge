package http

import (
	"context"
	"expvar"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/narslan/dezge"
	"github.com/rs/zerolog/log"
)

// Server represents an HTTP server.
type Server struct {
	server *http.Server

	stat              *expvar.Map
	EventService      dezge.EventService
	GeneratorService  dezge.GeneratorService
	EngineInfoService dezge.EngineInfoService

	Addr   string
	Domain string
	Cert   string
	Key    string
}

// NewServer returns a new instance of Server.
func NewServer(addr string, stat *expvar.Map) *Server {

	gin.SetMode(gin.ReleaseMode)
	s := &Server{
		server: &http.Server{
			Addr: addr,
		},
		stat: stat,
	}

	router := gin.Default()
	router.Use(initCors())
	//router.Use(otelgin.Middleware("goGinApp"))

	router.GET("/ping", func(c *gin.Context) {
		//	_, span := tracer.Start(c.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", "yok")))
		//	defer span.End()
		c.String(200, "pong")
	})

	router.GET("/time", s.timeStream)
	router.GET("/doc", s.documentStream)

	s.server.Handler = router

	//router.LoadHTMLFiles("index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/echo", func(c *gin.Context) {
		s.handleEcho(c)

	})

	router.GET("/event", func(c *gin.Context) {
		s.handleEvent(c)
	})

	router.GET("/eventru", func(c *gin.Context) {
		s.handleEventRandomUser(c)
	})

	router.POST("/publish", func(c *gin.Context) {
		s.handlePublishEvent(c)
	})

	router.GET("/debug/expvar", func(c *gin.Context) {
		s.handleDebug(c)
	})

	router.GET("/debug/expvarws", func(c *gin.Context) {
		s.handleDebugWS(c)
	})

	router.GET("/chess/board/subscribe", func(c *gin.Context) {
		s.subscribePositionSVG(c)
	})

	router.POST("/chess/board/publish", func(c *gin.Context) {
		s.publishPositionSVG(c)
	})

	router.GET("/engines", func(c *gin.Context) {
		s.ListEngines(c)
	})

	return s
}

func (s *Server) Open(ctx context.Context) (err error) {

	log.Debug().Msgf("addr: %+v", s.server.Addr)
	go func() {
		if err = s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen:%+s\n", err)
		}
	}()

	log.Info().Msg("server started")
	<-ctx.Done()

	log.Info().Msg("server stoped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = s.server.Shutdown(ctxShutDown); err != nil {
		log.Fatal().Err(err).Msg("server Shutdown Failed:")
	}

	log.Info().Msg("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return

}
func initCors() gin.HandlerFunc {

	config := cors.Config{

		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods"},
	}
	err := config.Validate()
	if err != nil {
		log.Fatal().Msg("")
	}
	return cors.New(config)
}
