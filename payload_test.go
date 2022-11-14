package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMultipartFormPayload(t *testing.T) {
	fields := map[string]string{
		"F1": "V1",
		"F2": "V2",
	}
	payload := GenerateMultipartFormPayload(fields)
	body := payload.Body.String()

	assert.True(t, strings.HasPrefix(payload.Type, "multipart/form-data; boundary="))

	for key, val := range fields {
		assert.Contains(t, body, "name=\""+key+"\"")
		assert.Contains(t, body, val)
	}
}

func TestGenerateLoginPayload(t *testing.T) {
	credentials := &Credentials{
		ID:       "ITITIU19180",
		Password: "Mypassword",
	}
	payload := GenerateLoginPayload(credentials)
	body := payload.Body.String()

	assert.Contains(t, body, "name=\""+IDInputName+"\"")
	assert.Contains(t, body, "name=\""+PasswordInputName+"\"")
	assert.Contains(t, body, "name=\""+LoginActionInputName+"\"")
	assert.Contains(t, body, credentials.ID)
	assert.Contains(t, body, credentials.Password)
}

func TestGenerateRegisterCoursePayload(t *testing.T) {
	courseID := "IT093IU02  01|IT093IU|Web Application Development|02|4|0|01/01/0001|0|0|0| |0|ITIT19CS31"
	extractedCourseInfo := strings.Split(courseID, "|")
	payload := GenerateRegisterCoursePayload(courseID)
	var body RegisterCourseBody
	json.Unmarshal(payload.Body.Bytes(), &body)

	assert.Equal(t, payload.Type, "text/plain; charset=utf-8")
	assert.Equal(t, body.MaDK, extractedCourseInfo[0])
	assert.Equal(t, body.MaMH, extractedCourseInfo[1])
	assert.Equal(t, body.TenMH, extractedCourseInfo[2])
	assert.Equal(t, body.MaNh, extractedCourseInfo[3])
	assert.Equal(t, body.Sotc, extractedCourseInfo[4])
	assert.Equal(t, body.StrngayThi, extractedCourseInfo[6])
	assert.Equal(t, body.SoTiet, extractedCourseInfo[8])
	assert.Equal(t, body.SoTiet, extractedCourseInfo[9])
	assert.Equal(t, body.OldMaDK, extractedCourseInfo[10])
}

func TestGenerateCourseSavePayload(t *testing.T) {
	payload := GenerateCourseSavePayload()
	body := payload.Body.String()

	assert.Equal(t, payload.Type, "text/plain; charset=utf-8")
	assert.NotEmpty(t, body)
}
