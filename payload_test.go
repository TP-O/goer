package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginInPayload(t *testing.T) {
	payloadGenerator := PayloadGenerator{
		credentials: Credentials{
			ID:       "ITITIU19180",
			Password: "Mypassword",
		},
	}
	payload := payloadGenerator.LoginPayload()

	assert.True(t, strings.HasPrefix(payload.Type, "multipart/form-data; boundary="))
	assert.Contains(t, payload.Body.String(), "name=\"__EVENTTARGET\"")
	assert.Contains(t, payload.Body.String(), "name=\"__EVENTARGUMENT\"")
	assert.Contains(t, payload.Body.String(), "name=\"ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$txtTaiKhoa\"")
	assert.Contains(t, payload.Body.String(), "name=\"ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$txtMatKhau\"")
	assert.Contains(t, payload.Body.String(), "name=\"ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$btnDangNhap\"")
	assert.Contains(t, payload.Body.String(), payloadGenerator.credentials.ID)
	assert.Contains(t, payload.Body.String(), payloadGenerator.credentials.Password)
}

func TestRegistrationPayload(t *testing.T) {
	payloadGenerator := PayloadGenerator{}
	payload, course := payloadGenerator.RegistrationPayload("MaDK|MaMH|TenMH|MaNh|Sotc||StrngayThi||TietBD|SoTiet|")

	assert.Equal(t, course, "TenMH")
	assert.Equal(t, payload.Type, "text/plain; charset=utf-8")
	assert.NotEmpty(t, payload.Body)
}

func TestSavePayload(t *testing.T) {
	payloadGenerator := PayloadGenerator{}
	payload := payloadGenerator.SavePayload()

	assert.Equal(t, payload.Type, "text/plain; charset=utf-8")
	assert.Contains(t, payload.Body.String(), "\"isCheckSongHanh\": false")
	assert.Contains(t, payload.Body.String(), "\"ChiaHP\": false")
}
