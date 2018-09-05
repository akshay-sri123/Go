package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCalculator(t *testing.T) {
	t.Run("Should add the 2 numbers", func(t *testing.T) {
		a := 1
		b := 2

		c := &Calculator{num1: a, num2: b}
		res := c.add()

		if res != 3 {
			t.Fatalf("Expected 3, got %v", res)
		}
	})

	t.Run("Should subtract the 2 numbers", func(t *testing.T) {
		a := 6
		b := 2

		c := &Calculator{num1: a, num2: b}
		res := c.sub()

		if res != 4 {
			t.Fatalf("Expected 4, got %v", res)
		}
	})

	t.Run("Should multiply the 2 numbers", func(t *testing.T) {
		a := 1
		b := 2

		c := &Calculator{num1: a, num2: b}
		res := c.mul()

		if res != 2 {
			t.Fatalf("Expected 2, got %v", res)
		}
	})

	t.Run("Should divide the 2 numbers", func(t *testing.T) {
		a := 1
		b := 2

		c := &Calculator{num1: a, num2: b}
		res := c.div()

		if res != 0 {
			t.Fatalf("Expected 0, got %v", res)
		}
	})
}
