package vo

type AuthorizationData struct {
	CryptoAddress string
	Nonce         string
	Sig           string
	SignedToken   string
}
