package main

import "fmt"

func printNameAndDepartment(name, department string) {
	fmt.Printf("Name: %s, Department: %s\n", name, department)
}

func main() {
	// Call your functions here
	printNameAndDepartment("Vinay", "IT")

	introduce_punit("Anonymous User")

	fmt.Println("Name: Aravind Kumar")
	fmt.Println("Details:")
	favoriteMoviesAndDialogue()

}
