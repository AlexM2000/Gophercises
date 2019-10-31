package types

type Story map[string]Adventure

type Adventure struct {
	Title       string    `json:"title"`
	StoryText   []string  `json:"story"`
	OptionsList []Options `json:"options"`
}

type Options struct {
	Text     string `json:"text"`
	StoryArc string `json:"arc"`
}
