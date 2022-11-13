package main

// import (
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateMultipartFormPayload(t *testing.T) {
// 	fields := []MultipartFormField{
// 		{
// 			Name:  "field01",
// 			Value: "value01",
// 		},
// 	}
// 	payload := CreateMultipartFormPayload(fields)

// 	assert.True(t, strings.HasPrefix(payload.Type, "multipart/form-data; boundary="))
// 	assert.Contains(t, payload.Body.String(), "Content-Disposition: form-data; name=\""+fields[0].Name+"\"")
// 	assert.Contains(t, payload.Body.String(), fields[0].Value)
// }

// func TestNewHttp(t *testing.T) {
// 	http := NewHttp()

// 	assert.NotEmpty(t, http)
// 	assert.NotEmpty(t, http.GetClient().Jar)
// 	assert.Equal(t, http.GetClient().Timeout, 60*time.Second)
// }

// func TestCreateRegistrationBody(t *testing.T) {
// 	body, course := CreateRegistrationBody("MaDK|MaMH|TenMH|MaNh|Sotc||StrngayThi||TietBD|SoTiet|")

// 	assert.True(t, course == "TenMH")
// 	assert.True(t, strings.Contains(body.String(), "\"isValidCoso\":false"))
// 	assert.True(t, strings.Contains(body.String(), "\"isValidTKB\":false"))
// 	assert.True(t, strings.Contains(body.String(), "\"strsoTCHP\":\"0\""))
// 	assert.True(t, strings.Contains(body.String(), "\"isCheck\":\"true\""))
// 	assert.True(t, strings.Contains(body.String(), "\"isMHDangKyCungKhoiSV\":\"0\""))
// 	assert.True(t, strings.Contains(body.String(), "\"maDK\":\"MaDK\""))
// 	assert.True(t, strings.Contains(body.String(), "\"maMH\":\"MaMH\""))
// 	assert.True(t, strings.Contains(body.String(), "\"tenMH\":\"TenMH\""))
// 	assert.True(t, strings.Contains(body.String(), "\"maNh\":\"MaNh\""))
// 	assert.True(t, strings.Contains(body.String(), "\"sotc\":\"Sotc\""))
// 	assert.True(t, strings.Contains(body.String(), "\"strngayThi\":\"StrngayThi\""))
// 	assert.True(t, strings.Contains(body.String(), "\"tietBD\":\"TietBD\""))
// 	assert.True(t, strings.Contains(body.String(), "\"soTiet\":\"SoTiet\""))
// }
