package main

import "fmt"

type Calculator struct {
	num1 int
	num2 int
}

func (calc *Calculator) add() int {
	return (calc.num1 + calc.num2)
}

func (calc *Calculator) sub() int{
  return (calc.num1 - calc.num2)
}

func (calc *Calculator) mul() int{
  return (calc.num1 * calc.num2)
}

func (calc *Calculator) div() int{
	res := 0
	if calc.num2 == 0{
		fmt.Println("Divide by 0 error")
		res = -1
	}else{
		res = (calc.num1 / calc.num2)
	}
	return res
}

func NewCalculator(num1, num2 int) *Calculator {
	return &Calculator{num1: num1, num2: num2}
}
