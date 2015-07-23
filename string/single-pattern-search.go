package string

/*
import (
	"fmt"
)
*/

func KMPSearch(source string, pattern string) int {
	si := 0
	pi := 0
	nextPatternIndexArray := generateNextPatternIndexArray(pattern)
	for si < len(source) && pi < len(pattern) {
		if pi == -1 || ([]byte(source)[si] == []byte(pattern)[pi]) {
			si++
			pi++
		} else {
			pi = nextPatternIndexArray[pi]
		}
	}
	if pi == len(pattern) {
		return si - pi
	} else {
		return -1
	}
}

func generateNextPatternIndexArray(pattern string) []int {
	length := len(pattern)
	array := make([]int, length)
	array[0] = -1
	prefixI := -1
	suffixI := 0
	for suffixI < length - 1 {
		if prefixI == -1 || ([]byte(pattern)[prefixI] == []byte(pattern)[suffixI]) {
			prefixI++
			suffixI++
			if []byte(pattern)[prefixI] != []byte(pattern)[suffixI] {
				array[suffixI] = prefixI
			} else {
				array[suffixI] = array[prefixI]
			}
		} else {
			prefixI = array[prefixI]
		}
	}
	return array
}

func BMSearch(source string, pattern string) int {
	//TODO
	return -1
}
