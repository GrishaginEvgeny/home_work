package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func isScreeningByWithoutRepeat(prevStringedRune string, errActual error, stringedRune string) bool {
	return prevStringedRune == "\\" && (errActual == nil || stringedRune == "\\")
}

func isError(errActual error, stringedRune string, errPrev error, prevPrevStringedRune string,
	prevStringedRune string, index int,
) bool {
	return errActual == nil && stringedRune != "\\" &&
		(errPrev == nil && prevPrevStringedRune != "\\" && prevStringedRune != "\\" || index == 0)
}

func getNextRune(index int, r string) string {
	var nextRune string
	if index != len(r)-1 {
		nextRune = string(r[index+1])
	} else {
		nextRune = ""
	}
	return nextRune
}

func getPrevRune(index int, r string) string {
	var prevRune string
	if index != 0 {
		prevRune = string(r[index-1])
	} else {
		prevRune = ""
	}
	return prevRune
}

func getPrevPrevRune(index int, r string) string {
	var prevPrevRune string
	if index > 1 {
		prevPrevRune = string(r[index-2])
	} else {
		prevPrevRune = ""
	}
	return prevPrevRune
}

func getIncrementOfIndex(index int, size int) int {
	var inc int
	if index+2 > size-1 {
		inc = 1
	} else {
		inc = 2
	}
	return inc
}

func Unpack(line string) (string, error) {
	var result strings.Builder
	for i := 0; i < len(line); i++ {
		prevStringedRune := getPrevRune(i, line)
		prevPrevStringedRune := getPrevPrevRune(i, line)
		nextStringedRune := getNextRune(i, line)
		stringedRune := string(line[i])
		intAct, errActual := strconv.Atoi(stringedRune)
		_, errPrev := strconv.Atoi(prevStringedRune)
		intNext, errNext := strconv.Atoi(nextStringedRune)
		switch {
		case prevStringedRune == "\\" && (errActual == nil || stringedRune == "\\") && errNext == nil:
			{
				result.WriteString(strings.Repeat(stringedRune, intNext))
				i += getIncrementOfIndex(i, len(line))
				break
			}
		case isScreeningByWithoutRepeat(prevStringedRune, errActual, stringedRune):
			{
				result.WriteString(stringedRune)
				i++
				break
			}
		case isError(errActual, stringedRune, errPrev, prevPrevStringedRune, prevStringedRune, i):
			{
				return "", ErrInvalidString
			}
		case errActual == nil && stringedRune != "\\" && prevPrevStringedRune != "\\":
			{
				result.WriteString(strings.Repeat(prevStringedRune, intAct))
				break
			}
		case errNext != nil && prevStringedRune != "\\" && stringedRune != "\\":
			{
				result.WriteString(stringedRune)
				break
			}
		}
	}
	return result.String(), nil
}
