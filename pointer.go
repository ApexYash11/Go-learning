package main

import "fmt"

func change(x *int) {
	*x = 20
}

type book struct {
	title string
	id int
}

func (b *book) setTitle(title string) {
	b.title = title
}

func main(){
	a := 10
	change(&a)
	fmt.Println(a)

	b := book{ "Programming in Go",1}
	b.setTitle("Go Programming")
	fmt.Println(b)
}