//练习 1.4： 修改 dup2 ，出现重复的行时打印文件名称。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			judge := countLines(f, counts)
			if judge == true{
				fmt.Println(f.Name())
			}
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	nums := 0
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		nums++
	}
	if nums > 1 {
		return true
	} else {
		return false
	}
}
