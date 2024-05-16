package cypherer

import (
	"encoding/base64"
	"errors"
)

func Encrypt(rawString, key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("secret key cannot be empty")
	}
	encryptedByted, err := process(
		[]byte(rawString),
		[]byte(key),
	)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(
		encryptedByted,
	), nil
}

func Decrypt(rawString, key string) (string, error) {
	cypheredByted, err := base64.StdEncoding.DecodeString(rawString)
	if err != nil {
		return "", errors.New("error while base64 decoding")
	}
	decryptedBytes, err := process(cypheredByted, []byte(key))
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

func process(input, secretKey []byte) ([]byte, error) {
	if len(secretKey) == 0 {
		return []byte{}, errors.New("secret key cannot be empty")
	}
	for i, b := range input {
		input[i] = b ^ secretKey[i%len(secretKey)]
	}
	return input, nil
}
