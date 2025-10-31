package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/nakul-krishnakumar/gt-actions-tut/internal"
)

func main() {
	adder := adder.New(8, 9)

	fmt.Printf("Output: %d\n", adder.GetResult())

	myFigure := figure.NewFigure("Hello World", "", true)
  	myFigure.Print()
}