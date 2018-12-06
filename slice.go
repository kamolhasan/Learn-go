package main

import "fmt"

func main() {
	var x []float64
	x = append(x, 10)
	fmt.Println(x)

	y := make([]int, 3)
	fmt.Println(len(y))
	fmt.Println(y)
	y = append(y, 111)
	fmt.Println(y)

	z := make([]int64, 2)

	for i := 1; i <= 5; i++ {

		z = append(z, int64(i))
		fmt.Println(z)
		fmt.Println(len(z))

		z=append(z[:1],z[2:]...)

		fmt.Println(z)
		fmt.Println(len(z))
	}

	ar:= [5]int{1,2,2,3,5}
	fmt.Println(ar)
	ns:=ar[1:4]
	fmt.Println(ns)

}
