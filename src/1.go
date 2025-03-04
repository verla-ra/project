package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	x := rand.Intn(100)
	y := rand.Intn(100)
	z := x + y
	fmt.Println("The sum of", x, "and", y, "is", z)
}
