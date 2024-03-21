package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	MinBytes = 16
	MaxBytes = 100000
)

func main() {
	numberOfBytesArg := flag.Int("n", 32, "Usage: rand -n <numberOfBytes>")
	outputArg := flag.String("o", "", "Usage: rand -n <numberOfBytes> -o <outputFile>")
	flag.Parse()

	var numberOfBytes int
	switch numberOfBytesArg {
	case nil:
		numberOfBytes = 32
	default:
		numberOfBytes = *numberOfBytesArg
	}

	if numberOfBytes < MinBytes || numberOfBytes > MaxBytes {
		fmt.Printf("%d bytes is the minimum and %d bytes is the maximum. try any number in range [%d-%d]\n", MinBytes, MaxBytes, MinBytes, MaxBytes+1)
		os.Exit(1)
	}

	randStruct := &RandStrct{
		NumberOfBytes: numberOfBytes,
		FileFlagSet:   false,
	}

	if outputArg != nil && *outputArg != "" {
		randStruct.FileFlagSet = true
		randStruct.FileName = *outputArg
	}

	randomBytes, err := RandomBytes(randStruct)
	if err != nil {
		fmt.Printf("failed to generate random bytes: %s", err)
		os.Exit(1)
	}

	fmt.Println(randomBytes)
}
