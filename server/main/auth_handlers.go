package main 

import(
  "net/http"
  "github.com/labstack/echo"
)

func basicHandler (c echo.Context) error {
  return c.String(http.StatusOK, "Basic handler")
}
