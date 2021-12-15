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

	for client.Login() == false {
	}

	client.SayHi()
	client.Register("CH012IU06    |CH012IU|Chemistry Laboratory|06|1|0|01/01/0001|0|0|0||0|ITIT19CS31")
	client.Save()
}
