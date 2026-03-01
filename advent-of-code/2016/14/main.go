package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	// salt := "abc" // test input
	salt := "ihaygndm"
	fmt.Println("part 1:", solvePart1(salt))
	fmt.Println("part 2:", solvePart2(salt))
}

func solvePart1(salt string) int {
	targetKeyIdx := 64
	return findKey(salt, targetKeyIdx, md5Sum)
}

func solvePart2(salt string) int {
	targetKeyIdx := 64
	return findKey(salt, targetKeyIdx, makeRepeatedHashFunc(md5Sum, 2017))
}

func findKey(salt string, targetKeyIdx int, hashFunc func(string) string) int {
	keysFound := 0
	var hashes []string
	for idx := range 1001 {
		hashes = append(hashes, hashFunc(fmt.Sprintf("%s%d", salt, idx)))
	}
	idx := 0
	for {
		hash := hashes[0]
		hashes = hashes[1:]
		if c := getRepeatedDigit3(hash); c != 0 {
			if hasOneWithRepeatedDigit5(hashes[:1000], c) {
				keysFound++
				if keysFound == targetKeyIdx {
					return idx
				}
			}
		}
		hashes = append(hashes, hashFunc(fmt.Sprintf("%s%d", salt, idx+1001)))
		idx++
	}
}

func getRepeatedDigit3(s string) byte {
CHAR:
	for i := 0; i < len(s)-2; i++ {
		for k := i + 1; k < i+3; k++ {
			if s[k] != s[i] {
				continue CHAR
			}
		}
		return s[i]
	}
	return 0
}

func hasOneWithRepeatedDigit5(ss []string, c byte) bool {
	repeated := strings.Repeat(string(c), 5)
	for _, s := range ss {
		if strings.Index(s, repeated) != -1 {
			return true
		}
	}
	return false
}

func md5Sum(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func makeRepeatedHashFunc(hashFunc func(string) string, n int) func(string) string {
	return func(s string) string {
		for range n {
			s = hashFunc(s)
		}
		return s
	}
}
