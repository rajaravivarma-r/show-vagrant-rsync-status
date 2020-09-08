package main

import (
  "bufio"
  "fmt"
  "os"
  "time"
)

func main() {
  inputChannel := make(chan string)
  secondsElapsedSinceLastInput := 0
  finishedFlag := false
  ticker := time.NewTicker(time.Second)

  go func() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
      inputChannel <- scanner.Text()
    }
    finishedFlag = true

    if scanner.Err() != nil {
      panic(scanner.Err())
    }
  }()

  go func() {
    for {
      inputVal := <-inputChannel
      fmt.Println()
      fmt.Println(inputVal)
      if finishedFlag {
        break
      }
      secondsElapsedSinceLastInput = 0
    }
  }()

  for {
    if finishedFlag {
      ticker.Stop()
      break
    }
    <-ticker.C
    secondsElapsedSinceLastInput = secondsElapsedSinceLastInput + 1
    fmt.Printf("\033[2K\r")
    fmt.Printf("Time elapsed since last sync: %d", secondsElapsedSinceLastInput)
  }
  fmt.Println("\nExiting...")
}
