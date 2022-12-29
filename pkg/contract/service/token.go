package contract

type TokenProvider interface {
	Create(subject string) (string, error)
	VerifyAndGetSubject(tokenString string) (string, error)
}
