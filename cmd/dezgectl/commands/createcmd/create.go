package createcmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/narslan/dezge"
	"github.com/narslan/dezge/cmd/dezgectl/commands/rootcmd"
	"github.com/narslan/dezge/sqlite"
	"github.com/pelletier/go-toml"
	"github.com/peterbourgon/ff/v3/ffcli"
)

// Config for the create subcommand, including a reference to the API client.
type FFConfig struct {
	rootConfig *rootcmd.Config
	out        io.Writer
	overwrite  bool
}

// New returns a usable ffcli.Command for the create subcommand.
func New(rootConfig *rootcmd.Config, out io.Writer) *ffcli.Command {
	cfg := FFConfig{
		rootConfig: rootConfig,
		out:        out,
	}

	fs := flag.NewFlagSet("objectctl create", flag.ExitOnError)
	fs.BoolVar(&cfg.overwrite, "overwrite", false, "overwrite existing object, if it exists")
	rootConfig.RegisterFlags(fs)

	return &ffcli.Command{
		Name:       "create",
		ShortUsage: "ssectl create [flags] /path/to/engine>",
		ShortHelp:  "Create an engine info on db",
		FlagSet:    fs,
		Exec:       cfg.Exec,
	}
}

// Exec function for this command.
func (c *FFConfig) Exec(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("create requires at least 1 args")
	}
	if c.rootConfig.File == "" {
		return errors.New("supply a conf file")
	}
	for _, v := range args {
		fmt.Println("args", v)
	}
	cfg, err := ReadConfigFile(c.rootConfig.File)
	if err != nil {
		return err
	}
	m := NewMain(cfg)

	db := sqlite.NewDB(m.Config.DB.DSN)
	err = db.Open()
	if err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}
	es := sqlite.NewEngineInfo(db)

	engineInfo := dezge.NewEngineInfo(args[0])

	return es.Create(ctx, engineInfo)
}

func NewMain(config Config) *Main {

	return &Main{
		Config: config,
	}
}

// Main represents the program.
type Main struct {
	Config Config
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

// ReadConfigFile unmarshals config from config file.
func ReadConfigFile(filename string) (Config, error) {
	config := Config{}
	if buf, err := ioutil.ReadFile(filename); err != nil {
		return config, err
	} else if err := toml.Unmarshal(buf, &config); err != nil {
		return config, err
	}
	return config, nil
}
