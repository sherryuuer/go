package main

import "fmt"

func printInfo(title string, route []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(route); i++ {
		route := route[i]
		fmt.Println(route)
	}
}

func main() {
	route := []string{"Home", "Bus Stop", "Cafe Shop", "Office"}
	printInfo("Route #1", route)

	route = append(route, "Nekko Shop")
	printInfo("Route #2", route)

	fmt.Println()
	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")
	route = route[2:]

	printInfo("Remaining locations", route)

}
