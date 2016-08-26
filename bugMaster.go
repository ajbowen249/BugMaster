package main

import (
	"fmt"

	"github.com/ajbowen249/GoSandbox/console"
)

func main() {
	numCols := 80
	numRows := 25
	console.ClearScreen(numCols, numRows)
	console.SetTitle("Bug Master!")

	fmt.Println("password: ")
}
