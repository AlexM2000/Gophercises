package deck

type Suit uint8

type Rank uint8

type Card struct {
	Suit 
	Rank
}

const DECKSIZE = 36

const (
	Spade = iota
	Diamond
	Club
	Heart
)

suits := []Suit{Spade, Diamond, Club, Heart}

const (
	Six = iota + 5
	Seven 
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	MinRank = Six
	MaxRank = Ace
)

type lessFunc func(p1, p2 *Card)

type sortDeck struct {
	deck []Card
	less []lessFunc
}

func generateColor(deck []Card, color string) []Card { 
}

func New() []Card {
	newDeck := make([]Card, DECKSIZE, DECKSIZE)
	for _, suit := range suits {
		for value := MinRank; value <= MaxRank; value++ {
			newDeck = append(newDeck, Card{Suit: suit, Rank: value})
		}
	}
	return newDeck
}

func (d *sortDeck) Len() int {
	return len(d.deck)
}

func (d *sortDeck) Swap(i, j int) {
	d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
}

func orderedBy(less ...lessFunc) *sortDeck {
	return &sortDeck{
		less: less,
	}
}

func (d *sortDeck) Less(i, j int) bool {
	p, q := &d.deck[i], &d.deck[j]

}
