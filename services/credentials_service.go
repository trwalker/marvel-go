package services

import (
	"github.com/trwalker/marvel-go/models"
)

type CredentialsService interface {
	GenerateCredentials() models.CredentialsModel
}
