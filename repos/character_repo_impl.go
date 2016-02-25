package repos

import (
	"github.com/trwalker/marvel-go/models"
)

var CharacterRepoInstance CharacterRepo = &CharacterRepoImpl{}

var characterDetailsMap map[int]models.CharacterDetailsModel = make(map[int]models.CharacterDetailsModel)

type CharacterRepoImpl struct {
}

func (characterRepo *CharacterRepoImpl) GetCharacter(id int) *models.CharacterDetailsModel {
	return nil
}
