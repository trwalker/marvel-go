package characters

import (
	"sync"
)

var CharacterCacheRepoInstance CharacterCacheRepo = &CharacterCacheRepoImpl {
	cache: make(map[int]*CharacterDetailsModel),
	lock: &sync.RWMutex{},
}

type CharacterCacheRepoImpl struct {
	cache map[int]*CharacterDetailsModel
	lock *sync.RWMutex
}

func (characterCacheRepo *CharacterCacheRepoImpl) Get(id int) (characterDetails *CharacterDetailsModel, found bool) {
	characterCacheRepo.lock.RLock()
	defer characterCacheRepo.lock.RUnlock()

	characterDetails, found = characterCacheRepo.cache[id]

	return
}

func (characterCacheRepo *CharacterCacheRepoImpl) Add(characterDetails *CharacterDetailsModel) {
	if characterDetails != nil {
		characterCacheRepo.lock.Lock()
		defer characterCacheRepo.lock.Unlock()

		characterCacheRepo.cache[characterDetails.Id] = characterDetails
	}
}