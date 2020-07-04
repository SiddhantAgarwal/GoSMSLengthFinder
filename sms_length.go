package smslength

import (
	"math"
	"unicode/utf8"
)

func getCharSet(message string) string {
	for _, char := range message {
		if _, ok := gsm7Bit[char]; !ok {
			if _, ok := gsm7BitExt[char]; !ok {
				return gsmCharsetUnicode
			}
		}
	}
	return gsmCharset7Bit
}

func calculate7BitLengthParts(message string) int64 {

	payload := ""

	for _, char := range message {
		if _, ok := gsm7BitExt[char]; ok {
			payload += string(gsm7BitEsc)
		}
		payload += string(char)
	}

	if len(payload) <= 160 {
		return 1
	}

	parts := int(math.Ceil(float64(len(payload)) / 153))
	return int64(parts)

}

// GetSMSParts : Return numbers of SMS in multipart sms to be sent.
func GetSMSParts(message string) int64 {
	charSet := getCharSet(message)
	charsInMessage := utf8.RuneCountInString(message)

	if charSet == gsmCharset7Bit {
		return calculate7BitLengthParts(message)
	} else if charSet == gsmCharsetUnicode {
		if charsInMessage <= 70 {
			return 1
		}
		return int64(math.Ceil(float64(charsInMessage) / 67))
	}

	return -1
}
