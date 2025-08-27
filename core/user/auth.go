package user

import "net/http"

type JWTAuth struct {
	AccesSalt         string
	AccessExpiration  int64 //insecond
	RefreshSalt       string
	RefreshExpiration int64 //insecond
}

func (JWTAuth) MiddlewareHTTP() func(next http.Handler) http.Handler
