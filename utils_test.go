package main

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateMultipartFormPayload(t *testing.T) {
	fields := []MultipartFormField{
		{
			Name:  "field01",
			Value: "value01",
		},
	}
	payload := CreateMultipartFormPayload(fields)

	assert.True(t, strings.HasPrefix(payload.Type, "multipart/form-data; boundary="))
	assert.Contains(t, payload.Body.String(), "Content-Disposition: form-data; name=\""+fields[0].Name+"\"")
	assert.Contains(t, payload.Body.String(), fields[0].Value)
}

func TestNewHttp(t *testing.T) {
	http := NewHttp()

	assert.NotEmpty(t, http)
	assert.NotEmpty(t, http.GetClient().Jar)
	assert.Equal(t, http.GetClient().Timeout, 50*time.Second)
}
