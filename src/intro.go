package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("hello world\n")

	// types declare
	var k int             // zero - default
	var i int = 10        // explicit
	const hello = "Hello" //
	b := true             // short -- if has already exist, it is just a assign
	fmt.Println(k, i, hello, b)

	const a = 5
	c := 5
	fmt.Printf("%v\n", a*math.Pi)
	fmt.Printf("%v\n", float64((c))*math.Pi)

	// for loop .
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			fmt.Print("Even ")
		} else {
			fmt.Print("Odd ")
		}
	}
	//while loop .
	n := 0
	for n < 5 {
		fmt.Print(n)
		n++
	}
	fmt.Print("\n")
	//-----data structure---
	//fixed
	var e [5]int
	for _, v := range e {
		fmt.Println("%d\n", v)
	}

	//maps
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 8
	delete(m, "k2")
	//check
	val, ok := m["k2"]
	fmt.Println("1. val:", val, "| ok:", ok)
	q := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("2.map: ", q)

	//slices - dynamic length
	s := make([]string, 2) //zero default
	s[0] = "a"
	s[1] = "b"
	s = append(s, "c", "d", "e")
	fmt.Println("1.len:", len(s), "| cap:", cap(s), "| s:", s)
	//capacity vs len, mostly not use cap
	t := s[2:4] //from 2 to 3
	fmt.Println("1.len:", len(t), "| cap:", cap(t), "| t:", t)
	//copy
	p := make([]string, len(s))
	copy(p, s)

	//---functions---
	fmt.Println("plus:", plus(3, 4))
	//defer: run ASAP after return

	//---struct interface---
	/*
		Structs are typed collection of fields
		Interfaces are a collection of method signatures
	*/
	r := rect{3, 4}
	measure(r)

	x := box{r, "red"}
	fmt.Println(x.color, " ", x.showcolor())

	//---IO---
	//os.ReadFile()

}

// Structs are typed collection of fields
type rect struct {
	width, height float64
}

// Interfaces are a collection of method signatures
type geometry interface {
	area() float64
	perim() float64
}

// implement
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// print
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// interfaces are composable
type object interface {
	geometry
	showcolor() string
}

// structs are composable
type box struct {
	rect
	color string
}

func (b box) showcolor() string {
	return b.color
}

// func process(r rect)
func plus(a /*int*/, b int) int {
	return a + b
}
