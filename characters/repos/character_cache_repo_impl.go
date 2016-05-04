package charrepos

import (
	"github.com/trwalker/marvel-go/characters/models"
	"sync"
)

var CharacterCacheRepoInstance CharacterCacheRepo = &CharacterCacheRepoImpl{
	cache: make(map[int]*charmodels.CharacterModel),
	lock:  &sync.RWMutex{},
}

type CharacterCacheRepoImpl struct {
	cache map[int]*charmodels.CharacterModel
	lock  *sync.RWMutex
}

func (characterCacheRepo *CharacterCacheRepoImpl) Get(id int) (character *charmodels.CharacterModel, found bool) {
	characterCacheRepo.lock.RLock()
	defer characterCacheRepo.lock.RUnlock()

	character, found = characterCacheRepo.cache[id]

	return
}

func (characterCacheRepo *CharacterCacheRepoImpl) Add(character *charmodels.CharacterModel) {
	if character != nil {
		characterCacheRepo.lock.Lock()
		defer characterCacheRepo.lock.Unlock()

		characterCacheRepo.cache[character.Id] = character
	}
}
