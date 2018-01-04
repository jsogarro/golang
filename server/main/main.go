package main 

import(
  "fmt"
  "github.com/labstack/echo"
)

func main() {
  fmt.Println("Server initializing")
  
  e := echo.New()
  
  // "/" represents the root URL.
  e.GET("/", rootHandler)
    
  e.GET("/users", getUsers)
  
  e.GET("/users/:userId", getUser)
  
  e.POST("/users", createUser)
  
  // e.POST("/users/:userId", editUser)
  // 
  // e.POST("/users/:userId", deleteUser)
  
  // e.Start starts the server.
  e.Start(":8000")
}

