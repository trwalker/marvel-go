package auth

type ApiKeyRepo interface {
	GetApiKeyConfig() *ApiKeyModel
}
