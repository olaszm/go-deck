package main

import (
	"fmt"

	"github.com/olaszm/go-deck/card"
)

func main() {
	deck := card.New(card.AddJokers(5), card.WithShuffleDeck(), card.FilterOutCard(0))

	for _, c := range *deck {
		fmt.Println(c.String())
	}

	fmt.Println(len(*deck))
}
