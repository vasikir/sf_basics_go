package main

import (
	"fmt"

	"github.com/vasikir/sf_basics_go/maths"
)

func main() {
	print(maths.Sum, 5, 2)
	print(maths.Sub, 4, 3)
	print(func(a int, b int) int { return a * b }, 3, 5)
}

func print(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}
