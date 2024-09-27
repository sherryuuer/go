package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	rand.New(src)

	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter++
		go func() {
			defer func() {
				fmt.Println(counter, "goroutines remaining")
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("Waiting for", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("done!")
}
