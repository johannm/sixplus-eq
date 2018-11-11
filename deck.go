package main

import (
	"math/rand"
)

type card struct {
	rank int // 4, 5, 6, 7, 8, 9, 10, 11, 12 <=> 6, 7, 8, 9, ten, jack, queen, king, ace
	suit int // 0 == club, 1 == diamond, 2 == heart, 3 == spade
}

type hand struct {
	cards []card
}

func createDeck() []card {
	var deck [36]card
	for i := 0; i < 36; i++ {
		rank := i % 9 + 4
		suit := i % 4
		deck[i] = card{rank: rank, suit: suit}
	}
	return deck[:]
}

func shuffle(deck []card, r *rand.Rand) {
	for i := 0; i < 35; i++ {
		j := r.Intn(36-i) + i
		tmp := deck[j]
		deck[j] = deck[i]
		deck[i] = tmp
	}
}

func remove(removeCard card, deck []card) []card {
	var newDeck []card
	for i := 0; i < 36; i++ {
		if removeCard == deck[i] {
			newDeck = append(deck[:i], deck[i+1:]...)
			return newDeck
		}
	}
	return deck
}
