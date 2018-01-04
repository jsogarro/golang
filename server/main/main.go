package main 

import(
  "fmt"
  "net/http"
  "github.com/labstack/echo"
)

func rootHandler (c echo.Context) error {
  return c.String(http.StatusOK, "Server started")
}

func userHandler (c echo.Context) error {
  userId := c.Param("userId")
  fmt.Println(userId)
  
  return c.String(http.StatusOK, fmt.Sprintf("User Id: %s", userId))
}

func usersHandler (c echo.Context) error {
  return c.String(http.StatusOK, "user handler")
}

func createUser(c echo.Context) error {
  return true;
}

func main() {
  fmt.Println("Server initializing")
  
  e := echo.New()
  
  // "/" represents the root URL.
  e.GET("/", rootHandler)
    
  e.GET("/users", usersHandler)
  
  e.GET("/users/:userId", userHandler)
  
  // e.Start starts the server.
  e.Start(":8000")
}

