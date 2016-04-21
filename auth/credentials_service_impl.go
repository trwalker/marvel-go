package auth

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

var CredentialsServiceInstance CredentialsService = &CredentialsServiceImpl{ApiKeyRepoInferace: ApiKeyRepoInstance}

type CredentialsServiceImpl struct {
	ApiKeyRepoInferace ApiKeyRepo
}

func (credentialsService *CredentialsServiceImpl) GenerateCredentials() *CredentialsModel {
	apiKeyConfig := credentialsService.ApiKeyRepoInferace.GetApiKeyConfig()
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
