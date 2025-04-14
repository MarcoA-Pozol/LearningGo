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
	
	// Defining Vehicle 
	vehicle_one := Vehicle{Name:"Ford Raptor 2025", Horsepower:690, VType:"Ground", Color:"Orange"}
	vehicle_two := Vehicle{Name:"Boing B-97 AA", Horsepower:690, VType:"Air", Color:"White-Gray"}
	vehicle_three := Vehicle{Name:"Dgiant Vessel G200", Horsepower:850, VType:"Sea", Color:"Blue-White"}
	
	fmt.Println(vehicle_one.StartEngine())
	fmt.Println(vehicle_two.StartEngine())
	fmt.Println(vehicle_three.StartEngine())
	
	fmt.Println(vehicle_one.MoveTo(12500, "Minessota"))
	fmt.Println(vehicle_two.MoveTo(135000, "Houston, Texas"))
	fmt.Println(vehicle_three.MoveTo(96600, "Pacific Ocean"))
	
	fmt.Println(vehicle_one.StopEngine())
	fmt.Println(vehicle_one.StopEngine())
	fmt.Println(vehicle_one.StopEngine())
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


// Vehicle struct
type Vehicle struct {
	Name string
	Horsepower int
	VType string
	Color string
}

// Interface for "Vehicle"
type Vehicle interface {
	StartEngine(Vehicle) string
	MoveTo(Vehicle, direction string, meters int) string
	StopEngine(Vehicle) string
}

func (v Vehicle) StartEngine() string {
	result = fmt.Sprintf("%s has started its engine.", v.Name)
	return result
}

func (v Vehicle) MoveTo(direction string, meters string) string {
	result = fmt.Sprintf("This vehicle '%s' has moved %s meters to %s.", v.Name, meters, direction)
	return result
}

func (v Vehicle) StopEngine() string {
	result = fmt.Sprintf("%s has stop its engine.", v.Name)
	return result
}