package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type Options struct {
	ID           string
	Password     string
	Origin       string
	Workers      uint64
	CarefulMode  bool
	SpamInterval uint64
	CourseIDs    []string
}

func RunCLI() *Options {
	shouldExit := true
	options := &Options{}
	app := &cli.App{
		Name:    "goer",
		Usage:   "A simple tool to help students enroll in their courses on the Edusoft website",
		Version: "2.0.2",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				Aliases:     []string{"i"},
				Usage:       "Login to account with the provided `ID`",
				Required:    true,
				Destination: &options.ID,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Usage:       "Account's password of the provied ID",
				Required:    true,
				Destination: &options.Password,
			},
			&cli.StringFlag{
				Name:        "origin",
				Aliases:     []string{"o"},
				Usage:       "Origin",
				Value:       "https://edusoftweb.hcmiu.edu.vn",
				Destination: &options.Origin,
			},
			&cli.BoolFlag{
				Name:        "careful",
				Aliases:     []string{"c"},
				Usage:       "Save after each single successful registration",
				Value:       true,
				Destination: &options.CarefulMode,
			},
			&cli.Uint64Flag{
				Name:        "spam",
				Aliases:     []string{"s"},
				Usage:       "Repeat the registrations every `TIME` seconds",
				Destination: &options.SpamInterval,
				Action: func(ctx *cli.Context, u uint64) error {
					if u < 1 {
						logrus.Fatalf("Flag `spam` value %v must be equal to or greater than 1", u)
					}

					return nil
				},
			},
			&cli.Uint64Flag{
				Name:        "workers",
				Aliases:     []string{"w"},
				Usage:       "Set the number of workers which register for courses",
				Value:       1,
				Destination: &options.Workers,
				Action: func(ctx *cli.Context, u uint64) error {
					if u < 1 {
						logrus.Fatalf("Flag `workers` value %v must be equal to or greater than 1", u)
					}

					return nil
				},
			},
			&cli.StringSliceFlag{
				Name:     "course-id",
				Aliases:  []string{"I"},
				Required: true,
				Usage:    "ID of registered course",
			},
		},
		Action: func(ctx *cli.Context) error {
			options.CourseIDs = ctx.StringSlice("course-id")
			shouldExit = false

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

	if shouldExit {
		os.Exit(0)
	}

	return options
}
