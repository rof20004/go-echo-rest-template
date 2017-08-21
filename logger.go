package main

import (
	"fmt"

	"github.com/labstack/echo"
)

// RequestLog middleware adds generate logs for requests.
func RequestLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

// ResponseLog bodydump - adds generate logs for responses.
func ResponseLog(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("%s\n", resBody)
}
