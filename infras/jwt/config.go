package jwt

type Config struct {
	PrivateKey string // jwt编码私钥
	ExpSeconds int    // jwt编码超时时间
}
