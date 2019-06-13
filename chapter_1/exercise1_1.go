//练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字。
package main

import (
	"os"
	"fmt"
)
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
