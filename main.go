package main

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/withmandala/go-log"
)

func main() {
	// Receive variables from CLI.
	id, password, host, careful, spam, courseIds := RunCLI()
	registeredIds := []string{}

	// Determine if saving registrations is required or not.
	saved := true

	// Init logger.
	logger := log.New(os.Stderr)

	// Disble colorf on windows.
	if runtime.GOOS == "windows" {
		logger.WithoutColor()
	}

	// Declare client with student credentials.
	client := Client{
		Host: host,
		Http: NewHttp(),
		PayloadGenerator: &PayloadGenerator{
			credentials: Credentials{
				ID:       id,
				Password: password,
			},
		},
	}

	// Retry loggin in until registration is ready.
	for true {
		if ok, message := client.Login(); ok {
			logger.Info(message)

			if isReady, isReadyMessage := client.IsReady(); isReady {
				logger.Info(isReadyMessage)

				break
			} else {
				logger.Warn(isReadyMessage)
				logger.Warn("Logging in again...")

				client.Reset()
			}
		} else {
			logger.Warn(message)
		}
	}

	// Get student ID.
	// logger.Info(client.SayHi())

	if spam {
		for true {
			// Register for courses
			for _, courseId := range courseIds {
				// Get course name
				course := strings.Split(courseId, "|")[2]

				if ok, messsage := client.Register(courseId); ok {
					logger.Infof("[%s] %s", course, messsage)

					_, messsage := client.Save()
					logger.Infof(messsage)
				} else {
					logger.Warnf("[%s] %s", course, messsage)
				}

				// Avoid constant request sending
				time.Sleep(5 * time.Second)
			}
		}
	} else {
		// Checking saving registrations is ignored if careful mode is enable
		// because saving is executed after select one course. Run until
		// all courses are registered.
		for (!saved && !careful) || len(registeredIds) < len(courseIds) {
			registeredString := strings.Join(registeredIds, "")

			// Register for courses
			for _, courseId := range courseIds {
				// Skip if course is already selected
				if strings.Contains(registeredString, courseId) {
					continue
				}

				// Get course name
				course := strings.Split(courseId, "|")[2]

				if ok, messsage := client.Register(courseId); ok {
					// Update registerIds list
					registeredIds = append(registeredIds, courseId)

					logger.Infof("[%s] %s", course, messsage)

					// Save after selecting
					if careful {
						if ok, messsage := client.Save(); ok {
							logger.Infof(messsage)
						} else {
							registeredIds = registeredIds[:len(registeredIds)-1]

							logger.Warnf(messsage)
						}
					}
				} else {
					logger.Warnf("[%s] %s", course, messsage)
				}

				// Avoid constant request sending
				time.Sleep(2 * time.Second)
			}

			// Save registration if new courses are selected.
			// Ignore if careful mode is enabled.
			if !careful && (!saved || registeredString != strings.Join(registeredIds, "")) {
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
}
