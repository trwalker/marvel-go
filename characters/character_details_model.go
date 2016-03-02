package characters

type CharacterDetailsModel struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Comics      []string `json:"comics"`
}