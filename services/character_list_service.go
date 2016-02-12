package services

import (
	"github.com/trwalker/marvel-go/models"
)

type CharacterListService interface {
	GetCharacterList() *models.CharacterListModel
}
