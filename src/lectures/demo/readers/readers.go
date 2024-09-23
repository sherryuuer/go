package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := r.ReadString(' ')

		// 检查输入是否为空
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}

		// 检查数值转换
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr) // 如果不是一个合法数字，就会被无视和跳过
		} else {
			sum += num
		}

		// 检查是否是输入的最后
		if inputErr == io.EOF { // 这种结束输入机制需要通过手动Ctrl+D来发送（Windows是Ctrl+Z）
			break
		}
		// 检查输入错误
		if inputErr != nil {
			fmt.Println(inputErr)
		}
	}

	fmt.Printf("sum: %v\n", sum)
}
