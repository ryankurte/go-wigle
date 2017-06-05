/**
 * WiGLE API Connector
 * Based on: https://github.com/nfvs/wigle-api/tree/master/lib
 *
 * Copyright 2017 Ryan Kurte
 */

package wigle

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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

	QueryUrl      = "/gps/gps/main/confirmquery/"
	QueryLatSouth = "latrange1"
	QueryLatNorth = "latrange2"
	QueryLngEast  = "longrange1"
	QueryLngWest  = "longrange2"
	QuerySsid     = "ssid"
	QueryOffset   = "pagestart"
)

func NewWiGLE() *WiGLE {
	wigle := WiGLE{}

	jar, _ := cookiejar.New(&cookiejar.Options{})

	wigle.client = http.Client{
		Jar: jar,
	}

	return &wigle
}

func (gle *WiGLE) login(username, password string) error {
	data := url.Values{}
	data.Set(LoginUsername, username)
	data.Set(LoginPassword, password)
	data.Set(LoginNoExpire, "on")

	resp, err := gle.client.PostForm(BaseUrl+LoginUrl, data)
	if err != nil {
		return err
	}

	if len(resp.Cookies()) == 0 {
		return fmt.Errorf("Error logging in: no cookie set")
	}

	// TODO: prolly a better check of login success

	return nil
}
