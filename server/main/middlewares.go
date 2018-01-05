package main 

import(
  "github.com/labstack/echo"
)

func AuthMiddleware(username string, password string, c echo.Context) (bool, error) {
  if username == "guest" && password == "User123" {
    return true, nil
  }
  
  return false, nil
}

func CreateHeaders(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    c.Response().Header().Set(echo.HeaderServer, "EchoServer/Golang")
    
    return next(c)
  }
}