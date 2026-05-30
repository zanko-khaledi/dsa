package algorithmes

import (
	"math/rand"
)

func RandomStr(length int) string {

	str := new(string)

	chars := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'j', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	for i := 0; i < length; i++ {
		*str += string(chars[rand.Intn(len(chars)-1)])
	}

	return *str
}

func Frequency(str string) map[string]int {

	freq := make(map[string]int, 0)

	for _, char := range str {
		freq[string(char)]++
	}

	return freq
}

func MostRepeatedChar(str string) (string, int) {

	freq := Frequency(str)

	var chars []rune
	maxCount := 0

	for _, n := range freq {
		if n > maxCount {
			maxCount = n
		}
	}

	for ch, count := range freq {
		if maxCount == count {
			chars = append(chars, []rune(ch)...)
		}
	}

	return string(chars), maxCount
}
