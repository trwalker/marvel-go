package charrepos

import (
	"github.com/trwalker/marvel-go/characters/models"
)

type CharacterCacheRepo interface {
	Get(id int) (characterDetails *charmodels.CharacterModel, found bool)
	Add(characterDetails *charmodels.CharacterModel)
}
