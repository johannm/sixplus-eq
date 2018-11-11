package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func parseHand(s string) []card {
	return []card{parseCard(s[0:2]), parseCard(s[2:4])}
}

func parseBoard(s string) []card {
	var b []card
	for len(s) > 0 {
		b = append(b, parseCard(s[0:2]))
		s = s[2:]		
	}
	return b
}

func parseCard(s string) card {
	rank, suit := s[0], s[1]
	var c card
	switch rank {
	case 'A':
		c.rank = 12
	case 'K':
		c.rank = 11
	case 'Q':
		c.rank = 10
	case 'J':
		c.rank = 9
	case 'T':
		c.rank = 8
	case '9':
		c.rank = 7
	case '8':
		c.rank = 6
	case '7':
		c.rank = 5
	case '6':
		c.rank = 4
	}
	switch suit {
	case 's':
		c.suit = 3
	case 'h':
		c.suit = 2
	case 'd':
		c.suit = 1
	case 'c':
		c.suit = 0
	}
	return c
}

func combos(cards []card, length int, startPos int, result []card, comboList *[]hand) {
	if length == 0 {
		h := make([]card, len(result))
		copy(h, result)
		*comboList = append(*comboList, hand{cards: h})
		return
	}
	for i := startPos; i <= len(cards)-length; i++ {
		result[len(result)-length] = cards[i]
		combos(cards, length-1, i+1, result, comboList)
	}
}

func main() {
	fmt.Printf("Hold-em Hand Equity Calculator\n")
	for i, arg := range os.Args[1:] {
		fmt.Printf("arg %d: %v\n", i, arg)
	}
	hand1 := parseHand(os.Args[1])
	hand2 := parseHand(os.Args[2])
	var board []card
	if len(os.Args) > 3 {
		board = parseBoard(os.Args[3])
		fmt.Printf("board: %v\n", board)
	}
	fmt.Printf("hand1: %v, hand2: %v\n", hand1, hand2)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	n := 100000
	win, lose, draw := 0, 0, 0
	for i := 0; i < n; i++ {
		deck := createDeck()
		shuffle(deck, r1)
		for _, c := range hand1 {
			deck = remove(c, deck)
		}
		for _, c := range hand2 {
			deck = remove(c, deck)
		}
		dealtBoard := append(board, deck[0:5-len(board)]...)

		cards1 := append(hand1, dealtBoard...)
		var result [5]card
		var comboList1 []hand
		combos(cards1, 5, 0, result[:], &comboList1)

		cards2 := append(hand2, dealtBoard...)
		var result2 [5]card
		var comboList2 []hand
		combos(cards2, 5, 0, result2[:], &comboList2)

		maxhand1 := comboList1[0]
		for _, h := range comboList1 {
			if compare(maxhand1.cards, h.cards) < 0 {
				maxhand1 = h
			}
		}
		maxhand2 := comboList2[0]
		for _, h := range comboList2 {
			if compare(maxhand2.cards, h.cards) < 0 {
				maxhand2 = h
			}
		}

		won := compare(maxhand1.cards, maxhand2.cards)
		if won > 0 {
			win++
		} else if won < 0 {
			lose++
		} else {
			draw++
		}
	}

	fmt.Printf("n: %v, win: %v, lose: %v, draw: %v\n", n, float64(win)/float64(n), float64(lose)/float64(n), float64(draw)/float64(n))
	fmt.Printf("equity: %v\n", float64(win)/float64(win+lose))
}
