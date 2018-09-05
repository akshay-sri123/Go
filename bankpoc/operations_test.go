package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}


func TestOperations(t *testing.T){
  t.Run("Should deposit amount into account", func(t *testing.T){
    depositAmount := 147.22
    cust := &Customer{firstName: "A", lastName: "B", id: "010", balance: 1470.00}
    cust.deposit(depositAmount)
    if cust.balance != 1617.22{
      t.Fatalf("Expected 1617.22, got %.2f", cust.balance)
      }
  })

  t.Run("Should withdraw amount from account", func(t *testing.T){
    withdrawalAmount := 147.22
    cust := &Customer{firstName: "A", lastName: "B", id: "010", balance: 1470.00}
    cust.withdraw(withdrawalAmount)
    if cust.balance != 1322.78{
      t.Fatalf("Expected 1617.22, got %.2f", cust.balance)
      }
  })


}
