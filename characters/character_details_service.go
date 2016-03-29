package characters

type CharacterDetailsService interface {
	GetCharacter(name string) *CharacterDetailsModel
}
