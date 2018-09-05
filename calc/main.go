package main

import (
  "log"
  "os"
  cli "gopkg.in/urfave/cli.v1"
  "operations"
)

func main()  {
  app := cli.NewApp()
  app.Name = "Calc"
  app.Usage = "For calculating"
  app.UsageText = "calc <num1> <num2>"

  app.Commands = []cli.Command{
    {
      Name: "add",
      Aliases: []string{"a"},
      Usage: "Add two numbers.",
      Action: func(c *cli.Context) error {
        a := c.Args().Get(0)
        b := c.Args().Get(1)
        return operations.Add(a, b)
      },
    },
    {
      Name: "sub",
      Aliases: []string{"s"},
      Usage: "Subtract two numbers.",
      Action: func(c *cli.Context) error {
        a := c.Args().Get(0)
        b := c.Args().Get(1)
        return operations.Sub(a, b)
      },
  },
  {
    Name: "mul",
    Aliases: []string{"m"},
    Usage: "Multiply two numbers.",
    Action: func(c *cli.Context) error {
      a := c.Args().Get(0)
      b := c.Args().Get(1)
      return operations.Mul(a, b)
    },
  },
  {
    Name: "div",
    Aliases: []string{"d"},
    Usage: "Divide two numbers.",
    Action: func(c *cli.Context) error {
      a := c.Args().Get(0)
      b := c.Args().Get(1)
      return operations.Div(a, b)
    },
  },
}

  err := app.Run(os.Args)
  if err != nil{
    log.Fatal(err)
  }

}
