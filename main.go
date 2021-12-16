package main

import (
	"os"
	"strings"
	"time"

	"github.com/withmandala/go-log"
)

func main() {
	// Receive variables from CLI
	id, password, host, session, courseId := RunCLI()
	registeredId := []string{}
	saved := false

	// Init logger
	logger := log.New(os.Stderr)

	// Client
	client := Client{
		Host:    host,
		Session: session,
		Http:    NewHttp(),
		PayloadGenerator: &PayloadGenerator{
			credentials: Credentials{
				ID:       id,
				Password: password,
			},
		},
	}

	for true {
		if ok, message := client.Login(); ok {
			logger.Info(message)

			break
		} else {
			logger.Warn(message)
		}
	}

	// Get student ID
	// logger.Info(client.SayHi())

	for !saved || len(registeredId) < len(courseId) {
		registeredString := strings.Join(registeredId, "")

		// Register for courses
		for _, id := range courseId {
			// Skip if course is already selected
			if strings.Contains(registeredString, id) {
				continue
			}

			course := strings.Split(id, "|")[2]

			if ok, messsage := client.Register(id); ok {
				// Update registerId list
				registeredId = append(registeredId, id)

				logger.Infof("[%s] %s", course, messsage)
			} else {
				logger.Warnf("[%s] %s", course, messsage)
			}

			// Avoid constant request sending
			time.Sleep(2 * time.Second)
		}

		// Save registration if new course is selected
		if !saved || registeredString != strings.Join(registeredId, "") {
			if ok, messsage := client.Save(); ok {
				saved = ok

				logger.Infof(messsage)
			} else {
				logger.Warnf(messsage)
			}
		}

		// Avoid constant request sending
		time.Sleep(2 * time.Second)
	}
}
