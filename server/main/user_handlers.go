package main 

import(
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "github.com/labstack/echo"
)

func getUsers (c echo.Context) error {
  return c.String(http.StatusOK, "user handler")
}

func getUser (c echo.Context) error {
  userId := c.Param("userId")
  fmt.Println(userId)
  
  return c.String(http.StatusOK, fmt.Sprintf("User Id: %s", userId))
}

func createUser(c echo.Context) error {
  user := User{}
  
  defer c.Request().Body.Close()
  
  b, err := ioutil.ReadAll(c.Request().Body)
  if err != nil {
    log.Printf("BODY REQUEST FAILURE: %s", err)
    return c.String(http.StatusInternalServerError, "")
  }
  
  err = json.Unmarshal(b, &user)
  if err != nil {
    log.Printf("UNMARSHAL FAILURE: %s", err)
    return c.String(http.StatusInternalServerError, "")
  }
  
  fmt.Println("Save successful!")
  return c.String(http.StatusOK, "Saved!")
}

// func editUser(c echo.Context) error {
//   return true;
// }
// 
// func deleteUser(c echo.Context) error {
//   return true;
// }