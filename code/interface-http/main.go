package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	robots := "http://google.com/robots.txt"
	url := "http://google.com"

	fmt.Println("Robots url:", robots)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	_, _ = io.Copy(lw, resp.Body)

	// content, err := ioutil.ReadAll(resp.Body)
	//
	// _ = resp.Body.Close()
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("%s", content)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
