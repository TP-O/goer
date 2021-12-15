package main

import (
	"os"
	"strings"

	"github.com/withmandala/go-log"
)

func main() {
	// Receive variables from CLI
	id, password, courseId := RunCLI()

	// Init logger
	logger := log.New(os.Stderr)

	// Client
	client := Client{
		Host: "https://edusoftweb.hcmiu.edu.vn",
		Http: NewHttp(),
		PayloadGenerator: &PayloadGenerator{
			credentials: Credentials{
				ID:       id,
				Password: password,
			},
		},
	}

	for true {
		loggedIn, message := client.Login()

		if loggedIn == true {
			logger.Info(message)

			break
		} else {
			logger.Warn(message)
		}
	}

	// Get student ID
	// logger.Info(client.SayHi())

	for len(courseId) != 0 {
		// Register for courses
		for i, id := range courseId {
			if ok, err := client.Register(id); ok {
				// Append next part to previous part
				courseId = append(courseId[:i], courseId[i+1:]...)

				logger.Infof("Registered: %s", strings.Split(id, "|")[2])
			} else {
				logger.Warnf(err.Error())
			}
		}

		// Save registration
		client.Save()
	}
}
