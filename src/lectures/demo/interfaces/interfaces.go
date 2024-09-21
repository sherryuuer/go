package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook Chicken!")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chop Salad!")
	fmt.Println("Add dressing!")
}

// 这里的接口作为参数，类似于Java的多态的感觉
func PrepareDishes(dishes []Preparer) {
	fmt.Println("---Prepare dishes---")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("---Dish: %v---\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken wings"),
		Salad("Salad"),
	}
	PrepareDishes(dishes)
}
