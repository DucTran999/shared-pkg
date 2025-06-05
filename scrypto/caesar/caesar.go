package caesar

import (
	"unicode"
)

// CaesarEncrypt shifts letters by 'shift' positions
func CaesarEncrypt(text string, shift int) string {
	shift = ((shift % 26) + 26) % 26
	result := []rune{}

	for _, ch := range text {
		if unicode.IsLetter(ch) {
			base := 'A'
			if unicode.IsLower(ch) {
				base = 'a'
			}
			// Rotate and wrap using modulo
			encrypted := ((ch-base+rune(shift))%26+26)%26 + base
			result = append(result, encrypted)
		} else {
			// Keep non-letter characters unchanged
			result = append(result, ch)
		}
	}

	return string(result)
}

// CaesarDecrypt reverses the Caesar encryption
func CaesarDecrypt(text string, shift int) string {
	return CaesarEncrypt(text, 26-shift)
}
