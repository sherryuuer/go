package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		out <- rand.Intn(max-min+1) + min
	}()

	return out
}

func generateRandIntN(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	randInt := generateRandInt(1, 100)
	fmt.Println("Generated number:")
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	randIntRange := generateRandIntN(5, 1, 100)
	fmt.Println("Generated number:")
	for i := range randIntRange {
		fmt.Println("randIntNrange:", i)
	}

}
