package characters

import (
	"sync"
)

var CharacterMapRepoInstance CharacterMapRepo = &CharacterMapRepoImpl{
	characterMap: make(map[string]int),
	lock:         &sync.Mutex{},
}

type CharacterMapRepoImpl struct {
	characterMap map[string]int
	lock         *sync.Mutex
}

func (characterMapRepo *CharacterMapRepoImpl) GetCharacterMap() map[string]int {
	if len(characterMapRepo.characterMap) == 0 {
		characterMapRepo.lock.Lock()
		defer characterMapRepo.lock.Unlock()

		if len(characterMapRepo.characterMap) == 0 {
			buildCharacterMap(characterMapRepo)
		}
	}

	return characterMapRepo.characterMap
}

func buildCharacterMap(characterMapRepo *CharacterMapRepoImpl) {
	characterMapRepo.characterMap["spider-man"] = 1009610
	characterMapRepo.characterMap["hulk"] = 1009351
	characterMapRepo.characterMap["captain-america"] = 1009220
	characterMapRepo.characterMap["iron-man"] = 1009368
	characterMapRepo.characterMap["thor"] = 1009664
	characterMapRepo.characterMap["wolverine"] = 1009718
	characterMapRepo.characterMap["storm"] = 1009629
	characterMapRepo.characterMap["jean-grey"] = 1009496
	characterMapRepo.characterMap["gambit"] = 1009313
	characterMapRepo.characterMap["cyclops"] = 1009257
	characterMapRepo.characterMap["beast"] = 1009175
}
