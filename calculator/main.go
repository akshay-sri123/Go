package main

import (
	"fmt"
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("Calculator", "CLI Calculator")

	addCommand = app.Command("add", "Adds two numbers.")
	addArg1   = addCommand.Arg("addArg1", "1st value").Required().Int()
	addArg2   = addCommand.Arg("addArg2", "2nd value").Required().Int()

	subCommand = app.Command("sub", "Subtracts two numbers.")
	subArg1   = subCommand.Arg("subArg1", "1st value").Required().Int()
	subArg2   = subCommand.Arg("subArg2", "2nd value").Required().Int()

	mulCommand = app.Command("mul", "Multiplies two numbers.")
	mulArg1   = mulCommand.Arg("mulArg1", "1st value").Required().Int()
	mulArg2   = mulCommand.Arg("mulArg2", "2nd value").Required().Int()

	divCommand = app.Command("div", "Divides two numbers.")
	divArg1   = divCommand.Arg("divArg1", "1st value").Required().Int()
	divArg2   = divCommand.Arg("divArg2", "2nd value").Required().Int()
)

func main() {
	var calculator *Calculator
	// fmt.Println()
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		case addCommand.FullCommand():
			calculator = NewCalculator(*addArg1, *addArg2)
			fmt.Printf("Result is: %v\n", calculator.add())

		case subCommand.FullCommand():
			calculator = NewCalculator(*subArg1, *subArg2)
			fmt.Printf("Result is: %v\n", calculator.sub())

		case mulCommand.FullCommand():
			calculator = NewCalculator(*mulArg1, *mulArg2)
			fmt.Printf("Result is: %v\n", calculator.mul())

		case divCommand.FullCommand():
			calculator = NewCalculator(*divArg1, *divArg2)
			fmt.Printf("Result is: %v\n", calculator.div())
	}
}
