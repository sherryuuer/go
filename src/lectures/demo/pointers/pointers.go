package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// (*counter).hits += 1
	// Go 会自动解引用指向结构体的指针来访问它的字段，所以不需要像上面这样写
	counter.hits += 1
	fmt.Println("Counter: ", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new // dereference 解除引用，将旧的指针指向新的string，从而改变指针指向的，真正内容！
	increment(counter)
}

func main() {
	counter := Counter{}
	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)
	// 这里的两个&都表达了他们是指针，关注一下replace函数中的参数，就是指针哦
	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	// 一个对数组元素的指针进行修改的例子，有趣
	phrase := []string{"hello", "world"}
	fmt.Println(phrase)
	replace(&phrase[1], "Go", &counter)
	fmt.Println(phrase)
}
