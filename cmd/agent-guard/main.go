// Command agent-guard is the generic-purpose cli-guard consumer entry point.
//
// Wraps a small fixed surface of dev verbs (build, test, vet, lint, tidy)
// behind the cli-guard policy gate, for repos that take external contributions.
// See README.md for the audience and scope rationale.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

// Version is set at build time via -ldflags.
var Version = "dev"

func main() {
	app := &cli.Command{
		Name:    "agent-guard",
		Usage:   "generic cli-guard consumer for external-contributor repos",
		Version: Version,
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "print the build version and exit",
				Action: func(_ context.Context, _ *cli.Command) error {
					fmt.Println(Version)
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "agent-guard:", err)
		os.Exit(1)
	}
}
