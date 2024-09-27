package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		// 这两者之间的时间，会导致最后的结果打印顺序不同
		fmt.Printf("%c done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("capitalized: %c\n", capitalized)
}

// for i in {1...15}; do go run ./demo/goroutines |rg ':'; done
// 学习PubSub的时候说的顺序不同，一定是因为这里的goroutine
