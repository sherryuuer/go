package main

import (
	"github.com/sherryuuer/go/lectures/demo/mypkg/display"
	"github.com/sherryuuer/go/lectures/demo/mypkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hi")
	// 可以看到这里的pkg名和文件名没有关系，只和文件夹有关系
	// 在函数定义的头部定义了package的名字，引用那个就OK
	msg.Exciting("Hellloooo!")
}
