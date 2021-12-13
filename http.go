package main

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

type HttpInterface interface {
	GetClient() *http.Client
	Do(req *http.Request) (*http.Response, error)
}

type Http struct {
	HttpClient *http.Client
}

func (h *Http) Initialize() {
	if jar, err := cookiejar.New(nil); err == nil {
		h.HttpClient = &http.Client{
			Jar:     jar,
			Timeout: 50 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}
}

func (h *Http) GetClient() *http.Client {
	return h.HttpClient
}

func (h Http) Do(req *http.Request) (*http.Response, error) {
	return h.HttpClient.Do(req)
}
