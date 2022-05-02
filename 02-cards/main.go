// Defines package as an executable.
package main

// Defines package as a library.
// package my-lib

func main() {
	deck := newDeck()
	deck.shuffle()
	hand, deck := deal(deck, 5)
	hand.print()

	// write(deck, "deck")
	// deck := read("deck")
}
