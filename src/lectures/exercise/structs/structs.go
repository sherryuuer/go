//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

// 这个结构基本就是class相当的东西了，struct
type Rectangle struct {
	length int
	width  int
}

func areaPerimeter(rect Rectangle) (int, int) {
	return rect.length * rect.width, (rect.length + rect.width) * 2
}

func printInfo(rect Rectangle) {
	area, perimeter := areaPerimeter(rect)
	fmt.Println("rectangle area is", area, " perimeter is", perimeter)
}

func main() {
	rectangle := Rectangle{2, 3}

	printInfo(rectangle)

	rectangle.length = 4
	rectangle.width = 6

	printInfo(rectangle)

}
