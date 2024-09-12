package main

import "fmt"

func main() {
	var myName = "Sally"
	fmt.Println("My name is ", myName)

	var age int
	fmt.Println("Age is ", age)

	name := "admin"
	fmt.Println("name = ", name)

	part1, other := 1, 5
	fmt.Println("Part1 = ", part1, "Other = ", other)
	part2, other := 2, 6
	fmt.Println("Part2 = ", part2, "Other = ", other)

	age = part1 + part2
	fmt.Println("Age = ", age)

	var (
		lessonName = "variable"
		lessonType = "demo"
	)
	fmt.Println(lessonName, lessonType)

}
