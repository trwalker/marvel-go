package characters

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/rest"
	"time"
)

const getCharacterUrlFormat string = "http://gateway.marvel.com/v1/public/characters/%d?ts=%s&apikey=%s&hash=%s";
const timeout time.Duration = time.Millisecond * 4000

var CharacterDetailsServiceInstance CharacterDetailsService = &CharacterDetailsServiceImpl{
	CharacterMapRepoInterface:   CharacterMapRepoInstance,
	CredentialsServiceInterface: auth.CredentialsServiceInstance,
	RestClientAdapterInterface: rest.RestClientAdapterInstance,
}

type CharacterDetailsServiceImpl struct {
	CharacterMapRepoInterface   CharacterMapRepo
	CredentialsServiceInterface auth.CredentialsService
	RestClientAdapterInterface rest.RestClientAdapter
}

func (characterDetailsService *CharacterDetailsServiceImpl) GetCharacter(name string) (characterDetails *CharacterDetailsModel, found bool, err error) {
	// Get ID
	var character *CharacterModel
	character, found = tryGetCharacterByName(characterDetailsService, name)

	if found {
		// Get Credentials
		credentials := characterDetailsService.CredentialsServiceInterface.GenerateCredentials()

		// Build URL
		requestUrl := fmt.Sprintf(getCharacterUrlFormat, character.Id, credentials.TimeStamp, credentials.PublicKey, credentials.Hash)

		// Make REST Call
		resp, body, restErr := characterDetailsService.RestClientAdapterInterface.Get(requestUrl, timeout)

		// Parse response JSON into Model
		if restErr != nil {
			err = restErr
		} else {
			characterDetails, found, err = tryParseCharacterJson(body, resp)
		}
	}

	return 
}

func tryGetCharacterByName(characterDetailsService *CharacterDetailsServiceImpl, name string) (*CharacterModel, bool) {
	characterMap := characterDetailsService.CharacterMapRepoInterface.GetCharacterMap()
	characterModel, found := characterMap[name]

	return characterModel, found
}

func tryParseCharacterJson(body string, resp *http.Response) (characterDetails *CharacterDetailsModel, found bool, err error) {
	switch resp.StatusCode {
	case 200:
		found = true
	case 404:
		found = false
	default:
		err = errors.New(fmt.Sprintf("Unexpected response code: %d", resp.StatusCode))
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(body), &jsonData)

	// "data":map[string]interface {}
	// "results":[]interface {}

	/*
		var characterData = data.data.results[0];

    return {
        id: characterData.id,
        name: characterData.name,
        description: characterData.description,
        image: characterData.thumbnail.path + '.' + characterData.thumbnail.extension,
        comics: characterData.comics.items
    };
	*/

	_ = "breakpoint"

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
