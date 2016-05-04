package charservices

import (
	"github.com/trwalker/marvel-go/characters/models"
)

type CharacterService interface {
	GetCharacter(name string) (character *charmodels.CharacterModel, found bool, err error)
}
