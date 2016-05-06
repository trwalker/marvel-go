package characters

import (
	"sync"
)

var CharacterCacheRepoInstance CharacterCacheRepo = NewCharacterCacheRepo()

type characterCacheRepoImpl struct {
	cache map[int]*CharacterModel
	lock  *sync.RWMutex
}

func NewCharacterCacheRepo() CharacterCacheRepo {
	characterCacheRepo := &characterCacheRepoImpl{
		cache: make(map[int]*CharacterModel),
		lock:  &sync.RWMutex{},
	}

	return characterCacheRepo
}

func (characterCacheRepo *characterCacheRepoImpl) Get(id int) (character *CharacterModel, found bool) {
	characterCacheRepo.lock.RLock()
	defer characterCacheRepo.lock.RUnlock()

	character, found = characterCacheRepo.cache[id]

	return
}

func (characterCacheRepo *characterCacheRepoImpl) Add(character *CharacterModel) {
	if character != nil {
		characterCacheRepo.lock.Lock()
		defer characterCacheRepo.lock.Unlock()

		characterCacheRepo.cache[character.Id] = character
	}
}
