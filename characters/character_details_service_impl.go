package characters

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/rest"
	"time"
	"sync"
)

const getCharacterUrlFormat string = "http://gateway.marvel.com/v1/public/characters/%d?ts=%s&apikey=%s&hash=%s";
const timeout time.Duration = time.Millisecond * 4000

var CharacterDetailsServiceInstance CharacterDetailsService = &CharacterDetailsServiceImpl{
	CharacterMapRepoInterface:   CharacterMapRepoInstance,
	CharacterCacheRepoInterface: CharacterCacheRepoInstance,
	CredentialsServiceInterface: auth.CredentialsServiceInstance,
	RestClientAdapterInterface: rest.RestClientAdapterInstance,
}

type CharacterDetailsServiceImpl struct {
	CharacterMapRepoInterface   CharacterMapRepo
	CharacterCacheRepoInterface CharacterCacheRepo
	CredentialsServiceInterface auth.CredentialsService
	RestClientAdapterInterface rest.RestClientAdapter
	characterDetailsCache map[int]*CharacterDetailsModel
	lock *sync.RWMutex
}

func (characterDetailsService *CharacterDetailsServiceImpl) GetCharacter(name string) (characterDetails *CharacterDetailsModel, found bool, err error) {
	characterDetails = nil
	found = false
	err = nil

	character, mappingFound := tryGetCharacterByName(characterDetailsService, name)

	if mappingFound {
		characterDetails, found = characterDetailsService.CharacterCacheRepoInterface.Get(character.Id)

		if !found {
			characterDetails, found, err = getCharacterFromMarvelApi(characterDetailsService, character)

			if found {
				characterDetailsService.CharacterCacheRepoInterface.Add(characterDetails)
			}
		}
	}

	return 
}

func tryGetCharacterByName(characterDetailsService *CharacterDetailsServiceImpl, name string) (characterModel *CharacterModel, found bool) {
	characterMap := characterDetailsService.CharacterMapRepoInterface.GetCharacterMap()
	characterModel, found = characterMap[name]

	return
}

func getCharacterFromMarvelApi(characterDetailsService *CharacterDetailsServiceImpl, character *CharacterModel) (characterDetails *CharacterDetailsModel, found bool, err error) {
	credentials := characterDetailsService.CredentialsServiceInterface.GenerateCredentials()

	requestUrl := fmt.Sprintf(getCharacterUrlFormat, character.Id, credentials.TimeStamp, credentials.PublicKey, credentials.Hash)

	resp, body, restErr := characterDetailsService.RestClientAdapterInterface.Get(requestUrl, timeout)

	if restErr != nil {
		err = restErr
	} else {
		switch resp.StatusCode {
			case 200:
				found = true
				characterDetails, found, err = parseCharacterJson(body, resp)
			case 404:
				found = false
			default:
				err = errors.New(fmt.Sprintf("Unexpected response code: %d", resp.StatusCode))
		}
	}

	return
}

func parseCharacterJson(body string, resp *http.Response) (characterDetails *CharacterDetailsModel, found bool, err error) {
	characterDetails = nil
	found = true

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonData)

	if err == nil {
		data := jsonData["data"].(map[string]interface {})
		results := data["results"].([]interface {})

		if len(results) == 0 {
			found = false
		} else {
			characterResult := results[0].(map[string]interface {})
			thumbnail := characterResult["thumbnail"].(map[string]interface {})

			characterDetails = &CharacterDetailsModel {
				Id: int(characterResult["id"].(float64)),
				Name: characterResult["name"].(string),
				Description: characterResult["description"].(string),
				Image: fmt.Sprintf("%v.%v", thumbnail["path"].(string), thumbnail["extension"].(string)),
			}
		}
	}

	return
}
