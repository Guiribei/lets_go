package main

import "fmt"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)


	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
}