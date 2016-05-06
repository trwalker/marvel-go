package characters

import (
	"github.com/trwalker/marvel-go/auth"
)

type CharacterRepo interface {
	GetCharacter(characterId int, credentials *auth.CredentialsModel) (character *CharacterModel, found bool, err error)
}
