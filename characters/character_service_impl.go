package characters

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/rest"
	"net/http"
	"sync"
	"time"
)

const getCharacterUrlFormat string = "http://gateway.marvel.com/v1/public/characters/%d?ts=%s&apikey=%s&hash=%s"
const timeout time.Duration = time.Millisecond * 4000

var CharacterServiceInstance CharacterService = &CharacterServiceImpl{
	CharacterMapRepoInterface:   CharacterMapRepoInstance,
	CharacterCacheRepoInterface: CharacterCacheRepoInstance,
	CredentialsServiceInterface: auth.CredentialsServiceInstance,
	RestClientAdapterInterface:  rest.RestClientAdapterInstance,
}

type CharacterServiceImpl struct {
	CharacterMapRepoInterface   CharacterMapRepo
	CharacterCacheRepoInterface CharacterCacheRepo
	CredentialsServiceInterface auth.CredentialsService
	RestClientAdapterInterface  rest.RestClientAdapter
	characterCache              map[int]*CharacterModel
	lock                        *sync.RWMutex
}

func (characterService *CharacterServiceImpl) GetCharacter(name string) (character *CharacterModel, found bool, err error) {
	character = nil
	found = false
	err = nil

	characterId, mappingFound := tryGetCharacterByName(characterService, name)

	if mappingFound {
		character, found = characterService.CharacterCacheRepoInterface.Get(characterId)

		if !found {
			character, found, err = getCharacterFromMarvelApi(characterService, characterId)

			if found {
				characterService.CharacterCacheRepoInterface.Add(character)
			}
		}
	}

	return
}

func tryGetCharacterByName(characterService *CharacterServiceImpl, name string) (characterId int, found bool) {
	characterMap := characterService.CharacterMapRepoInterface.GetCharacterMap()
	characterId, found = characterMap[name]

	return
}

func getCharacterFromMarvelApi(characterService *CharacterServiceImpl, characterId int) (character *CharacterModel, found bool, err error) {
	credentials := characterService.CredentialsServiceInterface.GenerateCredentials()

	requestUrl := fmt.Sprintf(getCharacterUrlFormat, characterId, credentials.TimeStamp, credentials.PublicKey, credentials.Hash)

	resp, body, restErr := characterService.RestClientAdapterInterface.Get(requestUrl, timeout)

	if restErr != nil {
		err = restErr
	} else {
		switch resp.StatusCode {
		case 200:
			found = true
			character, found, err = parseCharacterJson(body, resp)
		case 404:
			found = false
		default:
			err = errors.New(fmt.Sprintf("Unexpected response code: %d", resp.StatusCode))
		}
	}

	return
}

func parseCharacterJson(body string, resp *http.Response) (character *CharacterModel, found bool, err error) {
	character = nil
	found = true

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonData)

	if err == nil {
		data := jsonData["data"].(map[string]interface{})
		results := data["results"].([]interface{})

		if len(results) == 0 {
			found = false
		} else {
			characterResult := results[0].(map[string]interface{})
			thumbnail := characterResult["thumbnail"].(map[string]interface{})

			comicsObject := characterResult["comics"].(map[string]interface{})
			comicItems := comicsObject["items"].([]interface{})

			comics := make([]string, len(comicItems))

			for i, value := range comicItems {
				comic := value.(map[string]interface{})
				comicName := comic["name"].(string)

				comics[i] = comicName
			}

			character = &CharacterModel{
				Id:          int(characterResult["id"].(float64)),
				Name:        characterResult["name"].(string),
				Description: characterResult["description"].(string),
				Image:       fmt.Sprintf("%v.%v", thumbnail["path"].(string), thumbnail["extension"].(string)),
				Comics:      comics,
			}
		}
	}

	return
}
