package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
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

func (p *PayloadGeneratorMock) RegistrationPayload(id string) (Payload, string) {
	args := p.Called(id)

	return args.Get(0).(Payload), args.String(1)
}

func (p *PayloadGeneratorMock) SavePayload() Payload {
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
		Type: "xx",
		Body: &bytes.Buffer{},
	}

	payloadGeneratorMock.On("LoginPayload").Return(payload)

	client := Client{
		Http:             &httpMock,
		PayloadGenerator: &payloadGeneratorMock,
	}

	assert.NotEmpty(t, client)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
		StatusCode: 302,
		Header: map[string][]string{
			"Location": {""},
		},
	}, nil).Once()

	ok, _ := client.Login()

	assert.True(t, ok)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
		StatusCode: 302,
		Header: map[string][]string{
			"Location": {"sessionreuse"},
		},
	}, nil).Once()

	ok, _ = client.Login()

	assert.False(t, ok)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
		StatusCode: 200,
		Header: map[string][]string{
			"Location": {},
		},
	}, nil).Once()

	ok, _ = client.Login()

	assert.False(t, ok)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{}, errors.New("Error")).Once()

	ok, _ = client.Login()

	assert.False(t, ok)
}

func TestRegister(t *testing.T) {
	ID := "MaDK|MaMH|TenMH|MaNh|Sotc||StrngayThi||TietBD|SoTiet|"
	payloadGeneratorMock := PayloadGeneratorMock{}
	httpMock := HttpMock{}

	payload := Payload{
		Type: "xx",
		Body: &bytes.Buffer{},
	}

	payloadGeneratorMock.On("RegistrationPayload", ID).Return(payload, "TenMH")

	client := Client{
		Http:             &httpMock,
		PayloadGenerator: &payloadGeneratorMock,
	}

	assert.NotEmpty(t, client)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
		Body: io.NopCloser(strings.NewReader("TenMH")),
	}, nil).Once()

	ok, message := client.Register(ID)

	assert.True(t, ok)
	assert.NotEmpty(t, message)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
		Body: io.NopCloser(strings.NewReader("")),
	}, nil).Once()

	ok, message = client.Register(ID)

	assert.False(t, ok)
	assert.NotEmpty(t, message)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{}, errors.New("Error")).Once()

	ok, message = client.Register(ID)

	assert.False(t, ok)
	assert.NotEmpty(t, message)
}

func TestSave(t *testing.T) {
	payloadGeneratorMock := PayloadGeneratorMock{}
	httpMock := HttpMock{}

	payload := Payload{
		Type: "xx",
		Body: &bytes.Buffer{},
	}

	payloadGeneratorMock.On("SavePayload").Return(payload)

	client := Client{
		Http:             &httpMock,
		PayloadGenerator: &payloadGeneratorMock,
	}

	assert.NotEmpty(t, client)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
		Body: io.NopCloser(strings.NewReader("||default.aspx?page=dkmonhoc")),
	}, nil).Once()

	ok, message := client.Save()

	assert.True(t, ok)
	assert.NotEmpty(t, message)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
		Body: io.NopCloser(strings.NewReader("")),
	}, nil).Once()

	ok, message = client.Save()

	assert.False(t, ok)
	assert.NotEmpty(t, message)

	/* ============================= */
	httpMock.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{}, errors.New("Error")).Once()

	ok, message = client.Save()

	assert.False(t, ok)
	assert.NotEmpty(t, message)
}
