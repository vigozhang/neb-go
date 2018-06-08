package utils

import (
	"io"
	"crypto/rand"

	"github.com/spacemonkeygo/openssl"
)

// RandomCSPRNG a cryptographically secure pseudo-random number generator
func RandomCSPRNG(n int) []byte {
	buff := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, buff)
	if err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	return buff
}

// ZeroBytes clears byte slice.
func ZeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}

func OpensslEncrypt(input []byte, algorithm string, key []byte, iv []byte) ([]byte, error) {
	cipher, err := openssl.GetCipherByName(algorithm)
	if err != nil {
		return nil, err
	}

	ctx, err := openssl.NewEncryptionCipherCtx(cipher, nil, key, iv)
	if err != nil {
		return nil, err
	}

	cipherbytes, err := ctx.EncryptUpdate(input)
	if err != nil {
		return nil, err
	}

	finalbytes, err := ctx.EncryptFinal()
	if err != nil {
		return nil, err
	}

	cipherbytes = append(cipherbytes, finalbytes...)
	return cipherbytes, nil
}

func OpensslDecrypt(input []byte, algorithm string, key []byte, iv []byte) ([]byte, error) {
	cipher, err := openssl.GetCipherByName(algorithm)
	if err != nil {
		return nil, err
	}

	ctx, err := openssl.NewDecryptionCipherCtx(cipher, nil, key, iv)

	if err != nil {
		return nil, err
	}

	cipherbytes, err := ctx.DecryptUpdate(input)
	if err != nil {
		return nil, err
	}

	finalbytes, err := ctx.DecryptFinal()
	if err != nil {
		return nil, err
	}

	cipherbytes = append(cipherbytes, finalbytes...)
	return cipherbytes, nil
}
