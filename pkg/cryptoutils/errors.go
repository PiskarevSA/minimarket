package cryptoutils

import "errors"

const (
	errMsgFailedToReadPrivKeyFile = "failed to read private key file"
	errMsgFailedToReadPubKeyFile  = "failed to read private key file"
	errMsgFailedToParsePubKey     = "failed to parse public key"
	errMsgFailedToParsePrivKey    = "failed to parse private key"
	errMsgNotEcdsaPubKey          = "not ecdsa public key"
)

var (
	ErrInvalidPemPubKeyFormat  = errors.New("invalid pem public key format")
	ErrInvalidPemPrivKeyFormat = errors.New("invalid pem private key format")
	ErrNotEcdsaPubKey          = errors.New("not ecdsa public key")
)
