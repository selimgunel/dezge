package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/narslan/dezge/http"
	"github.com/narslan/dezge/inmem"
	"github.com/narslan/dezge/sqlite"
	"github.com/pelletier/go-toml"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Long: "dezge server is a messaging client",
		Run: func(cmd *cobra.Command, args []string) {

			config, err := ReadConfigFile(cfgFile)
			if os.IsNotExist(err) {
				log.Fatalf("config file not found: %s", cfgFile)
			}
			//log.Debug().Msgf("Config: %v", config)

			m := NewMain(config)

			// Setup signal handlers.
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			ctx, cancel := context.WithCancel(context.Background())

			ctx.Deadline()

			go func() {
				<-c
				//	log.Info().Msgf("system call:%+v", oscall)
				cancel()
			}()

			if err := m.Run(ctx); err != nil {
				log.Fatal(err.Error())
			}
		},
	}
)

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func NewMain(config Config) *Main {

	return &Main{
		HTTPServer: http.NewServer(config.HTTP.Addr, stats),
		Config:     config,
	}
}

// Main represents the program.
type Main struct {
	Config Config

	HTTPServer *http.Server
}

// DefaultConfig returns a new instance of Config with defaults set.
func DefaultConfig() Config {
	var config Config
	return config
}

// ReadConfigFile unmarshals config from config file.
func ReadConfigFile(filename string) (Config, error) {
	config := DefaultConfig()
	buf, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	if err := toml.Unmarshal(buf, &config); err != nil {
		return config, err
	}
	return config, nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	cobra.OnInitialize(checkConfigFile)
}

func checkConfigFile() {
	if cfgFile == "" {
		log.Fatal("no config supplied")
	}

}

// Config represents the server configuration file.
type Config struct {
	DB struct {
		DSN string `toml:"dsn"`
	} `toml:"db"`

	HTTP struct {
		Addr   string `toml:"addr"`
		Domain string `toml:"domain"`
		Cert   string `toml:"cert"`
		Key    string `toml:"key"`
	} `toml:"http"`
}

func (m *Main) Run(ctx context.Context) error {

	m.HTTPServer.Domain = m.Config.HTTP.Domain
	m.HTTPServer.Cert = m.Config.HTTP.Cert
	m.HTTPServer.Key = m.Config.HTTP.Key

	//log.Debug().Msgf("%+v", m.Config)
	db := sqlite.NewDB(m.Config.DB.DSN)
	err := db.Open()
	if err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}
	es := sqlite.NewEngineInfo(db)
	m.HTTPServer.EngineInfoService = es

	ev := inmem.NewEventService()
	m.HTTPServer.EventService = ev

	if err := m.HTTPServer.Open(ctx); err != nil {
		return err
	}

	return nil
}
