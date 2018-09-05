package main

import "fmt"

type Customer struct{
  firstName string
  lastName string
  id string
  balance float64
}

func (cust *Customer) deposit(amount float64){
  cust.balance += amount
}

func (cust *Customer) withdraw(amount float64){
  cust.balance -= amount
}

func (cust *Customer) printDetails(){
  fmt.Printf("Customer name is: %s %s, id is: %s and balance in account is: Rs %.2f\n", cust.firstName, cust.lastName, cust.id, cust.balance)
}

func NewCustomer(firstName, lastName, id string, balance float64) *Customer{
  return &Customer{firstName: firstName, lastName: lastName, id: id, balance: balance}
}
