package main

import (
  "fmt"
  "os"
)

func main() {
  file, _ := os.ReadFile("file01")
  content := (file)
  fmt.Println(content)
  fmt.Println("AAAA BBBB CCCC")
}
