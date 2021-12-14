package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	color.Red(fmt.Sprintf("Result = %d", calculator.Add(100, 200)))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.GetOpCount())
	fmt.Println(utils.IsEven(100))
}
