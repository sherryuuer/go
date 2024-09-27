package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count++
	fmt.Println("server iteration", iteration)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup

	hitCounter := Hits{}

	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Printf("waiting for goroutines...\n")
	wg.Wait()

	// 这里是在以上的处理都结束之后，可以锁定整个hitCounter来取得count
	// 这里的lock只是为了练习
	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()

	fmt.Printf("\ntotal hits = %d\n", totalHits)
}
