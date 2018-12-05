package main

import "fmt"

func main() {
	fmt.Println("hello masud")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("1+1.5 = ", 1.0+1.5)
	fmt.Println("5/3 = ", 5.0/3)
	s := "Good "
	var m string = "Morning"
	fmt.Println(s + m)
	for i:=0;i<len(s);i++ {
		fmt.Print(string(s[i]))
	}

}
