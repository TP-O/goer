package main

import (
	"os"

	"github.com/withmandala/go-log"
)

var logger = log.New(os.Stderr)

func main() {
	client := Client{
		Host: "https://edusoftweb.hcmiu.edu.vn",
		Http: NewHttp(),
		PayloadGenerator: &PayloadGenerator{
			credentials: Credentials{
				ID:       "ITITIU19180",
				Password: "Password",
			},
		},
	}

	client.Login()
	client.SayHi()
}
