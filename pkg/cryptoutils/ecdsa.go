package cryptoutils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func RestoreEcdsaPrivateKey(file string) (*ecdsa.PrivateKey, error) {
	rawKey, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("%s: %w", errMsgFailedToReadPrivKeyFile, err)

		return nil, err
	}

	block, _ := pem.Decode(rawKey)
	if block == nil {
		return nil, ErrInvalidPemPrivKeyFormat
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("%s: %w", errMsgFailedToParsePrivKey, err)

		return nil, err
	}

	return key, nil
}

func RestoreEcdsaPublicKey(file string) (*ecdsa.PublicKey, error) {
	rawKey, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("%s: %w", errMsgFailedToReadPubKeyFile, err)

		return nil, err
	}

	block, _ := pem.Decode(rawKey)
	if block == nil {
		return nil, ErrInvalidPemPubKeyFormat
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("%s: %w", errMsgFailedToParsePubKey, err)

		return nil, err
	}

	key, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil, ErrNotEcdsaPubKey
	}

	return key, nil
}

type Ecdsa struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func RestoreEcdsaKeyPair(privKeyFile, pubKeyFile string) (*Ecdsa, error) {
	priv, err := RestoreEcdsaPrivateKey(privKeyFile)
	if err != nil {
		return nil, err
	}

	pub, err := RestoreEcdsaPublicKey(pubKeyFile)
	if err != nil {
		return nil, err
	}

	return &Ecdsa{
		PrivateKey: priv,
		PublicKey:  pub,
	}, nil
}
