package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "total run")
	const defaultN = 100000
	const defaultK = 10

	k := defaultK
	n := defaultN
	if len(os.Args) >= 2 {
		nstr := os.Args[1]
		var err error
		n, err = strconv.Atoi(nstr)
		if err != nil {
			n = defaultN
		}
	}

	if len(os.Args) >= 3 {
		kstr := os.Args[2]
		var err error
		k, err = strconv.Atoi(kstr)
		if err != nil {
			k = defaultK
		}
	}
	for i := 0; i < k; i++ {
		findPrimeNumber(n)
	}
}

func findPrimeNumber(n int) int64 {
	defer timeTrack(time.Now(), "single")
	count := 0
	var a int64 = 2
	for count < n {
		var b int64 = 2
		prime := 1
		for b*b <= a {
			if a%b == 0 {
				prime = 0
				break
			}
			b++
		}
		if prime > 0 {
			count++
		}
		a++
	}
	a--
	return a
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
