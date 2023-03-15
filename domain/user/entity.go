package user

import (
	"github.com/enfil/metamask-auth/domain/user/vo"
	"github.com/google/uuid"
	"regexp"
)

var (
	HexRegex   *regexp.Regexp = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	NonceRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]+$`)
)

type Entity struct {
	uuid uuid.UUID
	auth vo.AuthorizationData
}

func New(uuid uuid.UUID, cryptoAddress string, nonce string, sig string) (Entity, error) {
	err := matchAddress(cryptoAddress)
	if err != nil {
		return Entity{}, err
	}
	return Entity{
		uuid: uuid,
		auth: vo.AuthorizationData{
			CryptoAddress: cryptoAddress,
			Nonce:         nonce,
			Sig:           sig,
		},
	}, nil
}
func (u Entity) Edit(cryptoAddress string, nonce string, sig string) error {
	if err := validate(cryptoAddress, nonce, sig); err != nil {
		return err
	}
	u.auth.CryptoAddress = cryptoAddress
	u.auth.Nonce = nonce
	u.auth.Sig = sig
	return nil
}

func (u Entity) Uuid() uuid.UUID {
	return u.uuid
}

func (u Entity) CryptoAddress() string {
	return u.auth.CryptoAddress
}

func (u Entity) Nonce() string {
	return u.auth.Nonce
}

func (u Entity) SetNonce(nonce string) {
	u.auth.Nonce = nonce
}

func validate(cryptoAddress string, nonce string, sig string) error {
	err := matchAddress(cryptoAddress)
	if err != nil {
		return err
	}

	if !NonceRegex.MatchString(nonce) {
		return ErrInvalidNonce
	}

	if len(sig) == 0 {
		return ErrMissingSig
	}

	return nil
}

func matchAddress(cryptoAddress string) error {
	if !HexRegex.MatchString(cryptoAddress) {
		return ErrInvalidAddress
	}
	return nil
}
