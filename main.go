package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Soal no 1
	fmt.Println("===== Soal NO 1 ======")
	word := "We Always Mekar"
	charMap := MapCharLower(word)
	fmt.Printf("%v\n", charMap)

	// Soal no 2
	fmt.Println("===== Soal NO 2 ======")
	wordSlice := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	wordString := SliceToString(wordSlice)
	charMap2 := MapChar(wordString)
	sortedString := sortMapKey(charMap2)
	fmt.Println(charMap2)
	fmt.Println(sortedString)
}

func sortMapKey(mapString map[string]int) string {
	stringSlice := MapToArray(mapString)
	// loop string slice
	// sort stringslice based on the value of map[stringSlice[i]]
	// sort descendingly by the value
	// if there are two char that are same i.e (a and A) that have the same value
	// put the capital letter first
	// if there are two different char that have the same value sort alphabetically
	// but put capital letter first

	sort.Slice(stringSlice, func(i, j int) bool {
		if mapString[stringSlice[i]] == mapString[stringSlice[j]] {
			if stringSlice[i] == strings.ToUpper(stringSlice[i]) && stringSlice[j] == strings.ToLower(stringSlice[j]) {
				return true
			} else if stringSlice[i] == strings.ToLower(stringSlice[i]) && stringSlice[j] == strings.ToUpper(stringSlice[j]) {
				return false
			} else {
				return stringSlice[i] < stringSlice[j]
			}
		}
		return mapString[stringSlice[i]] > mapString[stringSlice[j]]
	})

	wordString := strings.Join(stringSlice, "")
	return wordString
}

func MapCharLower(words string) map[string]int {
	// remove all whitespace inside the string
	words = strings.ReplaceAll(words, " ", "")
	// turn strings into lower case
	words = strings.ToLower(words)

	// initiate empty map[string]int
	countMap := make(map[string]int)

	// turn words into slice of rune
	for _, char := range words {
		_, ok := countMap[string(char)]
		if ok {
			countMap[string(char)] += 1
		} else {
			countMap[string(char)] = 1
		}
	}

	return countMap
}

func MapChar(words string) map[string]int {
	// remove all whitespace inside the string
	words = strings.ReplaceAll(words, " ", "")
	// initiate empty map[string]int
	countMap := make(map[string]int)

	// turn words into slice of rune
	for _, char := range words {
		_, ok := countMap[string(char)]
		if ok {
			countMap[string(char)] += 1
		} else {
			countMap[string(char)] = 1
		}
	}

	return countMap
}

func SliceToString(slices []string) string {
	word := strings.Join(slices, "")
	return word
}

func MapToArray(mapChar map[string]int) []string {
	// initiate empty slice of string
	var stringSlice []string
	// loop mapChar
	for k := range mapChar {
		stringSlice = append(stringSlice, k)
	}
	return stringSlice
}
