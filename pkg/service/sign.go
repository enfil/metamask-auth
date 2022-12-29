package service

import (
	"github.com/enfil/metamask-auth/internal/domain/user"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

type Sign struct {
}

func (s Sign) Check(cryptoAddress string, nonce string, sigHex string) error {
	sig := hexutil.MustDecode(sigHex)
	// https://github.com/ethereum/go-ethereum/blob/master/internal/ethapi/api.go#L516
	// check here why I am subtracting 27 from the last byte
	sig[crypto.RecoveryIDOffset] -= 27
	msg := accounts.TextHash([]byte(nonce))

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return err
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	if cryptoAddress != strings.ToLower(recoveredAddr.Hex()) {
		return user.ErrAuthError
	}

	return nil
}
