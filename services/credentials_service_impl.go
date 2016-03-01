package services

import (
	//"crypto/md5"
	//"encoding/hex"
	"github.com/trwalker/marvel-go/models"
	"github.com/trwalker/marvel-go/repos"
	//"strconv"
	//"time"
)

var CredentialsServiceInstance CredentialsService = &CredentialsServiceImpl{ApiKeyRepoInferace: repos.ApiKeyRepoInstance}

type CredentialsServiceImpl struct {
	ApiKeyRepoInferace repos.ApiKeyRepo
}

func (credentialsService *CredentialsServiceImpl) GenerateCredentials() models.CredentialsModel {

	return models.CredentialsModel{}
}

func GetMd5Hash(apiKeyConfig *models.ApiKeyConfigModel) {
	// timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)

	// key := timestamp + apiKeyConfig.PrivateKey + apiKeyConfig.PublicKey

	//md5Crypto := md5.New()

	//hash := hex.EncodeToString(md5Crypto.Sum(key))
}

func GetTime() {
	return
}
