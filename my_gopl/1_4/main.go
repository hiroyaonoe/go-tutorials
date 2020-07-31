package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	flist := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, flist)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "1_4: %v\n", err)
				continue
			}
			countLines(f, counts, flist)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println(flist[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, flist map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		
		flist[input.Text()]+=f.Name()+" "
	}
	// NOTE: ignoring potential errors from input.Err()
}