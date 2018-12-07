package main

import (
	"fmt"
	"time"
)

func sender1(c chan<- string){
	for i:=0;;i++{
		c<-"Hi"
		time.Sleep(time.Second*1)
	}
}
func sender2(c chan<- string){
	for i:=0;;i++{
		c<-"Hello"
		time.Sleep(time.Second*1)
	}
}

func receiver1(c <-chan string){
	for{
		msg:= <-c
		fmt.Println("r1: ",msg)
		time.Sleep(time.Second*1)
	}

}

func receiver2(c <-chan string){
	for{
	msg:= <-c
	fmt.Println("r2: ",msg)
	time.Sleep(time.Second*1)
	}

}


func main() {
	var c chan string =make(chan string)
 	  go sender1(c)
	  go sender2(c)
	  go receiver1(c)
	  go receiver2(c)


	var input string
	fmt.Scanln(&input)
}
