package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ipParts struct {
	as []string
	bs []string
}

func main() {
	ips := readInput()
	fmt.Println("part1:", solvePart1(ips))
	fmt.Println("part2:", solvePart2(ips))
}

func readInput() []ipParts {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var ips []ipParts
	for s.Scan() {
		ip, err := parseIP(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		ips = append(ips, ip)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return ips
}

func parseIP(s string) (ipParts, error) {
	var ip ipParts
	for len(s) > 0 {
		o := strings.IndexByte(s, '[')
		c := strings.IndexByte(s, ']')
		if o == -1 {
			ip.as = append(ip.as, s)
			break
		}
		if o > 0 {
			ip.as = append(ip.as, s[:o])
		}
		ip.bs = append(ip.bs, s[o+1:c])
		s = s[c+1:]
	}
	return ip, nil
}

func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func supportsTLS(ip ipParts) bool {
	found := false
	for _, p := range ip.as {
		if containsABBA(p) {
			found = true
		}
	}
	if !found {
		return false
	}
	for _, p := range ip.bs {
		if containsABBA(p) {
			return false
		}
	}
	return true
}

func solvePart1(ips []ipParts) int {
	supportTLSCnt := 0
	for _, ip := range ips {
		if supportsTLS(ip) {
			supportTLSCnt++
		}
	}
	return supportTLSCnt
}

func supportsSSL(ip ipParts) bool {
	for _, a := range ip.as {
		for i := range len(a) - 2 {
			if a[i] == a[i+2] && a[i] != a[i+1] {
				for _, b := range ip.bs {
					s := string([]byte{a[i+1], a[i], a[i+1]})
					if strings.Contains(b, s) {
						return true
					}
				}
			}
		}
	}
	return false
}

func solvePart2(ips []ipParts) int {
	supportSSLCnt := 0
	for _, ip := range ips {
		if supportsSSL(ip) {
			supportSSLCnt++
		}
	}
	return supportSSLCnt
}
