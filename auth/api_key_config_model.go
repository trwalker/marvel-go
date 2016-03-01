package auth

type ApiKeyConfigModel struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}
