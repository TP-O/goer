package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	Host             string
	Http             HttpInterface
	PayloadGenerator PayloadGeneratorInterface
}

func (c *Client) Login() (bool, string) {
	path := "/default.aspx"
	payload := c.PayloadGenerator.LoginPayload()

	req, _ := http.NewRequest("POST", c.Host+path, payload.Body)
	req.Header.Add("Content-Type", payload.Type)

	if res, err := c.Http.Do(req); err != nil {
		return false, err.Error() + " 😢 Trying to login again..."
	} else if res.StatusCode != 302 || strings.Contains(res.Header.Values("Location")[0], "sessionreuse") {
		return false, "Login failed 😢 Trying to login again..."
	} else {
		return true, "Login successfully!!! 😆"
	}
}

func (c *Client) SayHi() string {
	path := "/default.aspx"

	req, _ := http.NewRequest("GET", c.Host+path, nil)

	if res, err := c.Http.Do(req); err != nil {
		return "..."
	} else {
		document, _ := goquery.NewDocumentFromReader(res.Body)

		return document.Find("#ctl00_Header1_Logout1_lblNguoiDung").Text()
	}
}

func (c *Client) Register(id string) (bool, error) {
	path := "/ajaxpro/EduSoft.Web.UC.DangKyMonHoc,EduSoft.Web.ashx"
	payload, course := c.PayloadGenerator.RegistrationPayload(id)

	req, _ := http.NewRequest("POST", c.Host+path, payload.Body)
	req.Header.Add("Content-Type", payload.Type)
	req.Header.Add("X-AjaxPro-Method", "LuuVaoKetQuaDangKy")

	if res, err := c.Http.Do(req); err != nil {
		return false, err
	} else {
		resBody, _ := io.ReadAll(res.Body)

		if bytes.Contains(resBody, []byte(course)) {
			return true, nil
		}

		return false, errors.New("Register failed 😢")
	}
}

func (c *Client) Save() (bool, error) {
	path := "/ajaxpro/EduSoft.Web.UC.DangKyMonHoc,EduSoft.Web.ashx"
	payload := c.PayloadGenerator.SavePayload()

	req, _ := http.NewRequest("POST", c.Host+path, payload.Body)
	req.Header.Add("Content-Type", payload.Type)
	req.Header.Add("X-AjaxPro-Method", "LuuDanhSachDangKy_HopLe")

	if res, err := c.Http.Do(req); err != nil {
		return false, err
	} else {
		resBody, _ := io.ReadAll(res.Body)

		if bytes.Contains(resBody, []byte("||default.aspx?page=dkmonhoc")) {
			return true, nil
		}

		return false, errors.New("Save failed 😢")
	}
}
