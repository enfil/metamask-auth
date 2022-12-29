package contract

type SignProvider interface {
	Check(cryptoAddress string, nonce string, sigHex string) error
}
