package repos

import (
	"github.com/trwalker/marvel-go/models"
)

type CharacterMapRepo interface {
	GetCharacterMap() map[string]*models.CharacterModel
}
