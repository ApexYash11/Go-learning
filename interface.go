package main

import "fmt"

type shape interface {
	getperimeter() uint
}

type square struct {
	width uint
}

func (s square) getperimeter() uint {
	return 4 * s.width
}

type triangle struct {
	a uint
	b uint
	c uint
}

func (t triangle) getperimeter() uint {
	return t.a + t.b + t.c
}

func (t triangle) getsides() (uint, uint, uint) {
	return t.a, t.b, t.c
}

func main() {
 var s [] shape = []shape{triangle{a: 3, b: 4, c: 5}, square{width: 2}}
 for _, sh := range s {
  fmt.Println(sh.getperimeter())
 }
 t := triangle{a: 3, b: 4, c: 5}
 fmt.Println(t.getsides())
 

}