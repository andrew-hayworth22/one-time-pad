package pad

import "errors"

type Pad struct {
	key []byte
}

func New(key string) *Pad {
	return &Pad{
		key: []byte(key),
	}
}

func (p Pad) Encrypt(plaintext string) (string, error) {
	if len(plaintext) > len(p.key) {
		return "", errors.New("the text length must be shorter than or the same length as the pad")
	}

	plaintextBytes := []byte(plaintext)
	ciphertextBytes := make([]byte, len(plaintextBytes))

	for idx, b := range plaintextBytes {
		ciphertextBytes[idx] = (b + p.key[idx]) % 128
	}

	ciphertext := string(ciphertextBytes)

	return ciphertext, nil
}

func (p Pad) Decrypt(ciphertext string) (string, error) {
	if len(ciphertext) > len(p.key) {
		return "", errors.New("the text length must be shorter than or the same length as the pad")
	}

	ciphertextBytes := []byte(ciphertext)
	plaintextBytes := make([]byte, len(ciphertextBytes))

	for idx, b := range ciphertextBytes {
		plaintextBytes[idx] = (b - p.key[idx]) % 128
	}

	plaintext := string(plaintextBytes)

	return plaintext, nil
}
