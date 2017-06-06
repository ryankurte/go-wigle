/**
 * WiGLE API Connector
 * Based on: https://github.com/nfvs/wigle-api/tree/master/lib
 *
 * Copyright 2017 Ryan Kurte
 */

package wigle

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"

	"github.com/google/go-querystring"
)

type WiGLE struct {
	client             http.Client
	username, password string
}

const (
	BaseUrl = "https://wigle.net"

	LoginUrl = "/gps/gps/main/login"

	LoginUsername = "credential_0"
	LoginPassword = "credential_1"
	LoginNoExpire = "noexpire"

	QueryUrl = "/gps/gps/main/confirmquery/"
)

func NewWiGLE() *WiGLE {
	wigle := WiGLE{}

	jar, _ := cookiejar.New(&cookiejar.Options{})

	wigle.client = http.Client{
		Jar: jar,
	}

	return &wigle
}

type Request struct {
	SSID   string `url:"ssid"`
	Offset uint32 `url:"pagestart"`

	LatNorth float64 `url:"latrange2"`
	LatSouth float64 `url:"latrange1"`
	LngEast  float64 `url:"longrange1"`
	LngWest  float64 `url:"longrange2"`
}

type Response struct {
}

func (gle *WiGLE) Login(username, password string) error {
	data := url.Values{}
	data.Set(LoginUsername, username)
	data.Set(LoginPassword, password)
	data.Set(LoginNoExpire, "on")

	resp, err := gle.client.PostForm(BaseUrl+LoginUrl, data)
	if err != nil {
		return err
	}

	body, _ := httputil.DumpResponse(resp, true)

	log.Printf("%s", string(body))

	if len(resp.Cookies()) == 0 {
		return fmt.Errorf("Error logging in: no cookie set")
	}

	// TODO: prolly a better check of login success

	return nil
}
