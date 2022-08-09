package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const minSecretSize = 32

type JWTMaker struct {
	secretKwy string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretSize {
		return nil, fmt.Errorf("Invalid key size atleased 32 charactrer")
	}
	return &JWTMaker{secretKey}, nil
}

func (naker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(naker.secretKwy))

}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(maker.secretKwy), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("Invalid token")) {
			return nil, errors.New("Invalid token")
		}
		return nil, errors.New("Invalid token")
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("Invalid token")
	}
	return payload, nil
}
