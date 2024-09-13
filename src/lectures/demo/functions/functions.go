package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello, world!")
}

func main() {
	// go的function写法真是和java反着哈哈
	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen) // 这里不需要在is后面加空格！

	bakerDozen := add(dozen, 1) // 这里也可以直接嵌套
	fmt.Println("A baker's dozen is", bakerDozen)
}
