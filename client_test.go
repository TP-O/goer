package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGoerClient(t *testing.T) {
	client := NewGoerClient()

	assert.NotNil(t, client)
	assert.NotNil(t, client.HttpClient.Jar)
}
