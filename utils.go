package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"
)

type MultipartFormField struct {
	Name  string
	Value string
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
