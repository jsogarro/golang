package main 

import (
  "log"
  "net/http"
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  
  // Middlewares
  g := e.Group("/api/v1/")
  
  // Routes
  g.GET("/webhook", VerifyRequestHandler)
  g.POSt("/webhook", SendMessageHandler)

  // Start the server.
  e.Start(":8000")
}