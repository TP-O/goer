package main

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PayloadGeneratorMock struct {
	mock.Mock
}

func (p *PayloadGeneratorMock) LoginPayload() Payload {
	args := p.Called()

	return args.Get(0).(Payload)
}

type HttpMock struct {
	mock.Mock
}

func (h *HttpMock) GetClient() *http.Client {
	return nil
}

func (h *HttpMock) Do(req *http.Request) (*http.Response, error) {
	args := h.Called(req)

	return args.Get(0).(*http.Response), args.Error(1)
}

func TestLogin(t *testing.T) {
	payloadGeneratorMock := PayloadGeneratorMock{}
	httpMock := HttpMock{}

	payload := Payload{
		Type: "cc",
		Body: &bytes.Buffer{},
	}

	payloadGeneratorMock.On("LoginPayload").Return(payload)

	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{StatusCode: 302}, nil).Once()
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{StatusCode: 200}, nil).Once()
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{}, errors.New("Error")).Once()

	client := Client{
		Host:             "mock.com",
		Http:             &httpMock,
		PayloadGenerator: &payloadGeneratorMock,
	}

	assert.NotEmpty(t, client)
	assert.True(t, client.Login())
	assert.False(t, client.Login())
	assert.False(t, client.Login())
}
