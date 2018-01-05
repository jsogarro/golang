package main 

import(
  "github.com/labstack/echo"
)

func authMiddleware(username string, password string, c echo.Context) (bool, error) {
  if username == "guest" && password == "User123" {
    return true, nil
  }
  
  return false, nil
}