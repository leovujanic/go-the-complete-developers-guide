package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("abcdef")

	buf := make([]byte, 4)
	for {
		// read 4 bytes to buf
		n, err := r.Read(buf)
		fmt.Println(n, err, buf[:n])
		if err == io.EOF {
			break
		}
	}

	// read exactly len(buf2) bytes into buf2

	r2 := strings.NewReader("abcdef")
	buf2 := make([]byte, 4)

	if _, err := io.ReadFull(r2, buf2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf2)

	if _, err := io.ReadFull(r, buf2); err != nil {
		fmt.Println(err)
	}

	// read everything

	r3 := strings.NewReader("ab")
	var buf3 []byte
	// buf3 := make([]byte, 4)

	buf3, err := ioutil.ReadAll(r3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf3)

	str := "Counting  words in sentence"

	fmt.Println(countWords(str))
}

func countWords(s string) int {

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords) // split

	cnt := 0
	for scanner.Scan() {
		cnt += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return cnt
}
