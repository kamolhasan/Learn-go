package main

import "fmt"

func main() {
	var x [10]int
	for i := 0; i < 10; i++ {
		x[i] = i + i
	}
	fmt.Println(x)

	var total float64
	for i := 0; i < 10; i++ {
		total += float64(x[i])
	}
	total /= float64(len(x))

	fmt.Println(total)

	for _, value := range x {
		fmt.Println(value)
	}

	newx := []float64{1, 2, 2, 3, 33, 9}
	for i := 0; i < len(newx); i++ {
		fmt.Printf("%.2f ", newx[i])
	}
}
