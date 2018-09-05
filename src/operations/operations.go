package operations

import (
  "fmt"
  "strconv"
  )

func Add(num1, num2 string) error {
  a, _ := strconv.Atoi(num1)
  b, _ := strconv.Atoi(num2)
  fmt.Printf("%v\n", a+b)
  return nil
}

func Sub(num1, num2 string) error {
  a, _ := strconv.Atoi(num1)
  b, _ := strconv.Atoi(num2)
  fmt.Printf("%v\n", a-b)
  return nil
}

func Mul(num1, num2 string) error {
  a, _ := strconv.Atoi(num1)
  b, _ := strconv.Atoi(num2)
  fmt.Printf("%v\n", a*b)
  return nil
}

func Div(num1, num2 string) error {
  a, _ := strconv.Atoi(num1)
  b, _ := strconv.Atoi(num2)
  if b != 0{
    fmt.Printf("%v\n", a/b)
  } else{
    fmt.Printf("Divide by 0 error")
    }
  return nil
}
