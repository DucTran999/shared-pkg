package caesar_test

import (
	"testing"

	"github.com/DucTran999/shared-pkg/scrypto/caesar"
	"github.com/stretchr/testify/assert"
)

func Test_CaesarCryptoGraphy(t *testing.T) {
	rawMessage := "daniel!"

	nonce := 8
	cipher := caesar.CaesarEncrypt(rawMessage, nonce)
	plaintext := caesar.CaesarDecrypt(cipher, nonce)

	assert.Equal(t, rawMessage, plaintext)
}
