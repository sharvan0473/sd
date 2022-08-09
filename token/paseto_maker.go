package token

import (
	"errors"
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	synmetricKey []byte
}

func NewPasetoMaker(synmetricKey string) (Maker, error) {
	if len(synmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid Key size")
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		synmetricKey: []byte(synmetricKey),
	}
	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.synmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.synmetricKey, payload, nil)
	if err != nil {
		return nil, errors.New("Invlaid token")
	}
	err = payload.Valid()

	if err != nil {
		return nil, errors.New("Ina")
	}
	return payload, nil
}
