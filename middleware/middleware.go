package middleware

import (
	"github.com/labstack/echo/middleware"
)

var IsAuthenticate = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
