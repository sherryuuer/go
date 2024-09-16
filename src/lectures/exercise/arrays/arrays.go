//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

// 这里的参数其实是一个slice
func printInfo(a []Product) {
	fmt.Println("The last item of the list is", a[len(a)-1])
	fmt.Println("The total number of the list is", len(a))

	total_price := 0.0
	for i := 0; i < len(a); i++ {
		total_price += a[i].price
	}
	fmt.Println("The total price is", total_price)
}

func main() {
	shoppingList := []Product{
		{name: "laptop", price: 20},
		{name: "Keyboard", price: 32},
		{name: "toybox", price: 12},
	}

	printInfo(shoppingList)

	shoppingList_view := append(shoppingList, Product{name: "ricebag", price: 5.5})

	printInfo(shoppingList_view)

}
