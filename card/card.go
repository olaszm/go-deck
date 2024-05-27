package card

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Suite int

const (
	Spade Suite = iota
	Heart
	Diamond
	Club
	Joker
)

func (s Suite) String() string {
	return [...]string{"Spades", "Hearts", "Diamonds", "Clubs", "Joker"}[s]
}

type Rank int

const (
	Ace Rank = iota + 1
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	Kings
)

func (r Rank) String() string {
	var ret string
	switch r {
	case 1:
		ret = "Ace"
	case 11:
		ret = "Jack"
	case 12:
		ret = "Queen"
	case 13:
		ret = "King"
	default:
		ret = strconv.Itoa(int(r))
	}

	return ret
}

type Card struct {
	Suite Suite
	Rank  Rank
}

func (c Card) String() string {
	if c.Suite == 5 {
		return "Joker"
	}

	return fmt.Sprintf("%s - %s", c.Suite, c.Rank)
}

type Opts func(Deck) Deck

func WithShuffleDeck() Opts {
	return func(d Deck) Deck {
		rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
		return d
	}
}

func AddJokers(count int) Opts {
	return func(d Deck) Deck {
		for i := 0; i < count; i++ {
			d = append(d, Card{Joker, 0})
		}
		return d
	}
}

func FilterOutCard(Rank Rank) Opts {
	return func(d Deck) Deck {
		b := d[:0] // This trick creates an empty slice with the backing array's capacity and length
		for _, c := range d {
			if c.Rank != Rank {
				b = append(b, c)
			}
		}
		return b
	}
}

func BuildMultiple(count int) Opts {
	return func(d Deck) Deck {
		for i := 1; i < count; i++ {
			newDeck := New()
			d = append(d, *newDeck...)
		}
		return d
	}
}

type Deck []Card

func SortReverseValue() Opts {
	return func(d Deck) Deck {
		sort.SliceStable(d, func(i, j int) bool { return d[i].Rank > d[j].Rank })
		return d
	}
}

func defaultSort(d Deck) Deck {
	sort.SliceStable(d, func(i, j int) bool { return d[i].Suite < d[j].Suite })
	return d
}

func New(opts ...Opts) *Deck {
	var deck Deck

	baseSuits := []Suite{Club, Diamond, Heart, Spade}
	for _, v := range baseSuits {
		for i := 1; i <= 13; i++ {
			deck = append(deck, Card{v, Rank(i)})
		}
	}

	deck = defaultSort(deck)

	for _, opt := range opts {
		deck = opt(deck)
	}

	return &deck
}
