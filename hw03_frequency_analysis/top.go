package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordRate struct {
	word  string
	entry int
}

func stringsArrayToLower(stringArray []string) []string {
	var res []string
	for i := 0; i < len(stringArray); i++ {
		res = append(res, strings.ToLower(stringArray[i]))
	}
	return res
}

func sortedWordRateArrayToStringArray(wordRateArray []wordRate) []string {
	var res []string
	sort.Slice(wordRateArray, func(i, j int) bool {
		if wordRateArray[j].entry == wordRateArray[i].entry {
			return wordRateArray[j].word > wordRateArray[i].word
		}
		return wordRateArray[j].entry < wordRateArray[i].entry
	})
	for i := 0; i < len(wordRateArray); i++ {
		res = append(res, wordRateArray[i].word)
	}
	return res
}

func isInWordRateArrayByIndex(word string, wordRateArray []wordRate) int {
	for j := 0; j < len(wordRateArray); j++ {
		if wordRateArray[j].word == word {
			return j
		}
	}
	return -1
}

func clearText(text string) string {
	deletedChars := []string{",", ".", ";", "?", "!", "'", `"`, "-"}
	for i := 0; i < len(deletedChars); i++ {
		if strings.Contains(text, deletedChars[i]) {
			text = strings.ReplaceAll(text, deletedChars[i], "")
		}
	}
	return text
}

func Top10(text string) []string {
	var wordRateArray []wordRate
	r := regexp.MustCompile("\\s+")
	clearedTextOfTabs := r.ReplaceAllString(text, " ")
	clearedText := clearText(clearedTextOfTabs)
	splitedText := strings.Split(clearedText, " ")
	var lowerCasedStringsArray = stringsArrayToLower(splitedText)
	for i := 0; i < len(lowerCasedStringsArray); i++ {
		if isInWordRateArrayByIndex(lowerCasedStringsArray[i], wordRateArray) == -1 {
			if lowerCasedStringsArray[i] != "" {
				wordRateArray = append(wordRateArray, wordRate{word: lowerCasedStringsArray[i], entry: 1})
			}
		} else {
			wordRateArray[isInWordRateArrayByIndex(lowerCasedStringsArray[i], wordRateArray)].entry++
		}
	}
	result := sortedWordRateArrayToStringArray(wordRateArray)
	if len(result) >= 10 {
		return result[0:10]
	} else {
		return result
	}
}
