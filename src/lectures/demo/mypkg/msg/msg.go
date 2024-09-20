package msg

// 由于安装了扩展，当保存，vscode会帮我自动import，这个包是从go.mod设置的根目录开始索引的
import "github.com/sherryuuer/go/src/lectures/demo/mypkg/display"

func Hi() {
	display.Display("Hi")
}
