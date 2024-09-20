package display

import "fmt"

// 在Go中这种大写字母开头的函数会自动被export
func Display(msg string) {
	fmt.Println(msg)
}
