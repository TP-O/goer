package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tp-o/goer/mock"
)

func TestNewGoer(t *testing.T) {
	origin := "http://google.com/"
	client := NewGoerClient()
	goer := NewGoer(origin, client)

	assert.NotNil(t, goer)
	assert.Equal(t, origin, goer.Origin)
	assert.Equal(t, client, goer.Client)
}

func TestGoerLogin(t *testing.T) {
	origin := "http://google.com/"
	credentials := &Credentials{}
	failedHeader := http.Header{}
	failedHeader.Add("Location", "sessionreuse")
	successfulHeader := http.Header{}
	successfulHeader.Add("Location", "/")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockIGoerClient(ctrl)
	//================================================
	goer := NewGoer(origin, mockClient)

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(nil, errors.New("Test error"))
	assert.False(t, goer.Login(credentials))

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			StatusCode: 200,
		}, nil)
	assert.False(t, goer.Login(credentials))

	mockClient.EXPECT().Do(gomock.Any()).
		Return(&http.Response{
			Header: failedHeader,
		}, nil)
	assert.False(t, goer.Login(credentials))

	mockClient.EXPECT().Do(gomock.Any()).
		Return(&http.Response{
			StatusCode: 302,
			Header:     successfulHeader,
		}, nil)
	assert.True(t, goer.Login(credentials))
}

func TestGoerClear(t *testing.T) {
	origin := "http://google.com/"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockIGoerClient(ctrl)
	//================================================
	goer := NewGoer(origin, mockClient)

	mockClient.
		EXPECT().
		DeleteSessionId(gomock.Eq(origin)).
		Return(errors.New("Test error"))
	assert.False(t, goer.Clear())

	mockClient.
		EXPECT().
		DeleteSessionId(gomock.Eq(origin)).
		Return(nil)
	assert.True(t, goer.Clear())
}

func TestGoerGreet(t *testing.T) {
	//
}

func TestGoerIsRegistrationOpen(t *testing.T) {
	origin := "http://google.com/"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockIGoerClient(ctrl)
	//================================================
	goer := NewGoer(origin, mockClient)

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(nil, errors.New("Test error"))
	assert.False(t, goer.IsRegistrationOpen())

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			Body: io.NopCloser(strings.NewReader(
				"<p id=\"" +
					strings.Replace(CourseAlertSelector, "#", "", 1) +
					"\">Alert message</p>",
			)),
		}, nil)
	assert.False(t, goer.IsRegistrationOpen())

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			Body: io.NopCloser(strings.NewReader("<p></p>")),
		}, nil)
	assert.True(t, goer.IsRegistrationOpen())
}

func TestRegisterCourse(t *testing.T) {
	origin := "http://google.com/"
	courseID := "IT093IU02  01|IT093IU|Web Application Development|02|4|0|01/01/0001|0|0|0| |0|ITIT19CS31"
	coruseName := strings.Split(courseID, "|")[2]

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockIGoerClient(ctrl)
	//================================================
	goer := NewGoer(origin, mockClient)

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(nil, errors.New("Test error"))
	assert.False(t, goer.RegisterCourse(courseID))

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			Body: io.NopCloser(strings.NewReader("Response ncc")),
		}, nil)
	assert.False(t, goer.RegisterCourse(courseID))

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			Body: io.NopCloser(strings.NewReader("{" + coruseName + "}")),
		}, nil)
	assert.True(t, goer.RegisterCourse(courseID))
}

func TestSaveRegistration(t *testing.T) {
	origin := "http://google.com/"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockIGoerClient(ctrl)
	//================================================
	goer := NewGoer(origin, mockClient)

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(nil, errors.New("Test error"))
	assert.False(t, goer.SaveRegistration())

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			Body: io.NopCloser(strings.NewReader("Response ncc")),
		}, nil)
	assert.False(t, goer.SaveRegistration())

	mockClient.
		EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			Body: io.NopCloser(strings.NewReader("||default.aspx?page=dkmonhoc")),
		}, nil)
	assert.True(t, goer.SaveRegistration())
}
