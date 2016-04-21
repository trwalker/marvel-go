package auth

import (
	"sync"
	"github.com/trwalker/marvel-go/config/key"
)

var ApiKeyRepoInstance ApiKeyRepo = &ApiKeyRepoImpl{}

type ApiKeyRepoImpl struct {
	apiKeyModel *ApiKeyModel
}

func (apiKeyRepo *ApiKeyRepoImpl) GetApiKeyConfig() *ApiKeyModel {

	if apiKeyRepo.apiKeyModel == nil {
		lock := &sync.Mutex{}

		lock.Lock()
		defer lock.Unlock()

		if apiKeyRepo.apiKeyModel == nil {
			apiKeyRepo.apiKeyModel = &ApiKeyModel{
				PublicKey: key.PublicKey,
				PrivateKey: key.PrivateKey,
			}
		}
	}

	return apiKeyRepo.apiKeyModel
}
