package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Defining Person structs
	person_one := Person{Name: "Joseph", Age: 29, City: "Oklahoma"}
	person_two := Person{Name: "Helen", Age: 32, City: "Luxenburg"}
	person_three := Person{Name: "Louise", Age: 37, City: "Los Angeles"}

	fmt.Println(person_one.AngryGreet(person_three))
	fmt.Println(person_three.HappyGreet(person_two))
}

// Person struct
type Person struct {
	Name string
	Age  int
	City string
}

// Defining an Interface with two methods (HappyGreet(), AngryGreet())
type Greeter interface {
	HappyGreet(Person) string
	AngryGreet(Person) string
}

// Assigning Interface's methods functionality
func (p Person) HappyGreet(r Person) string {
	greet := fmt.Sprintf("Hello, my name is %s, nice to meet you %s!!! %s", p.Name, r.Name, ":)")
	return greet
}

func (p Person) AngryGreet(r Person) string {
	angry_greetings := []string{fmt.Sprintf("Hi...I am %s, does it matter?", p.Name), fmt.Sprintf("Hi %s...", r.Name)}

	randomIndex := rand.Intn(len(angry_greetings))
	greet := angry_greetings[randomIndex]
	return greet
}
