package main

import (
  "fmt"
//  "strings"
  )

type newstringtype string;

type person struct {
  name string
  age int
}

func main() {

  fmt.Println(person{"Bob", 20})

  fmt.Println(person{name: "Alice", age: 30})

  fmt.Println(person{name: "Fred"})

  p := person{name:"Bruno", age: 33}

  fmt.Println(p.name)
  fmt.Println(p.age)

}

// Função pode receber PTR para struct
// ou a própria struct
 
