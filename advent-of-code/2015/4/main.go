package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const key = "bgvyzdsv"
	fmt.Println("part1:", solve(key, "00000"))
	fmt.Println("part2:", solve(key, "000000"))
}

func solve(key, prefix string) int {
	for n := 1; ; n++ {
		md5Sum := md5.Sum([]byte(key + strconv.Itoa(n)))
		md5SumStr := fmt.Sprintf("%x", md5Sum)
		if strings.HasPrefix(md5SumStr, prefix) {
			return n
		}
	}
	return -1
}
