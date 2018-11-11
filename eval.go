package main

import (
	"sort"
)

/*
	Ranks:
		Straight Flush 8
		Quads 7
		Flush 6
		Boat 5
		Trips 4
		Straight 3
		2-pair 2
		1-pair 1
		Hi-card 0
*/
func rankHand(cards []card) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].rank > cards[j].rank
	})
	isFlush := true
	for _, card := range cards {
		if card.suit != cards[0].suit {
			isFlush = false
		}
	}
	isStraight := true
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].rank != cards[i+1].rank+1 {
			isStraight = false
		}
	}
	// Handle special case for A-6-7-8-9 straight
	if cards[0].rank == 12 && cards[1].rank == 7 && cards[2].rank == 6 && cards[3].rank == 5 && cards[4].rank == 4 {
		isStraight = true
	}

	if isFlush && isStraight {
		return 8
	}
	if isFlush {
		return 6
	}
	if isStraight {
		return 3
	}

	m := make(map[int]int)
	for _, card := range cards {
		n, prs := m[card.rank]
		if prs {
			m[card.rank] = n + 1
		} else {
			m[card.rank] = 1
		}
	}
	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return 7
			}
		}
		return 5
	}
	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return 4
			}
		}
		return 2
	}
	if len(m) == 4 {
		return 1
	}
	return 0
}

func compare(hand1 []card, hand2 []card) int {
	r1 := rankHand(hand1)
	r2 := rankHand(hand2)
	if r1 != r2 {
		if r1 > r2 {
			return 1
		} else {
			return -1
		}
	}

	sort.Slice(hand1, func(i, j int) bool {
		return hand1[i].rank > hand1[j].rank
	})
	sort.Slice(hand2, func(i, j int) bool {
		return hand2[i].rank > hand2[j].rank
	})

	m1 := make(map[int]int)
	for _, card := range hand1 {
		n, prs := m1[card.rank]
		if prs {
			m1[card.rank] = n + 1
		} else {
			m1[card.rank] = 1
		}
	}
	m2 := make(map[int]int)
	for _, card := range hand2 {
		n, prs := m2[card.rank]
		if prs {
			m2[card.rank] = n + 1
		} else {
			m2[card.rank] = 1
		}
	}

	switch r1 {
	case 8, 3:
		// Straight-Flush, Straight
		if hand1[4].rank == 4 || hand2[4].rank == 4 { /* Handle special case for A-2-3-4-5 straight */
			if hand1[4].rank != 4 {
				return 1
			} else if hand2[4].rank != 4 {
				return -1
			} else {
				return 0
			}
		}
		for i, card := range hand1 {
			if card.rank > hand2[i].rank {
				return 1
			} else if card.rank < hand2[i].rank {
				return -1
			}
		}
		return 0
	case 6, 0:
		// Flush, High-card
		for i, card := range hand1 {
			if card.rank > hand2[i].rank {
				return 1
			} else if card.rank < hand2[i].rank {
				return -1
			}
		}
	case 7:
		// Quads
		q1, k1, q2, k2 := 0, 0, 0, 0
		for k, v := range m1 {
			if v == 4 {
				q1 = k
			} else if v == 1 {
				k1 = k
			}
		}
		for k, v := range m2 {
			if v == 4 {
				q2 = k
			} else if v == 1 {
				k2 = k
			}
		}
		if q1 > q2 {
			return 1
		} else if q1 < q2 {
			return -1
		} else if k1 > k2 {
			return 1
		} else if k1 < k2 {
			return -1
		}
		return 0
	case 5:
		// Boat
		t1, o1, t2, o2 := 0, 0, 0, 0
		for k, v := range m1 {
			if v == 3 {
				t1 = k
			} else if v == 2 {
				o1 = k
			}
		}
		for k, v := range m2 {
			if v == 3 {
				t2 = k
			} else if v == 2 {
				o2 = k
			}
		}
		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		} else if o1 > o2 {
			return 1
		} else if o1 < o2 {
			return -1
		}
		return 0
	case 4:
		// Trips
		t1, t2 := 0, 0
		var k1 []int
		var k2 []int
		for k, v := range m1 {
			if v == 3 {
				t1 = k
			} else if v == 1 {
				k1 = append(k1, k)
			}
		}
		for k, v := range m2 {
			if v == 3 {
				t2 = k
			} else if v == 1 {
				k2 = append(k2, k)
			}
		}
		sort.Slice(k1, func(i, j int) bool {
			return k1[i] > k1[j]
		})
		sort.Slice(k2, func(i, j int) bool {
			return k2[i] > k2[j]
		})
		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		} else {
			for i, v := range k1 {
				if v > k2[i] {
					return 1
				} else if v < k2[i] {
					return -1
				}
			}
		}
		return 0
	case 2:
		// Two Pair
		var p1 []int
		var p2 []int
		k1, k2 := 0, 0
		for k, v := range m1 {
			if v == 2 {
				p1 = append(p1, k)
			} else if v == 1 {
				k1 = k
			}
		}
		for k, v := range m2 {
			if v == 2 {
				p2 = append(p2, k)
			} else if v == 1 {
				k2 = k
			}
		}
		sort.Slice(p1, func(i, j int) bool {
			return p1[i] > p1[j]
		})
		sort.Slice(p2, func(i, j int) bool {
			return p2[i] > p2[j]
		})
		if p1[0] > p2[0] {
			return 1
		} else if p1[0] < p2[0] {
			return -1
		} else if p1[1] > p2[1] {
			return 1
		} else if p1[1] < p2[1] {
			return -1
		} else if k1 > k2 {
			return 1
		} else if k1 < k2 {
			return -1
		}
		return 0
	case 1:
		// One Pair
		p1, p2 := 0, 0
		var k1 []int
		var k2 []int
		for k, v := range m1 {
			if v == 2 {
				p1 = k
			} else if v == 1 {
				k1 = append(k1, k)
			}
		}
		for k, v := range m2 {
			if v == 2 {
				p2 = k
			} else if v == 1 {
				k2 = append(k2, k)
			}
		}
		sort.Slice(k1, func(i, j int) bool {
			return k1[i] > k1[j]
		})
		sort.Slice(k2, func(i, j int) bool {
			return k2[i] > k2[j]
		})
		if p1 > p2 {
			return 1
		} else if p1 < p2 {
			return -1
		} else {
			for i, v := range k1 {
				if v > k2[i] {
					return 1
				} else if v < k2[i] {
					return -1
				}
			}
		}
		return 0
	}

	return 0
}
