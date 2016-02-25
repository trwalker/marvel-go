package repos

import (
	"encoding/json"
	"github.com/trwalker/marvel-go/models"
	"io/ioutil"
	"sync"
)

var ApiKeyRepoInstance ApiKeyRepo = &ApiKeyRepoImpl{}

type ApiKeyRepoImpl struct {
	apiKeyConfigModel *models.ApiKeyConfigModel
}

func (apiKeyRepo *ApiKeyRepoImpl) GetApiKeyConfig() *models.ApiKeyConfigModel {

	if apiKeyRepo.apiKeyConfigModel == nil {
		lock := &sync.Mutex{}

		lock.Lock()
		defer lock.Unlock()

		if apiKeyRepo.apiKeyConfigModel == nil {
			rawJson, fileErr := ioutil.ReadFile("../config/apikey_config.json")

			if fileErr != nil {
				// TODO: error logging
			} else {
				var model models.ApiKeyConfigModel
				jsonErr := json.Unmarshal(rawJson, &model)

				if jsonErr != nil {
					// TODO: error logging
				} else {
					apiKeyRepo.apiKeyConfigModel = &model
				}
			}
		}
	}

	return apiKeyRepo.apiKeyConfigModel
}
