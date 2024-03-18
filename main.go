package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	min = 16
	max = 100000
)

func main() {
	arg := flag.String("n", "", "rand -n 16\n rand -n 32 \n rand -n 64\n rand -n 128\n rand -n 256")
	farg := flag.String("o", "", "rand -n 1024 -o rand.txt")
	flag.Parse()

	nflag := flag.NFlag()

	r := &RandStrct{}

	switch nflag {
	case 2:
		r.FileFlagSet = true
		r.FileName = *farg
		rVal := valInRange(*arg, r)
		if rVal {
			RandomBytes(r)
		}
		return
	case 1:
		rVal := valInRange(*arg, r)
		if rVal {
			RandomBytes(r)
		}
		return

	default:
		r.NumberofBytes = 32
		RandomBytes(r)
	}
}

func parseInt(s string, base int, bitSize int) (int, error) {
	val, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		fmt.Println("bad input...")
		flag.Usage()
		return 0, err
	}
	return int(val), nil
}

func valInRange(arg1 string, r *RandStrct) bool {
	rVal, err := parseInt(arg1, 10, 0)
	if err != nil {
		os.Exit(1)
	}

	r.NumberofBytes = rVal
	if r.NumberofBytes >= min && r.NumberofBytes <= max {
		return true
	}

	fmt.Printf("%d bytes is the minimum and %d bytes is the maximum. try any number in range [%d-%d]", min, max, min, max+1)
	return false
}
