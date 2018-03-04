package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(initDeck())
}

func initDeck() []string {
	var suits [4]string
	suits[0] = "Hearts"
	suits[1] = "Diamonds"
	suits[2] = "Spades"
	suits[3] = "Clubs"

	cards := []string{}
	for _, suit := range suits {
		for i := 1; i <= 13; i++ {

			var cardName = ""
			switch (i) {
			case 1:
				cardName = "Ace"
			case 11:
				cardName = "Jack"
			case 12:
				cardName = "Queen"
			case 13:
				cardName = "King"
			default:
				cardName = strconv.Itoa(i)
			}

			cards = append(cards, fmt.Sprintf("%s of %s", cardName, suit))
		}
	}

	return cards
}
