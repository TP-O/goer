package main

import (
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"github.com/withmandala/go-log"
)

var logger = log.New(os.Stderr)

func main() {
	var client Client

	if jar, err := cookiejar.New(nil); err == nil {
		client = Client{
			Host: "https://edusoftweb.hcmiu.edu.vn",
			PayloadGenerator: &PayloadGenerator{
				credentials: Credentials{
					ID:       "ITITIU19180",
					Password: "Password",
				},
			},
			HttpClient: &http.Client{
				Jar:     jar,
				Timeout: 50 * time.Second,
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			},
		}
	} else {
		panic(err)
	}

	client.Login()
	client.SayHi()
}
