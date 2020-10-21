package main

import (
	"fmt"
)
func plus(a int, b int) int {
	// Go需要使用return语句显式地返回值
	return a + b
}
func main() {
	fmt.Println(plus(3,4)) //7
}