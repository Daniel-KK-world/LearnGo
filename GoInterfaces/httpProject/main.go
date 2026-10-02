package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(resp) this just prints the whole ball of resp, which is not very useful.

	defer resp.Body.Close()

	// 1. Create an empty byte slice with enough space to hold the HTML
	bs := make([]byte, 99999)

	// 2. Pass the slice into the Read function dictated by the io.Reader interface
	resp.Body.Read(bs)

	// 3. Convert the byte slice to a string and print it
	fmt.Println(string(bs))

}
