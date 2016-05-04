package charrepos

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/characters/models"
	"github.com/trwalker/marvel-go/rest"
	"net/http"
	"time"
)

const getCharacterUrlFormat string = "http://gateway.marvel.com/v1/public/characters/%d?ts=%s&apikey=%s&hash=%s"
const getCharactertimeout time.Duration = time.Millisecond * 4000

var CharacterRepoInstance CharacterRepo = &CharacterRepoImpl{
	RestClientAdapterInterface: rest.RestClientAdapterInstance,
}

type CharacterRepoImpl struct {
	RestClientAdapterInterface rest.RestClientAdapter
}

func (characterRepo *CharacterRepoImpl) GetCharacter(characterId int, credentials *auth.CredentialsModel) (character *charmodels.CharacterModel, found bool, err error) {
	character, found, err = getCharacterFromMarvelApi(characterRepo, characterId, credentials)

	return
}

func getCharacterFromMarvelApi(characterRepo *CharacterRepoImpl, characterId int, credentials *auth.CredentialsModel) (character *charmodels.CharacterModel, found bool, err error) {
	requestUrl := fmt.Sprintf(getCharacterUrlFormat, characterId, credentials.TimeStamp, credentials.PublicKey, credentials.Hash)

	resp, body, restErr := characterRepo.RestClientAdapterInterface.Get(requestUrl, getCharactertimeout)

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

func parseCharacterJson(body string, resp *http.Response) (character *charmodels.CharacterModel, found bool, err error) {
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

			character = &charmodels.CharacterModel{
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
