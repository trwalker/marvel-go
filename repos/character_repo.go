package repos

import (
	"github.com/trwalker/marvel-go/models"
)

type CharacterRepo interface {
	GetCharacter(id int) *models.CharacterDetailsModel
}
