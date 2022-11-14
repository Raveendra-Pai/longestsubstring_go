package main

import (
	"fmt"
	"strings"
)

func FindLongestSubstring(str string) string {
	count := 0
	maxCount := 0
	mapF := make(map[byte]bool)
	var strBuilder strings.Builder
	strBuilder.Grow(64) //allocate 8 bytes to stop reallocation
	longestSubstring := ""

	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if _, ok := mapF[str[j]]; !ok {
				count++
				mapF[str[j]] = true
				strBuilder.WriteByte(str[j])
			} else {
				break
			}
		}
		if count >= maxCount {
			maxCount = count
			longestSubstring = strBuilder.String()
		}
		count = 0
		mapF = make(map[byte]bool)
		strBuilder.Reset()
	}
	return longestSubstring
}

func main() {

	str := "RAVIISDUMBGUY"

	longestString := FindLongestSubstring(str)
	fmt.Println("Longest string: ", longestString)

}
