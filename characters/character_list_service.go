package characters

type CharacterListService interface {
	GetCharacterList(filter string) *CharacterListModel
}
