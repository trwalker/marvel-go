package characters

type CharacterService interface {
	GetCharacter(name string) (character *CharacterModel, found bool, err error)
}
