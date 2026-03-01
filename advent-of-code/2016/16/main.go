package main

import "fmt"

func main() {
	// test input
	// state := []byte("10000")
	// diskSize1 := 20

	state := []byte("11110010111001001")
	diskSize1 := 272
	diskSize2 := 35651584
	fmt.Println("part 1: ", string(solve(state, diskSize1)))
	fmt.Println("part 2: ", string(solve(state, diskSize2)))
}

func solve(data []byte, diskSize int) []byte {
	for len(data) < diskSize {
		data = generate(data)
	}
	data = data[:diskSize]
	return checkSum(data)
}

func generate(data []byte) []byte {
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)
	reverse(dataCopy)
	invert(dataCopy)
	data = append(data, '0')
	data = append(data, dataCopy...)
	return data
}

func reverse(data []byte) {
	s, e := 0, len(data)-1
	for s < e {
		data[s], data[e] = data[e], data[s]
		s++
		e--
	}
}

func invert(data []byte) {
	for i := range data {
		if data[i] == '1' {
			data[i] = '0'
		} else {
			data[i] = '1'
		}
	}
}

func checkSum(data []byte) []byte {
	cs := checkSumRound(data)
	for len(cs)%2 == 0 {
		cs = checkSumRound(cs)
	}
	return cs
}

func checkSumRound(data []byte) []byte {
	cs := make([]byte, len(data)/2)
	for i := 0; i < len(data)-1; i += 2 {
		if data[i] == data[i+1] {
			cs[i/2] = '1'
		} else {
			cs[i/2] = '0'
		}
	}
	return cs
}
