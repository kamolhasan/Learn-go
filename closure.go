package main

import "fmt"

func main() {

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
	a := 1
	incriment := func() int {
		a++
		return a
	}
	fmt.Println(incriment())
	fmt.Println(incriment())
	fmt.Println(incriment())

	//closure (variable+function)
	nxtEven := evenGen()
	fmt.Println(nxtEven())
	fmt.Println(nxtEven())
	fmt.Println(nxtEven())
	fmt.Println(nxtEven())
}

func evenGen() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
