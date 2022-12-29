package command

type SignIn struct {
	CryptoAddress string `json:"address"`
	Nonce         string `json:"nonce"`
	Sig           string `json:"sig"`
}
