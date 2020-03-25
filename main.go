package main

import (
	"fmt"
	"github.com/linalinn/Bruch/bruch"
)

func main() {
	bruch1, _ := bruch.NewBruch(4, 2)
	bruch2, _ := bruch.NewBruch(1, 4)
	bruch3, _ := bruch.NewBruch(6, 5)
	bruch4, _ := bruch.NewBruch(-8, 10)
	bruch5, _ := bruch.NewBruch(12, -4)
	bruch6, _ := bruch.NewBruch(-5, -9)
	run(bruch1, bruch2)
	run(bruch3, bruch4)
	run(bruch5, bruch6)

}

func run(b1 *bruch.Bruch, b2 *bruch.Bruch) {
	fmt.Println(b1, " + ", b2, " = ", b1.Add(b2))
	fmt.Println(b1, " - ", b2, " = ", b1.Sub(b2))
	fmt.Println(b1, " * ", b2, " = ", b1.Mul(b2))
	result, err := b1.Div(b2)
	if err != nil {
		fmt.Println("Nicht durch 0 teilen!")
	}
	fmt.Println(b1, " / ", b2, " = ", result)
}
