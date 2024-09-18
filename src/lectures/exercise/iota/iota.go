//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import (
	"fmt"
)

const (
	add = iota
	sub
	mul
	div
)

type Operation int

func (ope Operation) calculate(a, b int) int {
	switch ope {
	case add:
		return a + b
	case sub:
		return a - b
	case mul:
		return a * b
	case div:
		if b != 0 {
			return a / b
		} else {
			fmt.Println("Can not be divided by 0")
		}
	}
	panic("unhandled operation")
}

func main() {
	add := Operation(add)
	fmt.Println(add.calculate(2, 2)) // = 4
	sub := Operation(sub)
	fmt.Println(sub.calculate(10, 3)) // = 7
	mul := Operation(mul)
	fmt.Println(mul.calculate(3, 3)) // = 9
	div := Operation(div)
	fmt.Println(div.calculate(100, 2)) // = 50
}

// iota感觉主要用于和定数，以及switch一起使用
