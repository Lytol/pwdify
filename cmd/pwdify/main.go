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
		ArgsUsage:       "[files... | directory]",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "directory",
				Aliases: []string{"d"},
				Value:   ".",
				Usage:   "working directory",
			},
			&cli.StringFlag{
				Name:  "password-env",
				Value: DefaultPasswordEnv,
				Usage: "environment variable to read the password from",
			},
		},
		Action: func(ctx *cli.Context) error {
			state := &state{
				password: "",
				files:    []string{},
				cwd:      ".",
			}

			// If provided, set the working directory
			directory := ctx.String("directory")
			di, err := os.Stat(directory)
			if err != nil {
				return fmt.Errorf("could not find directory: %s", err)
			}
			if !di.IsDir() {
				return fmt.Errorf("not a directory: %s", directory)
			}

			// Set password from environment variable
			state.password = os.Getenv(ctx.String("password-env"))

			// Enumerate files and/or directories from the arguments
			for _, arg := range ctx.Args().Slice() {
				fileOrDir := filepath.Join(directory, arg)
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
