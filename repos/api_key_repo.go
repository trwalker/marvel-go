package repos

import (
	"github.com/trwalker/marvel-go/models"
)

type ApiKeyRepo interface {
	GetApiKeyConfig() *models.ApiKeyConfigModel
}
