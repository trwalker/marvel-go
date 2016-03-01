package auth

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var ApiKeyRepoInstance ApiKeyRepo = &ApiKeyRepoImpl{}

type ApiKeyRepoImpl struct {
	apiKeyConfigModel *ApiKeyConfigModel
}

func (apiKeyRepo *ApiKeyRepoImpl) GetApiKeyConfig() *ApiKeyConfigModel {

	if apiKeyRepo.apiKeyConfigModel == nil {
		lock := &sync.Mutex{}

		lock.Lock()
		defer lock.Unlock()

		if apiKeyRepo.apiKeyConfigModel == nil {
			rawJson, fileErr := ioutil.ReadFile("../config/auth/api_key_config.json")

			if fileErr != nil {
				// TODO: error logging
			} else {
				var model ApiKeyConfigModel
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
