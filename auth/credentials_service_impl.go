package auth

import (
//"crypto/md5"
//"encoding/hex"
//"strconv"
//"time"
)

var CredentialsServiceInstance CredentialsService = &CredentialsServiceImpl{ApiKeyRepoInferace: ApiKeyRepoInstance}

type CredentialsServiceImpl struct {
	ApiKeyRepoInferace ApiKeyRepo
}

func (credentialsService *CredentialsServiceImpl) GenerateCredentials() CredentialsModel {

	return CredentialsModel{}
}

func GetMd5Hash(apiKeyConfig *ApiKeyConfigModel) {
	// timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)

	// key := timestamp + apiKeyConfig.PrivateKey + apiKeyConfig.PublicKey

	//md5Crypto := md5.New()

	//hash := hex.EncodeToString(md5Crypto.Sum(key))
}

func GetTime() {
	return
}
