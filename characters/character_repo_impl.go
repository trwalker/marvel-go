package characters

import ()

var CharacterRepoInstance CharacterRepo = &CharacterRepoImpl{}

var characterMap map[int]CharacterModel = make(map[int]CharacterModel)

type CharacterRepoImpl struct {
}

func (characterRepo *CharacterRepoImpl) GetCharacter(id int) *CharacterModel {
	return nil
}
