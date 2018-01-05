package main 

import(
  "fmt"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  fmt.Println("Server initializing")
  
  e := echo.New()
  
  // "/" represents the root URL.
  e.GET("/", rootHandler, middleware.Logger())
    
  e.GET("/users", getUsers, middleware.Logger())
  
  e.GET("/users/:userId", getUser, middleware.Logger())
  
  e.POST("/users", createUser, middleware.Logger())
  
  // e.POST("/users/:userId", editUser)
  // 
  // e.POST("/users/:userId", deleteUser)
  
  // e.Start starts the server.
  e.Start(":8000")
}

