package charservices

import (
	"github.com/trwalker/marvel-go/characters/models"
)

type CharacterListService interface {
	GetCharacterList() *charmodels.CharacterListModel
}
