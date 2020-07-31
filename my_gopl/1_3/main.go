package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start).Seconds())
}
