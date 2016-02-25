package services

import (
	"github.com/trwalker/marvel-go/models"
	"github.com/trwalker/marvel-go/repos"
)

var CredentialsServiceInstance CredentialsService = &CredentialsServiceImpl{ApiKeyRepoInferace: repos.ApiKeyRepoInstance}

type CredentialsServiceImpl struct {
	ApiKeyRepoInferace repos.ApiKeyRepo
}

func (credentialsService *CredentialsServiceImpl) GenerateCredentials() models.CredentialsModel {

	return models.CredentialsModel{}
}
