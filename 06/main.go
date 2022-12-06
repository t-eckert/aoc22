package main

import (
	"fmt"
	"os"
)

// Sample cases
var cases = []string{
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func main() {
	raw, err := os.ReadFile("./06/puzzle.input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	signal := string(raw)

	fmt.Printf("Part 1: Start of packet: %d\n", start(signal, 4))
	fmt.Printf("Part 2: Start of message: %d\n", start(signal, 14))
}

func start(signal string, window int) int {
	packetStart := 0

	for i := 0; i < len(signal)-window; i++ {
		if isUnique(signal[i : i+window]) {
			packetStart = i + window
			return packetStart
		}
	}

	return packetStart
}

func isUnique(signal string) bool {
	for i, a := range signal {
		for _, b := range signal[i+1:] {
			if a == b {
				return false
			}
		}
	}

	return true
}
