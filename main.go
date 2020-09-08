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

  go func(secondsElapsedSinceLastInput *int) {
    for {
      inputVal := <-inputChannel
      fmt.Println(inputVal)
      if finishedFlag {
        break
      }
      *secondsElapsedSinceLastInput = 0
    }
  }(&secondsElapsedSinceLastInput)

  for {
    if finishedFlag {
      ticker.Stop()
      break
    }
    <-ticker.C
    secondsElapsedSinceLastInput = secondsElapsedSinceLastInput + 1
    fmt.Printf("%d", secondsElapsedSinceLastInput)
  }
  fmt.Println("Exiting...")
}
