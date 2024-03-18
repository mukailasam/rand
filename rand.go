package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type RandStrct struct {
	NumberofBytes int    // number of bytes
	FileFlagSet   bool   // set true to write output to file
	FileName      string // output filename
}

func RandomBytes(r *RandStrct) (int, error) {
	byt := make([]byte, r.NumberofBytes)
	n, err := rand.Read(byt)
	if err != nil {
		return 0, err
	}

	output := hex.EncodeToString(byt)
	result := strings.TrimSpace(output)

	if r.FileFlagSet {
		f, err := os.OpenFile(r.FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return 0, err
		}
		n, err = f.WriteString(result)
		if err != nil {
			return 0, err
		}
		return n, nil
	}

	fmt.Println("Ouput👇\n", result)

	return n, nil
}
