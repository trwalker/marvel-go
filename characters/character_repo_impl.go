package characters

import ()

var CharacterRepoInstance CharacterRepo = &CharacterRepoImpl{}

var characterDetailsMap map[int]CharacterDetailsModel = make(map[int]CharacterDetailsModel)

type CharacterRepoImpl struct {
}

func (characterRepo *CharacterRepoImpl) GetCharacter(id int) *CharacterDetailsModel {
	return nil
}
