package main

import "fmt"

func PrintNameAndDepartment(name, department string) {
	fmt.Printf("Name: %s, Department: %s\n", name, department)
}

func main() {
	// Call your functions here
	PrintNameAndDepartment("Vinay", "IT")
}
