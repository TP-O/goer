package main

import (
	"bytes"
)

type Payload struct {
	Type string
	Body *bytes.Buffer
}

type Credentials struct {
	ID       string
	Password string
}

type PayloadGeneratorInterface interface {
	LoginPayload() Payload
	RegistrationPayload(id string) (Payload, string)
	SavePayload() Payload
}

type PayloadGenerator struct {
	credentials Credentials
}

func (p *PayloadGenerator) LoginPayload() Payload {
	fields := []MultipartFormField{
		{
			Name:  "__EVENTTARGET",
			Value: "",
		},
		{
			Name:  "__EVENTARGUMENT",
			Value: "",
		},
		{
			Name:  "ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$txtTaiKhoa",
			Value: p.credentials.ID,
		},
		{
			Name:  "ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$txtMatKhau",
			Value: p.credentials.Password,
		},
		{
			Name:  "ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$btnDangNhap",
			Value: "Đăng Nhập",
		},
	}

	return CreateMultipartFormPayload(fields)
}

func (p *PayloadGenerator) RegistrationPayload(id string) (Payload, string) {
	body, course := CreateRegistrationBody(id)

	return Payload{
			Type: "text/plain; charset=utf-8",
			Body: body,
		},
		course
}

func (p *PayloadGenerator) SavePayload() Payload {
	body := bytes.NewBuffer([]byte(`{
		"isCheckSongHanh": false,
		"ChiaHP": false
	}`))

	return Payload{
		Type: "text/plain; charset=utf-8",
		Body: body,
	}
}
