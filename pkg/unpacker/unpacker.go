package unpacker

import (
	"errors"
	"strings"
	"unicode"
)

func UnpackString(packedString string) (string, error) {
	if len(packedString) == 0 {
		return "", nil
	}
	packedStringRunes := []rune(packedString)
	builder := strings.Builder{}
	escapeMode := false
	var lastRune rune
	i := 0
	for i < len(packedStringRunes) {
		currentRune := packedStringRunes[i]
		if currentRune == '\\' {
			if escapeMode {
				builder.WriteRune(currentRune)
				lastRune = currentRune
				escapeMode = false
			} else {
				escapeMode = true
			}
		} else if unicode.IsDigit(currentRune) {
			if escapeMode {
				builder.WriteRune(currentRune)
				lastRune = currentRune
				escapeMode = false
			} else {
				if builder.Len() == 0 {
					return "", errors.New("unpack error: wrong sequence")
				}
				num, offset, err := extractNumber(packedStringRunes, i)
				if err != nil {
					return "", err
				}
				if num != 1 {
					builder.WriteString(strings.Repeat(string(lastRune), num-1))
				}
				i += offset
			}
		} else {
			if escapeMode {
				return "", errors.New("unpack error: invalid escape sequence")
			}
			builder.WriteRune(currentRune)
			lastRune = currentRune
		}
		i++
	}
	if escapeMode {
		return "", errors.New("unpack error: unclosed escape sequence")
	}
	return builder.String(), nil
}

func extractNumber(runes []rune, numberIndex int) (int, int, error) {
	// функция для извлечения числа из строки: возвращает число, сдвиг (после которого начинается не-число) и error
	offset := 0
	num := int(runes[numberIndex] - '0')
	if numberIndex+1 < len(runes) {
		numberIndex++
		if unicode.IsDigit(runes[numberIndex]) {
			num = num*10 + int(runes[numberIndex]-'0')
			offset++
		}
	}
	if num == 0 {
		return 0, 0, errors.New("unpack error: zero in sequence")
	}
	if num > 100 {
		return 0, 0, errors.New("unpack error: too long number sequence")
	}
	return num, offset, nil
}
