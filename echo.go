package main

import (
  "fmt"
  "os"
  "time"
  "strings"
)

// No idea yet how to cast a number to a string for concatenation
func main() {
  fmt.Println(concatenate() / join() * 100, "%")
}

// had no idea how to know that time.Now was int64 until compilation
// did not know that Nanoseconds was one of the apis for time until I looked
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

// did not know that specifying the return (end int64) counted as a declaration
// and therefor end := time.Since(...) was wrong
func join() (end int64) {
  start := time.Now()
  strings.Join(os.Args[1:], " ")
  end = time.Since(start).Nanoseconds()
  return
}