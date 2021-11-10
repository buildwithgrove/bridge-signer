package service

import (
	"context"
	"fmt"

	"github.com/poktbridge/protocols/sign"
)

type SignService interface {
	ECDSASign(ctx context.Context, pass string, data []byte) ([]byte, error)
	EdDSASign(ctx context.Context, pass string, data []byte) ([]byte, error)
}

type Wrapper struct {
	s        SignService
	ethPass  string
	poktPass string
}

func NewServerWrapper(s SignService, ethPass, poktPass string) *Wrapper {
	return &Wrapper{s, ethPass, poktPass}
}

func (w Wrapper) ECDSASign(ctx context.Context, in *sign.ECDSASignRequest) (*sign.ECDSASignReply, error) {
	pass := fmt.Sprintf("%s%s", w.ethPass, in.Pass)

	res, err := w.s.ECDSASign(ctx, pass, in.Data)
	if err != nil {
		return nil, err
	}

	return &sign.ECDSASignReply{Sig: res}, nil
}
func (w Wrapper) EdDsaSign(ctx context.Context, in *sign.EdDsaSignRequest) (*sign.EdDsaSignReply, error) {
	pass := fmt.Sprintf("%s%s", w.poktPass, in.Pass)

	res, err := w.s.EdDSASign(ctx, pass, in.Data)
	if err != nil {
		return nil, err
	}

	return &sign.EdDsaSignReply{Sig: res}, nil
}
