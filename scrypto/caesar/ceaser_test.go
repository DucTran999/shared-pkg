package caesar_test

import (
	"testing"

	"github.com/DucTran999/shared-pkg/scrypto/caesar"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name  string
	input string
	nonce int
}

func Test_CaesarCryptoGraphy(t *testing.T) {
	testcases := []testTable{
		{
			name:  "message with non chars",
			input: "daniel!",
			nonce: 8,
		},
		{
			name:  "message encrypt with nonce negative",
			input: "daniel",
			nonce: -2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			cipher := caesar.CaesarEncrypt(tc.input, tc.nonce)
			plaintext := caesar.CaesarDecrypt(cipher, tc.nonce)

			assert.Equal(t, tc.input, plaintext)
		})
	}

}
