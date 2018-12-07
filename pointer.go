package main

import "fmt"

func sqare(x *int) {
	*x *= *x
}

func main() {
	x := 5
	sqare(&x)
	fmt.Println(x)
	ad := new(int)
	*ad = 5
	fmt.Println(*ad)

}
