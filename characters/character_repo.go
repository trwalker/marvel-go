package characters

type CharacterRepo interface {
	GetCharacter(id int) *CharacterDetailsModel
}
