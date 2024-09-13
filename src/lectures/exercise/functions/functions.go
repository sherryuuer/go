//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.

package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello,", name)
}

func getMessage(message string) {
	fmt.Println(message)
}

func addThree(a, b, c int) int {
	return a + b + c
}

func getNum(n int) int {
	return n
}

func getTwoNum(n1, n2 int) (int, int) {
	return n1, n2
}

func main() {
	//--Requirements:
	//* Write a function that accepts a person's name as a function
	//  parameter and displays a greeting to that person.
	greeting("Jack")
	//* Write a function that returns any message, and call it from within
	//  fmt.Println()
	getMessage("How are you")
	//* Write a function to add 3 numbers together, supplied as arguments, and
	//  return the answer
	resOfAddThree := addThree(1, 2, 3)
	fmt.Println("The sum of the three is", resOfAddThree)
	//* Write a function that returns any number
	fmt.Println(getNum(7))
	//* Write a function that returns any two numbers
	fmt.Println(getTwoNum(3, 4))
	//* Add three numbers together using any combination of the existing functions.
	two, three := getTwoNum(2, 3)
	sumOfThree := addThree(getNum(1), two, three)
	//  * Print the result
	fmt.Println(sumOfThree)
	//* Call every function at least once

}
