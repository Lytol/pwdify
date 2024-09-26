package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	"github.com/lytol/pwdify/internal/util"
)

const DefaultPasswordEnv = "PWDIFY_PASSWORD"

var (
	// version should be provided by ldflags in release
	version = "dev"

	logger util.Logger
)

func main() {
	var err error

	logger, err = util.NewLogger()
	if err != nil {
		fmt.Printf("could not start logger: %s\n", err)
		os.Exit(1)
	}

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
			state := &state{
				password: "",
				files:    []string{},
				cwd:      ".",
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
			state.cwd = cwd

			// Set password from environment variable
			state.password = os.Getenv(ctx.String("password-env"))

			// Enumerate files and/or directories from the arguments
			for _, arg := range ctx.Args().Slice() {
				fileOrDir := filepath.Join(state.cwd, arg)
				_, err := os.Stat(fileOrDir)
				if err != nil {
					return fmt.Errorf("could not find file or directory: %s", err)
				}
				state.files = append(state.files, fileOrDir)
			}

			logger.Logf("state: %+v\n", state)

			_, err = tea.NewProgram(newModel(state)).Run()
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
