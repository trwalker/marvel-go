package auth

type ApiKeyRepo interface {
	GetApiKeyConfig() *ApiKeyConfigModel
}
