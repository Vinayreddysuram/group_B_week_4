package main

import "fmt"

func printNameAndDepartment(name, department string) {
	fmt.Printf("Name: %s, Department: %s\n", name, department)
}

func concatenateMyString(str1, str2 string) {
	fmt.Printf("\n\n\nThis is the function by Neha. It concatenates 2 strings %s%s", str1, str2)
}

func main() {
	// Call your functions here
	printNameAndDepartment("Vinay", "IT")

	introduce_punit("Anonymous User")

	fmt.Println("Name: Aravind Kumar")
	fmt.Println("Details:")
	favoriteMoviesAndDialogue()

	concatenateMyString("go", "train")

	input := "Hello, Go World!"

    uppercased := toUpperCase(input)

    fmt.Printf("Original: %s\n", input)
    fmt.Printf("Uppercase: %s\n", uppercased)
}
