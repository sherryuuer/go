package msg

import "github.com/sherryuuer/go/lectures/demo/mypkg/display"

// 由于安装了扩展，当保存，vscode会帮我自动import，这个包是从go.mod设置的根目录开始索引的

func Hi() {
	display.Display("Hi")
}
