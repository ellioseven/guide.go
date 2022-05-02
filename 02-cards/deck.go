package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type card struct {
	suit  string
	value string
}

type deck []string

func newDeck() deck {
	deck := deck{}

	// atts are assigned default value if not specified.
	spac := card{
		suit:  "Spades",
		value: "Ace",
	}

	// print object.
	// fmt.Printf("%+v", spac)

	suits := []string{
		"Spades",
		"Clubs",
		"Diamonds",
		"Hearts",
	}

	values := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, value+" of "+suit)
		}
	}

	return deck
}

// receiver.
func (_deck deck) print() {
	fmt.Println(_deck.toString())
}

func deal(_deck deck, size int) (deck, deck) {
	return _deck[:size], _deck[size:]
}

func (_deck deck) write(filename string) error {
	contents := _deck.toString()
	return ioutil.WriteFile(filename, []byte(contents), 0666)
}

func read(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	parts := strings.Split(string(bs), ",")
	return deck(parts)
}

func (_deck deck) toString() string {
	return strings.Join([]string(_deck), ",")
}

func (_deck deck) shuffle() {
	randSeed := time.Now().Unix()
	randSource := rand.NewSource(randSeed)
	randGen := rand.New(randSource)

	for i := range _deck {
		size := len(_deck) - 1
		posNew := randGen.Intn(size)

		// swap elements.
		_deck[i], _deck[posNew] = _deck[posNew], _deck[i]
	}
}
