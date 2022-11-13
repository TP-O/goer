package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strings"
)

type Credentials struct {
	ID       string
	Password string
}

type Payload struct {
	Type string
	Body *bytes.Buffer
}

type RegisterCourseBody struct {
	IsValidCoso          bool   `json:"isValidCoso"`
	IsValidTKB           bool   `json:"isValidTKB"`
	MaDK                 string `json:"maDK"`
	MaMH                 string `json:"maMH"`
	Sotc                 string `json:"sotc"`
	TenMH                string `json:"tenMH"`
	MaNh                 string `json:"maNh"`
	StrsoTCHP            string `json:"strsoTCHP"`
	IsCheck              string `json:"isCheck"`
	OldMaDK              string `json:"oldMaDK"`
	StrngayThi           string `json:"strngayThi"`
	TietBD               string `json:"tietBD"`
	SoTiet               string `json:"soTiet"`
	IsMHDangKyCungKhoiSV string `json:"isMHDangKyCungKhoiSV"`
}

func GenerateMultipartFormPayload(fields map[string]string) Payload {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	defer writer.Close()

	for key, val := range fields {
		fw, _ := writer.CreateFormField(key)
		io.Copy(fw, strings.NewReader(val))
	}

	return Payload{
		Type: writer.FormDataContentType(),
		Body: payload,
	}
}

func GenerateLoginPayload(credentials *Credentials) Payload {
	loginFileds := map[string]string{
		"__EVENTTARGET":      "",
		"__EVENTARGUMENT":    "",
		IDInputName:          credentials.ID,
		PasswordInputName:    credentials.Password,
		LoginActionInputName: "Đăng Nhập",
	}

	return GenerateMultipartFormPayload(loginFileds)
}

func GenerateRegisterCoursePayload(courseId string) Payload {
	extractedProperties := strings.Split(courseId, "|")
	body := RegisterCourseBody{
		IsValidCoso:          false,
		IsValidTKB:           false,
		MaDK:                 extractedProperties[0],
		MaMH:                 extractedProperties[1],
		Sotc:                 extractedProperties[4],
		TenMH:                extractedProperties[2],
		MaNh:                 extractedProperties[3],
		StrsoTCHP:            "0",
		IsCheck:              "true",
		OldMaDK:              extractedProperties[10],
		StrngayThi:           extractedProperties[6],
		TietBD:               extractedProperties[8],
		SoTiet:               extractedProperties[9],
		IsMHDangKyCungKhoiSV: "0",
	}
	jsonBody, _ := json.Marshal(body)

	return Payload{
		Type: "text/plain; charset=utf-8",
		Body: bytes.NewBuffer(jsonBody),
	}
}

func GenerateCourseSavePayload() Payload {
	body := bytes.NewBuffer([]byte(`{
		"isCheckSongHanh": false,
		"ChiaHP": false
	}`))

	return Payload{
		Type: "text/plain; charset=utf-8",
		Body: body,
	}
}
