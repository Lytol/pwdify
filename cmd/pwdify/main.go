package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	"github.com/lytol/pwdify/internal/util"
)

const DefaultPasswordEnv = "PWDIFY_PASSWORD"

var (
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
		Usage:           "Password protect your static HTML",
		ArgsUsage:       "[directory]",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "password-env",
				Value: DefaultPasswordEnv,
				Usage: "The environment variable to read the password from",
			},
		},
		Action: func(ctx *cli.Context) error {
			state := &state{
				password: "",
				files:    []string{},
				cwd:      ".",
			}

			if ctx.NArg() > 1 {
				return fmt.Errorf("too many arguments")
			}

			// If provided, set the working directory
			directory := ctx.Args().First()
			if directory != "" {
				// Ensure that the directory exists
				if _, err := os.Stat(directory); os.IsNotExist(err) {
					return fmt.Errorf("`%s` is not a valid directory", directory)
				}
				state.cwd = directory
			}

			// Set password from environment variable
			state.password = os.Getenv(ctx.String("password-env"))

			logger.Logf("state: %+v\n", state)

			_, err = tea.NewProgram(newModel(state)).Run()
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
