package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const input = "uqwqemis"
	fmt.Println("part1:", solvePart1(input))
	fmt.Println("part2:", solvePart2(input))
}

func solvePart1(input string) string {
	var password []byte
	for i := 0; len(password) < 8; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if strings.HasPrefix(hash, "00000") {
			password = append(password, hash[5])
		}
	}
	return string(password)
}

func solvePart2(input string) string {
	var password [8]byte
	filled := 0
	for i := 0; filled < 8; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))
		if strings.HasPrefix(hash, "00000") && hash[5] >= '0' && hash[5] < '8' && password[hash[5]-'0'] == 0 {
			password[hash[5]-'0'] = hash[6]
			filled++
		}
	}
	return string(password[:])
}
