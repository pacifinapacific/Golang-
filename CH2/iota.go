package main

import "fmt"

func main() {
	const (
		zero = iota // constの中で１づつ増える
		one  = iota
		two  = iota
	)

	fmt.Println(zero, one, two)

}
