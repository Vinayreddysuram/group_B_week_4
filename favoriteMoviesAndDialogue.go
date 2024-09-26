package main

import "fmt"

func favoriteMoviesAndDialogue() {
	// Favorite books
	movies := []string{
		"Batman: The Dark Knight",
		"Pursuit of Happiness",
		"Interstellar",
		"12 Angry Men",
	}

	// Favorite dialogue
	dialogue := "The world is your oyster. It’s up to you to find the pearls."

	// Why the dialogue is meaningful
	reason := "it encourages personal empowerment and responsibility, reminding us that while the world is full of opportunities, it’s up to us to actively seek and seize them for success and fulfillment."

	// Fun fact about the group member
	funFact := "I love the christian Bale Batman Trilology"

	// Displaying information
	fmt.Println("Favorite Movies:")
	for i, movie := range movies {
		fmt.Printf("%d. %s\n", i+1, movie) // Prints movie number and title
	}

	fmt.Println("\nFavorite dialogue:")
	fmt.Println("\"" + dialogue + "\"")

	fmt.Println("\nWhy this dialogue is meaningful:")
	fmt.Println(reason)

	fmt.Println("\nFun Fact:")
	fmt.Println(funFact)
}
