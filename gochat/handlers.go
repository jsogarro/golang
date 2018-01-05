package main 

import (
  "net/http"
  "github.com/labstack/echo"
)

func VerifyRequestHandler(c echo.Context) error {
  // TODO: authenticate the request
  return c.String(http.StatusOK, "VerifyRequestHandler")
}

func SendMessageHandler(c echo.Context) error {
  // TODO: Handle message responses
  return c.String(http.StatusOK, "SendMessageHandler")
}