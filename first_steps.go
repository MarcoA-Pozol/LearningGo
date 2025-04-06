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

	// Lists
	var namesList = []string{"Juan", "Andrea", "Jason", "Jennifer"}

	// Arrays
	var numbersArray [4]int
	numbersArray[0] = 78
	numbersArray[1] = 19
	numbersArray[2] = 22
	numbersArray[3] = 95

	var sizesArray [4]string
	sizesArray[0] = "S"
	sizesArray[1] = "M"
	sizesArray[2] = "L"
	sizesArray[3] = "XL"

	// Maps
	userMap := map[string]string{"username": "MarcoA-Pozol", "age": "23"}
	userMap["email"] = "marcoantoniopozolnarciso@gmail.com"

	// Printing
	fmt.Println(name)
	fmt.Println(last_name)
	fmt.Println(PI)

	fmt.Println(sumNumbers())
	fmt.Println(getTwoNumbersAverage(458, 540))

	fmt.Println(numbersArray)
	fmt.Println(namesList)
	fmt.Println(sizesArray)
	fmt.Println(userMap)
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
