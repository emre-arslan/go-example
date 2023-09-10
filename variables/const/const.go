package main

import "fmt"

const const_1 = "değer 1"

const (
	const_2 = "deger2"
	const_3 = "deger3"
)

const (
	const_4 = iota
	const_5
	const_6
)

func main() {

	fmt.Println(const_1, const_2, const_3, const_4, const_5, const_6)
}
