package repos

import (
	"github.com/trwalker/marvel-go/models"
)

var CharacterRepoInstance CharacterRepo = initCharacterRepo()

var characterDetailsMap map[int]models.CharacterDetailsModel

type CharacterRepoImpl struct {
}

func initCharacterRepo() CharacterRepo {
	characterDetailsMap = make(map[int]models.CharacterDetailsModel)

	return &CharacterRepoImpl{}
}

func (characterRepo *CharacterRepoImpl) GetCharacter(id int) *models.CharacterDetailsModel {
	return nil
}
