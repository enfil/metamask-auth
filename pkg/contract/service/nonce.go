package contract

type NonceProvider interface {
	GenerateNonce() (string, error)
}
