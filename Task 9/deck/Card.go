package deck

type Card struct {
	color string
	value int
}

type deck []Card

func New() []Card {
	return []Card{}
}

func (d deck) Len() int {
	return len(d)
}

func (d deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d deck) Less(i, j int) bool {
	return d[i].value < d[j].value
}
