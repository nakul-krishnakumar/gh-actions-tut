package main

import (
	"fmt"

	"github.com/nakul-krishnakumar/gt-actions-tut/internal"
)

func main() {
	adder := adder.New(8, 9)

	fmt.Printf("Output: %d", adder.GetResult())
}