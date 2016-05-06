package auth

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

var CredentialsServiceInstance CredentialsService = NewCredentialsService(ApiKeyRepoInstance)

type credentialsServiceImpl struct {
	apiKeyRepoInferace ApiKeyRepo
}

func NewCredentialsService(apiKeyRepo ApiKeyRepo) CredentialsService {
	credentialsService := &credentialsServiceImpl{apiKeyRepoInferace: apiKeyRepo}

	return credentialsService
}

func (credentialsService *credentialsServiceImpl) GenerateCredentials() *CredentialsModel {
	apiKeyConfig := credentialsService.apiKeyRepoInferace.GetApiKeyConfig()
	timeStamp := getTimeStamp()
	hash := generateHash(apiKeyConfig, timeStamp)

	return &CredentialsModel{
		PublicKey: apiKeyConfig.PublicKey,
		Hash:      hash,
		TimeStamp: timeStamp,
	}
}

func generateHash(apiKeyConfig *ApiKeyModel, timeStamp string) string {
	key := []byte(timeStamp + apiKeyConfig.PrivateKey + apiKeyConfig.PublicKey)

	md5Crypto := md5.New()
	md5Crypto.Write(key)
	hash := hex.EncodeToString(md5Crypto.Sum(nil))

	return hash
}

func getTimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}
