package main

import "fmt"

type Car struct {
	Color string
}

func (c *Car) findColor() {
	fmt.Println("Car's Color is : ", c.Color)
}

func (h *House) findColor() {

	fmt.Println("House's color is: ", h.Color)

}

type Resource interface {
	findColor()
}

func PrintColor(X ...Resource) {

	for _, value := range X {
		value.findColor()
	}
}

type House struct {
	Color string
}

func main() {
	A := House{"Blue"}
	B := Car{"red"}
	C := Car{"white"}
	D := House{"Off-White"}
	PrintColor(&A, &B, &C, &D)

}
