package main

import (
	"context"
	"fmt"
	"os"

	"github.com/narslan/dezge/cmd/dezgectl/commands/createcmd"
	"github.com/narslan/dezge/cmd/dezgectl/commands/rootcmd"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {

	out := os.Stdout
	rootCommand, rootConfig := rootcmd.New()
	createCommand := createcmd.New(rootConfig, out)

	rootCommand.Subcommands = []*ffcli.Command{
		createCommand,
	}

	if err := rootCommand.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error during Parse: %v\n", err)
		os.Exit(1)
	}

	if err := rootCommand.Run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
