package main

import (
	"fmt"
	"math/rand"
)

type fourLevelHierarchicalPageAddress struct {
	raw             string // all 64-bits
	readableVersion string // all 64-bits, hyphen-delimited
	ignoredBits     string // leading 16 bits, bits 63-48, are ignored
	levelFour       string // bits 39-47
	levelThree      string // bits 38-30
	levelTwo        string // bits 29-21
	levelOne        string // bits 20-12
	physicalAddress string // bits 11-0
}

func main() {
	// Generate a random 64-bit address as a byte slice
	address := make([]byte, 8)
	for i := 0; i < 8; i++ {
		// Generate a random byte (0 to 255)
		randomByte := byte(rand.Intn(256))
		// Assign the random byte to the address byte slice
		address[i] = randomByte
	}

	pageAddr := &fourLevelHierarchicalPageAddress{}
	for i, b := range address {
		delimiter := "-"
		if i == (len(address) - 1) {
			delimiter = ""
		}
		pageAddr.readableVersion += fmt.Sprintf("%08b%s", b, delimiter)
		pageAddr.raw += fmt.Sprintf("%08b", b)
	}
	fmt.Printf("Full 64-bits are: %s\n", pageAddr.readableVersion)
	bitIdx := 63
	for _, char := range pageAddr.raw {
		if bitIdx <= 63 && bitIdx >= 48 {
			pageAddr.ignoredBits += string(char)
		} else if bitIdx <= 47 && bitIdx >= 39 {
			pageAddr.levelFour += string(char)
		} else if bitIdx <= 38 && bitIdx >= 30 {
			pageAddr.levelThree += string(char)
		} else if bitIdx <= 29 && bitIdx >= 21 {
			pageAddr.levelTwo += string(char)
		} else if bitIdx <= 20 && bitIdx >= 12 {
			pageAddr.levelOne += string(char)
		} else if bitIdx <= 11 {
			pageAddr.physicalAddress += string(char)
		}
		bitIdx -= 1
	}
	fmt.Printf("ignored bits 63-48: %s\n", pageAddr.ignoredBits)
	fmt.Printf("level 4, bits 39-47: %s\n", pageAddr.levelFour)
	fmt.Printf("level 3, bits 38-30: %s\n", pageAddr.levelThree)
	fmt.Printf("level 2, bits 29-21: %s\n", pageAddr.levelTwo)
	fmt.Printf("level 1, bits 20-12: %s\n", pageAddr.levelOne)
	fmt.Printf("physical address, bits 11-0: %s\n", pageAddr.physicalAddress)
}
