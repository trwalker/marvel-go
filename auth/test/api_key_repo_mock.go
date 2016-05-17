package authmock

import (
	"github.com/trwalker/marvel-go/auth"
)

type ApiKeyRepoMock struct {
	GetApiKeyConfigMock func() *auth.ApiKeyModel
}

func (apiKeyRepoMock *ApiKeyRepoMock) GetApiKeyConfig() *auth.ApiKeyModel {
	return apiKeyRepoMock.GetApiKeyConfigMock()
}
