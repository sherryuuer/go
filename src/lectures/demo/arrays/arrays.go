package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i] // 在并发处理的时候，养成复制的习惯，可以防止很多问题
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Reception"},
		{name: "HeadOffice"},
		{name: "Operation"},
	}

	checkCleanliness(rooms)

	fmt.Println("Cleaning...")
	rooms[0].cleaned = true
	rooms[1].cleaned = true

	checkCleanliness(rooms)
}
