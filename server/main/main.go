package main 

import(
  "fmt"
  "net/http"
  "github.com/labstack/echo"
)

func main() {
  fmt.Println("Server initializing")
  
  e := echo.New()
  
  // "/" represents the root URL.
  e.GET("/", func (c echo.Context) error {
    return c.String(http.StatusOK, "Server started")
  })
  
  // e.Start starts the server.
  e.Start(":8000")
}

