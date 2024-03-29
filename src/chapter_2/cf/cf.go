package main

import (
	"fmt"
	"os"
	"strconv"
	"chapter_2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n",err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
	}
}