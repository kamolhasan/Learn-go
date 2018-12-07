package main

import "fmt"

var x int = 100

func main() {
	var s string
	s = "my name is kamol"
	fmt.Println(x)
	f(&s)

}

func f(ss *string) {

	fmt.Println(*ss)
}
