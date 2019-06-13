package main

import (
	"fmt"
	"os"
	"strconv"
)
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func Kelvin(c Fahrenheit) Celsius {
	fmt.Println("输入的为：", c)
	return Celsius(c) + AbsoluteZeroC
}

func main() {
	args := os.Args[1:]

	for _, v := range args{
		str, _ := strconv.ParseFloat(v, 10)
		fmt.Println("单位转换后的结果是:", Kelvin(Fahrenheit(str)))
	}
}