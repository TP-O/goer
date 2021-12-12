package main

import (
	"strings"
	"testing"

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
