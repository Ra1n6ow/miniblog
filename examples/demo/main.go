package main

import "fmt"

func main() {
	var a = make([]int, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
}
