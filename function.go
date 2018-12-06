package main

import "fmt"

func main() {
	ar := []float64{2, 3, 2, 3, 2, 3}
	average := avg(ar)
	fmt.Println(average)
	a, b, c := callF()
	fmt.Println(a, b, c)

	fmt.Println(totalSum(1, 2, 3, 4, 5))

	iar := []int{12, 12, 12}
	fmt.Println(totalSum(iar...))
	//sum from index 2 to end
	fmt.Println(totalSum(iar[2:]...))
}

func avg(nar []float64) float64 {
	var ans float64 = 0
	for i := 0; i < len(nar); i++ {
		ans += nar[i]
	}

	ans /= float64(len(nar))
	return ans
}

func callF() (a int, b int, c int) {
	a = 1
	b = 8
	c = 3
	return
}

func totalSum(ar ...int) int {
	sum := 0
	for _, value := range ar {
		sum += value
	}
	return sum
}
