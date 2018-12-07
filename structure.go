package main

import "fmt"

type Person struct {
	name   string
	age    int
	height float64
}

func make_young(X *Person) {
	(*X).age -= 25
	fmt.Println(*X, ": *")
	fmt.Println(X)
}

func heightIncRate(p Person) float64 {
	return p.height / float64(p.age)
}

//method
func (p *Person) methodM() float64 {
	return p.height / float64(p.age)
}

//embedded type

func (p *Person) dance() {
	fmt.Println(p.name, ": Hi, I can dance")
}

type Man struct {
	Person
	extra_quality string
}

func main() {
	var oneMan Person
	oneMan = Person{"kamol", 23, 176.0}
	twoMan := Person{"masud", 50, 150.0}
	fmt.Println(oneMan)
	fmt.Println(twoMan)
	fmt.Println(oneMan.name, twoMan.name)
	make_young(&twoMan)
	fmt.Println(twoMan)
	fmt.Println(heightIncRate(oneMan))
	fmt.Println(twoMan.methodM())

	Masud := new(Man)
	Masud.name = "Masud"
	Masud.dance()

}
