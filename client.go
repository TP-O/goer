package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	Host             string
	HttpClient       *http.Client
	PayloadGenerator *PayloadGenerator
}

func (c Client) Login() bool {
	path := "/default.aspx"
	payload := c.PayloadGenerator.LoginPayload()

	req, _ := http.NewRequest("POST", c.Host+path, payload.Body)
	req.Header.Add("Content-Type", payload.Type)

	if res, err := c.HttpClient.Do(req); err != nil {
		log.Println(err)
		log.Println("Trying to login again")

		c.Login()

		return false
	} else if res.StatusCode != 302 {
		log.Println("Login failed!!!")
		log.Println("Trying to login again")

		c.Login()

		return false
	}

	log.Println("Login successfully!!!")

	return true
}

func (c Client) SayHi() {
	path := "/default.aspx"

	req, _ := http.NewRequest("GET", c.Host+path, nil)

	if res, err := c.HttpClient.Do(req); err != nil {
		log.Println(err)
	} else {
		document, _ := goquery.NewDocumentFromReader(res.Body)
		fmt.Println(document.Find("#ctl00_Header1_Logout1_lblNguoiDung").Text())
	}
}
