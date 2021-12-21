package main

import (
	"bufio"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/withmandala/go-log"
)

func main() {
	// Receive variables from CLI
	id, password, host, session, careful, courseId := RunCLI()
	registeredId := []string{}

	// Careful mode will ignore this saved condition
	// in the while loop below (always true).
	saved := careful

	// Read input from console
	reader := bufio.NewReader(os.Stdin)
	input := ""

	// Init logger
	logger := log.New(os.Stderr)

	// Disble colorf on windows
	if runtime.GOOS == "windows" {
		logger.WithoutColor()
	}

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

	// Retry until successful login
	for true {
		if ok, message := client.Login(); ok {
			logger.Info(message)

			break
		} else {
			logger.Warn(message)
		}
	}

	// Wait for user accepting
	for len(input) == 0 {
		logger.Info("Press enter to continue")
		input, _ = reader.ReadString('\n')
	}

	// Get student ID
	// logger.Info(client.SayHi())

	// If careful mode is enabled, `!saved`` condition is always false.
	// Therefore, the condition to continue this while loop is
	// to check the lenghth of registered course array.
	for !saved || len(registeredId) < len(courseId) {
		registeredString := strings.Join(registeredId, "")

		// Register for courses
		for _, id := range courseId {
			// Skip if course is already selected
			if strings.Contains(registeredString, id) {
				continue
			}

			// Get course name
			course := strings.Split(id, "|")[2]

			if ok, messsage := client.Register(id); ok {
				// Update registerId list
				registeredId = append(registeredId, id)

				logger.Infof("[%s] %s", course, messsage)

				// Save after selecting
				if careful {
					if ok, messsage := client.Save(); ok {
						logger.Infof(messsage)
					} else {
						logger.Warnf(messsage)
					}
				}
			} else {
				logger.Warnf("[%s] %s", course, messsage)
			}

			// Avoid constant request sending
			time.Sleep(2 * time.Second)
		}

		// Save registration if new course is selected.
		// Ignore if careful mode is enabled.
		if !careful && (!saved || registeredString != strings.Join(registeredId, "")) {
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
