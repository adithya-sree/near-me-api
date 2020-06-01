package middleware

import (
	"nearme-api/src/config"
)

//Middleware middleware
type Middleware struct {
	config config.AppConfig
}

//NewMiddleware gets an instance of the middleware struct
func NewMiddleware(c config.AppConfig) Middleware {
	return Middleware{
		config: c,
	}
}
