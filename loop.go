package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Print(i)
		fmt.Print(" ")
		i = i + 1
	}
	fmt.Println()

	for j := 10; j > 0; j-- {
		if j%2 == 0 {
			fmt.Printf("Even number %d\n", j)
		} else if j%3 == 0 {
			fmt.Printf("Multiple of 3 = %d \n", j)
		} else {
			fmt.Printf("simple Odd number %d\n",j)
		}
	}

	if (true && true)||(false && true){
		fmt.Println("always true")
	}
}
