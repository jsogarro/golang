package main

import "fmt"

func main() {
  cities := map[string]string{
    "NYC": "New York City",
    "BOS": "Boston",
    "PHI": "Philiadelphia",
  }

  fmt.Println(cities)
  fmt.Println(cities["NYC"])

  numbers := make(map[string]int)
  numbers["One"] = 1
  numbers["Two"] = 2
  numbers["Three"] = 3
  fmt.Println(numbers)

  // delete property
  delete(numbers, "Three")
  fmt.Println(numbers)

  showMap(cities)
}

func showMap(m map[string]string) {
  for k, v := range m {
    fmt.Println(k)
    fmt.Println(v)
  }
}
