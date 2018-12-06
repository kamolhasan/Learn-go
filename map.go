package main

import (
	"fmt"
	"sort"
)

func main() {
	x := make(map[string]int)
	x["k"] = 1
	x["a"] = 2
	x["m"] = 3
	x["o"] = 4
	x["l"] = 5

	delete(x, "k")

	fmt.Println(x)
	for k, value := range x {
		println(k, value)
	}
	var tmp []string
	for k, _ := range x {
		tmp=append(tmp,k)
	}
	sort.Strings(tmp)

	for _,value:=range tmp{
		fmt.Println("here: ",x[value])
	}

}
