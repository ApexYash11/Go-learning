package main

import "fmt"

type Person struct { // struct
	Name      string
	Age       int
	FavSports []sport
}

type sport struct {
	Name      string
	age       int
	favsports []sport
}

func getname(p Person) string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func (p Person) getname1() string {
	return p.Name
}

func (p Person) setname(newName string) {
	p.Name = newName
	fmt.Println(p)
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	p1 := Person{Age: 15}
	p1.Name = "Bob"
	fmt.Println(p1)
	fmt.Println(p.Name, p.Age)
	value := p.getname1()
	fmt.Println(value)

	var p2 Person = Person{Name: "Charlie", Age: 25}
	p2.setname("David")
	fmt.Println(p2)

	p3 := Person{Name: "Eve", Age: 28, FavSports: []sport{{Name: "Soccer", age: 10}, {Name: "Basketball", age: 5}}}
	fmt.Println(p3)
}
