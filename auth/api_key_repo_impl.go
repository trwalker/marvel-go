package auth

import (
	"github.com/trwalker/marvel-go/config/key"
	"sync"
)

var ApiKeyRepoInstance ApiKeyRepo = &ApiKeyRepoImpl{
	lock: &sync.Mutex{},
}

type ApiKeyRepoImpl struct {
	apiKeyModel *ApiKeyModel
	lock        *sync.Mutex
}

func (apiKeyRepo *ApiKeyRepoImpl) GetApiKeyConfig() *ApiKeyModel {

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
