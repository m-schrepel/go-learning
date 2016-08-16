package main

import (
  "fmt"
  "os"
  "time"
  "strings"
)

func main() {
  fmt.Println(concatenate() / join() * 100, "%")
}

func concatenate() (end int64) {
  var s, sep string
  start := time.Now()
  for i := 0; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " || "
  }
  end = time.Since(start).Nanoseconds()
  return
}

func join()(end int64) {
  start := time.Now()
  strings.Join(os.Args[1:], " ")
  end = time.Since(start).Nanoseconds()
  return
}