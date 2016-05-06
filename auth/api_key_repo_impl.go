package auth

import (
	"github.com/trwalker/marvel-go/config/key"
	"sync"
)

var ApiKeyRepoInstance ApiKeyRepo = NewApiKeyRepo()

type apiKeyRepoImpl struct {
	apiKeyModel *ApiKeyModel
	lock        *sync.Mutex
}

func NewApiKeyRepo() ApiKeyRepo {
	apiKeyRepo := &apiKeyRepoImpl{
		lock: &sync.Mutex{},
	}

	return apiKeyRepo
}

func (apiKeyRepo *apiKeyRepoImpl) GetApiKeyConfig() *ApiKeyModel {

	if apiKeyRepo.apiKeyModel == nil {
		apiKeyRepo.lock.Lock()
		defer apiKeyRepo.lock.Unlock()

		if apiKeyRepo.apiKeyModel == nil {
			apiKeyRepo.apiKeyModel = &ApiKeyModel{
				PublicKey:  key.PublicKey,
				PrivateKey: key.PrivateKey,
			}
		}
	}

	return apiKeyRepo.apiKeyModel
}
