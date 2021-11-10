package service

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	posCrypto "github.com/pokt-network/pocket-core/crypto"
)

type Sign struct {
	ecdsaPK PK
	eddsaPK PK
}

func NewSign(ecdsaPK PK, eddsaPK PK) Sign {
	return Sign{
		ecdsaPK: ecdsaPK,
		eddsaPK: eddsaPK,
	}
}

func (s Sign) ECDSASign(ctx context.Context, pass string, data []byte) ([]byte, error) {
	k, err := s.ecdsaPK.Decode(pass)
	if err != nil {
		return nil, fmt.Errorf("ECDSASign error")
	}

	privateKey, err := crypto.HexToECDSA(string(k))
	if err != nil {
		return nil, fmt.Errorf("ECDSASign error")
	}
	sig, err := crypto.Sign(data, privateKey)
	if err != nil {
		return nil, fmt.Errorf("ECDSASign error")
	}

	return sig, nil
}

func (s Sign) EdDSASign(ctx context.Context, pass string, data []byte) ([]byte, error) {
	k, err := s.eddsaPK.Decode(pass)
	if err != nil {
		return nil, fmt.Errorf("EdDSASign error")
	}

	privKeyBytes, err := hex.DecodeString(string(k))
	if err != nil {
		return nil, err
	}

	priv, err := posCrypto.NewPrivateKeyBz(privKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("EdDSASign error")
	}
	sig, err := priv.Sign(data)
	if err != nil {
		return nil, fmt.Errorf("EdDSASign error")
	}

	return sig, nil
}
