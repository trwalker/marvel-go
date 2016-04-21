package characters

type CharacterDetailsService interface {
	GetCharacter(name string) (characterDetails *CharacterDetailsModel, found bool, err error)
}
