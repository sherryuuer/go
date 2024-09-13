//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//Requirements:
	//* Store your favorite color in a variable using the `var` keyword
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	//* Store your first & last initials in two variables using block assignment
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier

	var favoriteColor = "blue"
	fmt.Println(favoriteColor)

	birthYear1, age := 1998, 28
	birthYear2, age := 1997, 20
	fmt.Println(birthYear1, birthYear2)

	var days int // 如果下面不打印，他就会建议你和下面的赋值写在一行
	fmt.Println(days)
	days = age * 365
	fmt.Println(days)
}
