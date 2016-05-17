package authmock

import (
	"github.com/trwalker/marvel-go/auth"
)

type CredentialsServiceMock struct {
	GenerateCredentialsMock func() *auth.CredentialsModel
}

func (credentialsServiceMock *CredentialsServiceMock) GenerateCredentials() *auth.CredentialsModel {
	return credentialsServiceMock.GenerateCredentialsMock()
}
