package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

type Goer struct {
	Origin string
	Http   *http.Client
}

func NewGoer(origin string) *Goer {
	var httpClient *http.Client

	if jar, err := cookiejar.New(nil); err != nil {
		panic(SystemFailureMessage + err.Error())
	} else {
		httpClient = &http.Client{
			Jar:     jar,
			Timeout: 60 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	return &Goer{
		Origin: origin,
		Http:   httpClient,
	}
}

func (g *Goer) Login(credentials *Credentials) bool {
	payload := GenerateLoginPayload(credentials)
	req, _ := http.NewRequest("POST", g.Origin+LoginPath, payload.Body)
	req.Header.Add("Content-Type", payload.Type)

	if res, err := g.Http.Do(req); err != nil {
		logrus.Warn(SystemFailureMessage + err.Error())
		return false
	} else if res.StatusCode != 302 ||
		strings.Contains(res.Header.Values("Location")[0], "sessionreuse") {

		logrus.Warn(LoginFailureMessage)
		return false
	} else {
		logrus.Info(LoginSuccessMessage)
		return true
	}
}

func (g *Goer) Clear() bool {
	if url, err := url.Parse(g.Origin); err == nil {
		g.Http.Jar.SetCookies(url, []*http.Cookie{
			{
				Name:     SessionIDCookieField,
				Value:    "",
				HttpOnly: true,
			},
		})

		logrus.Info(LogoutSuccessMessage)
		return true
	}

	logrus.Warn(LogoutFailureMessage)
	return false
}

func (g *Goer) Greet() {
	req, _ := http.NewRequest("GET", g.Origin+HomePath, nil)

	if res, err := g.Http.Do(req); err != nil {
		logrus.Fatalf(SystemFailureMessage + err.Error())
	} else {
		document, _ := goquery.NewDocumentFromReader(res.Body)
		logrus.Info(document.Find(UserGreetingSelector).Text())
	}
}

func (g *Goer) IsRegistrationOpen() bool {
	req, _ := http.NewRequest("GET", g.Origin+CourseListPath, nil)

	if res, err := g.Http.Do(req); err != nil {
		logrus.Warn(SystemFailureMessage + err.Error())
		return false
	} else {
		document, _ := goquery.NewDocumentFromReader(res.Body)

		if document.Find(CourseAlertSelector).Text() == "" {
			logrus.Info(RegistrationIsOpenMessage)
			return true
		}

		logrus.Warn(RegistrationIsNotOpenMessage)
		return false
	}
}

func (g *Goer) RegisterCourse(courseId string) bool {
	courseName := strings.Split(courseId, "|")[2]
	payload := GenerateRegisterCoursePayload(courseId)
	req, _ := http.NewRequest("POST", g.Origin+RegisterCoursePath, payload.Body)
	req.Header.Add("Content-Type", payload.Type)
	req.Header.Add("X-AjaxPro-Method", RegisterCourseAjaxMethod)

	if res, err := g.Http.Do(req); err != nil {
		logrus.Warn(SystemFailureMessage + err.Error())
		return false
	} else {
		resBody, _ := io.ReadAll(res.Body)

		if bytes.Contains(resBody, []byte(courseName)) {
			logrus.Info(RegistrationSuccessMessage + "[" + courseName + "]")
			return true
		}

		logrus.Warn(RegistrationFailureMessage + "[" + courseName + "]")
		return false
	}
}

func (g *Goer) SaveRegistration() bool {
	payload := GenerateCourseSavePayload()
	req, _ := http.NewRequest("POST", g.Origin+SaveCoursePath, payload.Body)
	req.Header.Add("Content-Type", payload.Type)
	req.Header.Add("X-AjaxPro-Method", SaveCourseAjaxMethod)

	if res, err := g.Http.Do(req); err != nil {
		logrus.Warn(SystemFailureMessage + err.Error())
		return false
	} else {
		resBody, _ := io.ReadAll(res.Body)

		if bytes.Contains(resBody, []byte("||default.aspx?page=dkmonhoc")) {
			logrus.Info(SaveSuccessMessage)
			return true
		}

		logrus.Info(SaveFailureMessage)
		return false
	}
}
