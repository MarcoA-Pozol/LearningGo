package main
import "fmt"

// Main function
func main() {
	/* We do not have Classes in Go, so we use Structures instead (struct)*/
	
	
	
	// Defining Person structs
	personOne := Person{Name: "Joseph", Age: 35, Country: "USA", Occupation: "Medic"}
	personTwo := Person{Name: "Jennifer", Age: 21, Country: "Canada", Occupation: "Bartender"}
	personThree := Person{Name: "Marco", Age: 23, Country: "Mexico", Occupation: "Software Developer"}
	personFour := Person{Name: "Camila", Age: 32, Country: "USA", Occupation: "English Teacher"}

	// Accesing Person structs data
	fmt.Println(personOne)
	fmt.Println(personTwo)
	fmt.Println(personThree.Name)
	fmt.Println(personFour.Occupation)
	
	personOne.Greet(personFour)
	personThree.MeetPerson(personTwo)
}

// Person struct
type Person struct {
	Name string
	Age int
	Country string
	Occupation string
}

// Adding behavior to Person structs with methods
func (p Person) Greet(other_person Person) {
	fmt.Println("Hi, my name is", p.Name, "I am saying hi to", other_person.Name)
}

func (p Person) MeetPerson(other_person Person) {
	fmt.Println("Nice to meet you", other_person.Name, "my name´s", p.Name)
}