package main

import (
  "fmt"
  "strings"
  )

func main() {

  words := [3]string {"Gate", "Comet", "Pizza"}

  for _, word := range words {

    if(strings.Contains(word, "e")) {

      fmt.Printf("Achei: %c", word[0]);
    }
  }

}
