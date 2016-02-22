package repos

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/models"
	"io/ioutil"
	"sync"
)

var ApiKeyRepoInstance ApiKeyRepo = initApiKeyRepo()

var apiKeyConfigModel *models.ApiKeyConfigModel

type ApiKeyRepoImpl struct {
}

func initApiKeyRepo() ApiKeyRepo {
	apiKeyConfigModel = nil

	return &ApiKeyRepoImpl{}
}

func (apiKeyRepo *ApiKeyRepoImpl) GetApiKeyConfig() *models.ApiKeyConfigModel {

	if apiKeyConfigModel == nil {
		lock := &sync.Mutex{}

		lock.Lock()
		defer lock.Unlock()

		if apiKeyConfigModel == nil {
			rawJson, fileErr := ioutil.ReadFile("../config/apikey_config.json")

			if fileErr != nil {
				// TODO: error logging
			} else {
				var model models.ApiKeyConfigModel
				jsonErr := json.Unmarshal(rawJson, &model)

				if jsonErr != nil {
					// TODO: error logging
				} else {
					apiKeyConfigModel = &model
				}
			}
		}
	}

	return apiKeyConfigModel
}
