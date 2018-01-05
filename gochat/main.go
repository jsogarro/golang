package main 

import (
  "log"
  "net/http"
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  
  // Middlewares
  g := e.Group("/gobot/api/v1/")
  
  // Routes
  g.GET("/verify", VerifyRequestHandler)
  g.POSt("/message", SendMessageHandler)

  // Start the server.
  e.Start(":8000")
}