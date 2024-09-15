package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("sum is", sum)
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println("sum is", sum)
	}
	// 这种写法相当于while，go语言中没有while关键字
	for sum > 10 {
		sum -= 5
		fmt.Println("sum become", sum)
	}
}
