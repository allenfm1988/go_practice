package main

import (
	"fmt"
)

const (
	first = iota + 6
	second = 2 << iota
)

// another constant block without specifying a constant value which will get assigned the previous constant expression.

const (
	first1 = iota + 6
	second1 
)

const (
	first2 = iota
	 
)

func main() {
// usage of iota with constants
	fmt.Println(first, second)

	fmt.Println(first1, second1)

	fmt.Println(first2)

}
