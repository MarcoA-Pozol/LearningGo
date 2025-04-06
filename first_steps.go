// Main method with first package "fmt"
package main

import "fmt" //Formating Package

func main() {
	fmt.Println("This is Go!")
	fmt.Println("And go is really fast!")
	fmt.Println("You cannot stop it!")

	// Variables
	var name string = "Marco Antonio"
	last_name := "Pozol" // Var type autodetection
	const PI = 3.14

	fmt.Println(name)
	fmt.Println(last_name)
	fmt.Println(PI)

	fmt.Println(sumNumbers())
	fmt.Println(getTwoNumbersAverage(458, 540))
}

func sumNumbers() int {
	var numberone = 500
	var numbertwo = 600
	var result = numberone + numbertwo
	return result
}

func getTwoNumbersAverage(a int, b int) int {
	return (a + b) / 2
}
