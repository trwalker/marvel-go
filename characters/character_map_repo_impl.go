package characters

import (
	"sync"
)

var CharacterMapRepoInstance CharacterMapRepo = &CharacterMapRepoImpl{
	characterMap: make(map[string]*CharacterModel),
	lock: &sync.Mutex{},
}

type CharacterMapRepoImpl struct {
	characterMap map[string]*CharacterModel
	lock *sync.Mutex
}

func (characterMapRepo *CharacterMapRepoImpl) GetCharacterMap() map[string]*CharacterModel {
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
	spiderMan := &CharacterModel{
		Id:    1009610,
		Name:  "spider-man",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/3/50/526548a343e4b.jpg",
	}

	characterMapRepo.characterMap[spiderMan.Name] = spiderMan

	hulk := &CharacterModel{
		Id:    1009351,
		Name:  "hulk",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/5/a0/538615ca33ab0.jpg",
	}

	characterMapRepo.characterMap[hulk.Name] = hulk

	captainAmerica := &CharacterModel{
		Id:    1009220,
		Name:  "captain-america",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/3/50/537ba56d31087.jpg",
	}

	characterMapRepo.characterMap[captainAmerica.Name] = captainAmerica

	ironMan := &CharacterModel{
		Id:    1009368,
		Name:  "iron-man",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/9/c0/527bb7b37ff55.jpg",
	}

	characterMapRepo.characterMap[ironMan.Name] = ironMan

	thor := &CharacterModel{
		Id:    1009664,
		Name:  "thor",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/d/d0/5269657a74350.jpg",
	}

	characterMapRepo.characterMap[thor.Name] = thor

	wolverine := &CharacterModel{
		Id:    1009718,
		Name:  "wolverine",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/2/60/537bcaef0f6cf.jpg",
	}

	characterMapRepo.characterMap[wolverine.Name] = wolverine

	storm := &CharacterModel{
		Id:    1009629,
		Name:  "storm",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/40/526963dad214d.jpg",
	}

	characterMapRepo.characterMap[storm.Name] = storm

	jeanGrey := &CharacterModel{
		Id:    1009496,
		Name:  "jean-grey",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/f/30/4bc654cf9d0ac.jpg",
	}

	characterMapRepo.characterMap[jeanGrey.Name] = jeanGrey

	gambit := &CharacterModel{
		Id:    1009313,
		Name:  "gambit",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/a/40/52696aa8aee99.jpg",
	}

	characterMapRepo.characterMap[gambit.Name] = gambit

	cyclops := &CharacterModel{
		Id:    1009257,
		Name:  "cyclops",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/70/526547e2d90ad.jpg",
	}

	characterMapRepo.characterMap[cyclops.Name] = cyclops

	beast := &CharacterModel{
		Id:    1009175,
		Name:  "beast",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/2/80/511a79a0451a3.jpg",
	}

	characterMapRepo.characterMap[beast.Name] = beast
}
