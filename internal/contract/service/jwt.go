package contract

type JwtProvider interface {
	Create(subject string) (string, error)
	VerifyAndGetSubject(tokenString string) (string, error)
}
