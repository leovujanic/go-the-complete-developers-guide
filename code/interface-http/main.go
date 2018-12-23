package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	url := "http://google.com/robots.txt"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	content, err := ioutil.ReadAll(resp.Body)

	_ = resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", content)

}
