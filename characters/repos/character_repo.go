package charrepos

import (
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/characters/models"
)

type CharacterRepo interface {
	GetCharacter(characterId int, credentials *auth.CredentialsModel) (character *charmodels.CharacterModel, found bool, err error)
}
