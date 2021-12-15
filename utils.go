package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strings"
)

type MultipartFormField struct {
	Name  string
	Value string
}

type RegistrationBody struct {
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

func CreateMultipartFormPayload(fileds []MultipartFormField) Payload {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	defer writer.Close()

	for _, filed := range fileds {
		fw, _ := writer.CreateFormField(filed.Name)
		io.Copy(fw, strings.NewReader(filed.Value))
	}

	return Payload{
		Type: writer.FormDataContentType(),
		Body: payload,
	}
}

func NewHttp() HttpInterface {
	http := Http{}

	http.Initialize()

	return &http
}

func CreateRegistrationBody(id string) (*bytes.Buffer, string) {
	value := strings.Split(id, "|")

	body := RegistrationBody{
		IsValidCoso:          false,
		IsValidTKB:           false,
		MaDK:                 value[0],
		MaMH:                 value[1],
		Sotc:                 value[4],
		TenMH:                value[2],
		MaNh:                 value[3],
		StrsoTCHP:            "0",
		IsCheck:              "true",
		OldMaDK:              value[10],
		StrngayThi:           value[6],
		TietBD:               value[8],
		SoTiet:               value[9],
		IsMHDangKyCungKhoiSV: "0",
	}

	byteBody, _ := json.Marshal(body)

	return bytes.NewBuffer(byteBody), value[2]
}
