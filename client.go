package main

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type IGoerClient interface {
	Do(req *http.Request) (*http.Response, error)
	DeleteSessionId(domain string) error
}

type GoerClient struct {
	HttpClient *http.Client
}

func NewGoerClient() *GoerClient {
	if jar, err := cookiejar.New(nil); err != nil {
		panic(SystemFailureMessage + err.Error())
	} else {
		return &GoerClient{
			&http.Client{
				Jar:     jar,
				Timeout: 60 * time.Second,
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			},
		}
	}
}

func (c *GoerClient) Do(req *http.Request) (*http.Response, error) {
	return c.HttpClient.Do(req)
}

func (c *GoerClient) DeleteSessionId(domain string) error {
	if url, err := url.Parse(domain); err != nil {
		return err
	} else {
		c.HttpClient.Jar.SetCookies(url, []*http.Cookie{
			{
				Name:     SessionIDCookieField,
				Value:    "",
				HttpOnly: true,
			},
		})

		return nil
	}
}
