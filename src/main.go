package main

import (
  "os"
)

func main () {
  switch os.Args[1] {
  case "run":
    Run()
  case "child":
    Child()
  default:
    panic("Not sure, what to do")
  }
}

