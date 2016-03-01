package auth

type CredentialsService interface {
	GenerateCredentials() CredentialsModel
}
