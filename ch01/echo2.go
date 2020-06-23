package main

import (
	"fmt"
	"os"
)

func main() {
	
	// s := ""
	// var s string
	// var s = ""
	// var s string = ""

	s, sep := "", ""
	// for 循环的另一种形式, 在某种数据类型的区间(range)上遍历,如字符串或切片
	// 每次循环迭代, range 产生一对值;索引以及在该索引处的元素值。这个例子不需要索引
	// 用 空标识符 (blank identifier),即 _ (也就是下划线)表示
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}