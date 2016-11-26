package main

import (
  "fmt"
  "strings"
  )

/* Função corrida autotimaticamente antes do main */

func init() {
  // teste do segundo branch do git
}

func main() {

  fmt.Println("Welcome to the real world");

  var x int
  var y int

  x = 10
  y = x + 1

  fmt.Println("ola", " ", x, " ", y, " ", 999);

  //fmt.Println("div: ", (3 / 4));

  var teste bool = true

  complexNum := 4 + 1i

  fmt.Printf("%v is a %T \n", complexNum, complexNum);
  fmt.Printf("%v is a %T \n", teste, teste);

  var outroTeste = false

  fmt.Printf("%v is a %T \n", outroTeste, outroTeste);

  fmt.Printf("%v is a %T \n", x, x);
  fmt.Printf("%v is a %T \n", y, y);

  // outroTeste = 10

  var numero int8 = 120

  numero = 1

  fmt.Printf("%v is a %T \n", numero, numero);

  var nome = "Bruno"
  fmt.Println(nome);
  fmt.Println(strings.ToUpper(nome));

  if(numero > 0) {
    fmt.Println("show di bola");
  }

  // indentação do else - forçada ?
  if num := x + y; num < 0 {
    fmt.Println("Entrou no IF");
  } else {
    fmt.Println("Entrou no ELSE");
  }

  for i := 0; i < 10; i++ {

    if i % 2 == 0 {
      fmt.Println("", i, " par");
    } else {
      fmt.Println(" impar");
    }

  }

  var z int = 10 - 5

  switch z {

  case 1:
    fmt.Println("Um")
  case 2:
    fmt.Println("Dois")
  case 3, 4, 5:
    fmt.Println("Especiais")
  default:
    fmt.Println("Desconhecido")
  }

  // O switch tem break implicito
  // Para seguir, deve-se usar a keyword fallthrough

  var a [5]int
  a[4] = 100;
  fmt.Println(a);

  /* o [...] indica ao compilador que deverá calcular o tamanho */
  //var b := [...]int{1, 2, 3, 4, 5}

  // Slices - arrays que podem crescer

  numS := []int{2, 3, 4}

  sum := 0

  for i, num := range numS {
    fmt.Printf("Index: %d, value: %d\n", i, num)
    sum = sum + num
  }

  fmt.Println("Sum: ", sum)

  res := plus(1, 2)

  fmt.Println("Soma 1+2: ", res)

  res, err := morePlus(4, 5, 6);

  if err == nil {
    fmt.Println("Soma 4+5+6: ", res)
  } else {
    fmt.Println("Erro...")
  }

  // O _ permite ignorar a variável
  // err vs _
}

func plus(a int, b int) int {
  return a + b
}

// 2 args, 2 returns
func morePlus(a, b, c int) (int, error) {
  return a + b + c, nil;
}
