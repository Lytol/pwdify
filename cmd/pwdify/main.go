package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"

	"github.com/lytol/pwdify/cmd/pwdify/tui"
	"github.com/lytol/pwdify/pkg/pwdify"
)

const DefaultPasswordEnv = "PWDIFY_PASSWORD"

var (
	// version should be provided by ldflags in release
	version = "dev"
)

func main() {
	app := &cli.App{
		Name:            "pwdify",
		Usage:           "Password protect static web pages",
		ArgsUsage:       "[file | directory ...]",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "cwd",
				Aliases: []string{"d"},
				Value:   ".",
				Usage:   "working directory",
			},
			&cli.StringFlag{
				Name:  "password-env",
				Value: DefaultPasswordEnv,
				Usage: "environment variable for password",
			},
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show version",
			},
		},
		Action: func(ctx *cli.Context) error {
			cfg := &pwdify.Config{
				Password: "",
				Files:    []string{},
				Cwd:      ".",
			}

			// If version flag is set, print version and exit
			if ctx.Bool("version") {
				fmt.Println(version)
				return nil
			}

			// If provided, set the working directory
			cwd := ctx.String("cwd")
			di, err := os.Stat(cwd)
			if err != nil {
				return fmt.Errorf("could not find directory: %s", err)
			}
			if !di.IsDir() {
				return fmt.Errorf("not a directory: %s", cwd)
			}
			cfg.Cwd = cwd

			// Set password from environment variable
			cfg.Password = os.Getenv(ctx.String("password-env"))

			// Enumerate files and/or directories from the arguments
			for _, arg := range ctx.Args().Slice() {
				fileOrDir := filepath.Join(cfg.Cwd, arg)
				_, err := os.Stat(fileOrDir)
				if err != nil {
					return fmt.Errorf("could not find file or directory: %s", err)
				}
				cfg.Files = append(cfg.Files, fileOrDir)
			}

			if isatty.IsTerminal(os.Stdout.Fd()) {
				return tui.Run(cfg)
			} else {
				if cfg.Password == "" {
					return fmt.Errorf("you must specify a password")
				}

				if len(cfg.Files) == 0 {
					return fmt.Errorf("you must specify one or more files or directories")
				}

				status, total, err := pwdify.Encrypt(cfg.Files, cfg.Password)
				if err != nil {
					log.Fatal(err)
				}

				completed := 0
				errs := []error{}

				for {
					s, ok := <-status
					if !ok {
						return errors.Join(errs...)
					}
					completed++
					fmt.Fprintf(os.Stderr, "%d/%d | %s", completed, total, s.File)
					if s.Error != nil {
						errs = append(errs, s.Error)
						fmt.Fprintf(os.Stderr, " | %s\n", s.Error)
					} else {
						fmt.Fprintf(os.Stderr, "\n")
					}
				}
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
