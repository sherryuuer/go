package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap = map[string]int{
		"item1": 1,
		"item2": 2,
		"item3": 3,
	}

	fmt.Println(myMap)
	fmt.Println(len(myMap))

	itemNum := myMap["item1"]
	fmt.Println(itemNum)

	// 这里的found参数是一个很好的检查元素是否存在的方法
	num, found := myMap["item4"]
	fmt.Println(num, found)
	if !found {
		fmt.Println("The element not exsits")
	}

	myMap["item4"] = 4
	fmt.Println("After insert", myMap)

	delete(myMap, "item1")
	fmt.Println("After delete", myMap)

	fmt.Println("Get all the element in the map: ")
	for key, value := range myMap {
		fmt.Println(key, ": ", value)
	}

	totalValue := 0
	for _, v := range myMap {
		totalValue += v
	}
	fmt.Println("The total value of the map is", totalValue)

}
