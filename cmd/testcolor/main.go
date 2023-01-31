package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(); err != nil {
		log.Printf("run: %s", err.Error())
	}
}

func run() error {
	if err := app().Run(os.Args); err != nil {
		return fmt.Errorf("run: %w", err)
	}
	return nil
}

func app() *cli.App {
	return &cli.App{
		Action: func(ctx *cli.Context) error {
			scanner := bufio.NewScanner(ctx.App.Reader)
			w := ctx.App.Writer

			for scanner.Scan() {
				line := scanner.Text()

				switch {
				default:
					fmt.Fprintln(w, line)
				case regexp.MustCompile(`^\s*--- FAIL: `).MatchString(line):
					color.New(color.FgRed).Fprintln(w, line)
				}
			}

			if err := scanner.Err(); err != nil {
				return fmt.Errorf("scan: %w", err)
			}

			return nil
		},
	}
}
