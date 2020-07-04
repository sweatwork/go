// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
// Prints the names of all files in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	duplicate := make(map[string]int)
	for line, n := range counts {
		if n > 1 {
			// each duplicate line in files
			duplicate[line] = n
		}
	}

	fmt.Println("duplicated lines with count -", duplicate)

	if len(files) != 0 {
		filenames := countFiles(duplicate, files)
		if len(filenames) == 0 {
			fmt.Println("No file contains each duplicated line")
			return
		}
		fmt.Println("File(s) with each duplicated line: ", filenames)
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignoring potential errors from input.Err()
}

func countFiles(duplicate map[string]int, files []string) []string {
	filenames := []string{}
	for _, file := range files {
		fdata := make(map[string]int)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			fdata[line]++
		}
		fdup := make(map[string]int)
		for key, value := range fdata {
			if value > 1 {
				fdup[key] = value
			}
		}
		var dontaddfile = false
		for k := range duplicate {
			_, contains := fdup[k]
			if !contains {
				dontaddfile = true
				break
			}
		}
		if !dontaddfile {
			filenames = append(filenames, file)
		}

	}

	return filenames
}
