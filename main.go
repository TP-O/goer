package main

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)
}

func main() {
	options := RunCLI()
	credentials := &Credentials{
		ID:       options.ID,
		Password: options.Password,
	}
	goer := NewGoer(options.Origin, NewGoerClient())

	logrus.Warn("=======================================================")
	logrus.Warn("DO NOT ACCESS YOUR ACCOUNT WHEN THIS TOOL IS RUNNING!!!")
	logrus.Warn("=======================================================")

	// Try to log in again until the registration is ready
	for true {
		if ok := goer.Login(credentials); ok {
			// goer.Greet()

			if isOpen := goer.IsRegistrationOpen(); isOpen {
				break
			} else {
				goer.Clear()
			}
		}

		time.Sleep(1 * time.Second)
		logrus.Info("Login again...")
	}

	// Start registration
	var wg sync.WaitGroup
	var mu sync.Mutex
	courseCounter := 0
	courseIDChannel := make(chan string, options.Workers)

	go func() {
		for _, courseID := range options.CourseIDs {
			courseIDChannel <- courseID
		}
	}()

	for i := 0; i < int(options.Workers); i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			for courseID := range courseIDChannel {
				if ok := goer.RegisterCourse(courseID); !ok {
					courseIDChannel <- courseID
				} else {
					// Save the enrolled course after successful registration
					if options.CarefulMode {
						goer.SaveRegistration()
					}

					// Put the course ID back if spam is enable
					if options.SpamInterval != 0 {
						courseIDChannel <- courseID
						time.Sleep(time.Duration(options.SpamInterval) * time.Second)
					} else {
						mu.Lock()

						courseCounter++

						// Close the channel if all courses are registered
						if courseCounter == len(options.CourseIDs) {
							close(courseIDChannel)
						}

						mu.Unlock()
					}
				}
			}
		}()
	}

	wg.Wait()

	// Save all registered courses
	if !options.CarefulMode {
		goer.SaveRegistration()
	}

	logrus.Info("Done!")
}
