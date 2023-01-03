package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(line string) (string, error) {
	var result strings.Builder
	r := []rune(line)
	for i := 0; i < len(r); i++ {
		prevStringedRune := ""
		prevPrevStringedRune := ""
		nextStringedRune := ""
		stringedRune := string(r[i])
		if i != 0 {
			prevStringedRune = string(r[i-1])
		}
		if i != len(r)-1 {
			nextStringedRune = string(r[i+1])
		}
		if i > 1 {
			prevPrevStringedRune = string(r[i-2])
		}

		intAct, errActual := strconv.Atoi(stringedRune)
		_, errPrev := strconv.Atoi(prevStringedRune)
		intNext, errNext := strconv.Atoi(nextStringedRune)

		if prevStringedRune == "\\" && (errActual == nil || stringedRune == "\\") && errNext == nil {
			result.WriteString(strings.Repeat(stringedRune, intNext))
			if i+2 > len(r)-1 {
				i++
			} else {
				i += 2
			}
		} else if prevStringedRune == "\\" && (errActual == nil || stringedRune == "\\") {
			result.WriteString(stringedRune)
			i++
		} else if errActual == nil && stringedRune != "\\" {
			if i == 0 {
				return "", ErrInvalidString
			}
			if errPrev == nil && prevPrevStringedRune != "\\" && prevStringedRune != "\\" {
				return "", ErrInvalidString
			}
			if prevPrevStringedRune != "\\" {
				result.WriteString(strings.Repeat(prevStringedRune, intAct))
			}
		} else {
			if errNext != nil && prevStringedRune != "\\" && stringedRune != "\\" {
				result.WriteString(stringedRune)
			}
		}
	}
	return result.String(), nil
}
