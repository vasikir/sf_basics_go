package main

import (
	"fmt"

	"github.com/vasikir/sf_basics_go/maths"
)

func main() {
	print(maths.Sum, 5, 2)
	print(maths.Sub, 4, 3)
}

func print(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}
