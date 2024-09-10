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

func (p Pad) Encrypt(str string) (string, error) {
	if len(str) > len(p.key) {
		return "", errors.New("the text length must be shorter than or the same length as the pad")
	}

	strBytes := []byte(str)
	resultBytes := make([]byte, len(strBytes))

	for idx, b := range strBytes {
		resultBytes[idx] = (b + p.key[idx]) % 255
	}

	result := string(resultBytes)

	return result, nil
}

func (p Pad) Decrypt(str string) (string, error) {
	if len(str) > len(p.key) {
		return "", errors.New("the text length must be shorter than or the same length as the pad")
	}

	strBytes := []byte(str)
	resultBytes := make([]byte, len(strBytes))

	for idx, b := range strBytes {
		resultBytes[idx] = (b - p.key[idx]) % 255
	}

	result := string(resultBytes)

	return result, nil
}
