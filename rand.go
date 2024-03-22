package main

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strings"
)

type RandStrct struct {
	NumberOfBytes int
	FileFlagSet   bool   // set true to write output to file
	FileName      string // output filename
}

func RandomBytes(r *RandStrct) (string, error) {
	byt := make([]byte, r.NumberOfBytes)
	if _, err := rand.Read(byt); err != nil {
		return "", err
	}

	output := hex.EncodeToString(byt)
	result := strings.TrimSpace(output)

	if r.FileFlagSet {
		f, err := os.OpenFile(r.FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return "", err
		}

		if _, err := f.WriteString(result); err != nil {
			return "", err
		}

		return result, nil
	}

	return result, nil
}
