//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("Checking out...")
	// 下面的方法会创建一个副本，而不会改变原本的struct，要使用index来修改
	// for _, item := range items {
	// 	deactivate(&item.tag)
	// }
	for i, _ := range items {
		deactivate(&items[i].tag)
	}
}

func main() {
	frontEnd := Item{name: "frontEnd", tag: Active}
	backEnd := Item{name: "backEnd", tag: Active}
	clientSide := Item{name: "clientSide", tag: Active}
	fireWall := Item{name: "fireWall", tag: Active}

	items := []Item{frontEnd, backEnd, clientSide, fireWall}
	fmt.Println(items)

	deactivate(&items[0].tag)
	checkout(items)
	fmt.Println(items)

}
