package tiktok

import (
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	r         *resty.Client
	deviceID  string
	userAgent string
	debug     bool
}

/*
Cookies take from browser cookies when u doing request, device id check on browser local storage and find user_unique_id.

The Cookies must come from firefox browser on pc.
*/
func New(cookies string, deviceID string) *Client {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0"
	r := resty.New()

	// assign cookies if exist
	if cookies != "" {
		splittedCookies := strings.Split(cookies, ";")
		httpCookies := make([]*http.Cookie, 0, len(splittedCookies))
		for _, v := range splittedCookies {
			keyVal := strings.Split(v, "=")
			cookie := &http.Cookie{
				Name:   keyVal[0],
				Value:  keyVal[1],
				Secure: true,
				Domain: ".tiktok.com",
			}
			httpCookies = append(httpCookies, cookie)
		}
		r.SetCookies(httpCookies)
	}
	r.SetHeader("user-agent", userAgent)
	return &Client{
		r:         r,
		deviceID:  deviceID,
		userAgent: userAgent,
	}
}
