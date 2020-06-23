package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string	// 默认初始化为空字符串""
	for i:=1; i <len(os.Args); i++ {	// Go语言只有for循环这一种循环语句
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}