package repos

import (
	"github.com/trwalker/marvel-go/models"
	"sync"
)

var CharacterMapRepoInstance CharacterMapRepo = constructor()

var characters map[string]*models.CharacterModel

type CharacterMapRepoImpl struct {
}

func (characterMapRepo *CharacterMapRepoImpl) GetCharacterMap() map[string]*models.CharacterModel {
	if len(characters) == 0 {
		lock := &sync.Mutex{}

		lock.Lock()
		defer lock.Unlock()

		if len(characters) == 0 {
			buildCharacterMap()
		}
	}

	return characters
}

func constructor() CharacterMapRepo {
	var characterMapRepo *CharacterMapRepoImpl = &CharacterMapRepoImpl{}

	return characterMapRepo
}

func buildCharacterMap() {
	characters = make(map[string]*models.CharacterModel)

	spiderMan := &models.CharacterModel{
		Id:    1009610,
		Name:  "spider-man",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/3/50/526548a343e4b.jpg",
	}

	characters[spiderMan.Name] = spiderMan

	hulk := &models.CharacterModel{
		Id:    1009351,
		Name:  "hulk",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/5/a0/538615ca33ab0.jpg",
	}

	characters[hulk.Name] = hulk

	captainAmerica := &models.CharacterModel{
		Id:    1009220,
		Name:  "captain-america",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/3/50/537ba56d31087.jpg",
	}

	characters[captainAmerica.Name] = captainAmerica

	ironMan := &models.CharacterModel{
		Id:    1009368,
		Name:  "iron-man",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/9/c0/527bb7b37ff55.jpg",
	}

	characters[ironMan.Name] = ironMan

	thor := &models.CharacterModel{
		Id:    1009664,
		Name:  "thor",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/d/d0/5269657a74350.jpg",
	}

	characters[thor.Name] = thor

	wolverine := &models.CharacterModel{
		Id:    1009718,
		Name:  "wolverine",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/2/60/537bcaef0f6cf.jpg",
	}

	characters[wolverine.Name] = wolverine

	storm := &models.CharacterModel{
		Id:    1009629,
		Name:  "storm",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/40/526963dad214d.jpg",
	}

	characters[storm.Name] = storm

	jeanGrey := &models.CharacterModel{
		Id:    1009496,
		Name:  "jean-grey",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/f/30/4bc654cf9d0ac.jpg",
	}

	characters[jeanGrey.Name] = jeanGrey

	gambit := &models.CharacterModel{
		Id:    1009313,
		Name:  "gambit",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/a/40/52696aa8aee99.jpg",
	}

	characters[gambit.Name] = gambit

	cyclops := &models.CharacterModel{
		Id:    1009257,
		Name:  "cyclops",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/70/526547e2d90ad.jpg",
	}

	characters[cyclops.Name] = cyclops

	beast := &models.CharacterModel{
		Id:    1009175,
		Name:  "beast",
		Image: "http://i.annihil.us/u/prod/marvel/i/mg/2/80/511a79a0451a3.jpg",
	}

	characters[beast.Name] = beast
}