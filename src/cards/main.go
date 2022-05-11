package main

import "fmt"

func main() {

	fmt.Println("\n# Creating a new deck")
	cards := newDeck()

	fmt.Println("\n# Splitting the deck in two decks")
	d1, d2 := deal(cards, 3)

	fmt.Println("\n# Printing the decks by their reference function")
	fmt.Println(d1.toString())
	fmt.Println(d2.toString())

	fmt.Println("\n# Saving decks to files")
	d1.saveToFile("d1.txt")
	d2.saveToFile("d2.txt")

	fmt.Println("\n# Reading deck from the file")
	d3, err := newDeckFromFile("d1.txt")
	if err != nil {
		fmt.Println("Error: {}", err)
	} else {
		d3.print()
	}

	fmt.Println("\n# Shuffling Deck")
	fmt.Println("* Before Shuffling")
	fmt.Println(d1.toString())
	fmt.Println("* After Shuffling")
	fmt.Println(d1.toString())
}
