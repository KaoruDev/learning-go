package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	res, _ := http.Get("https://www.example.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	//fmt.Println(page) // cannot use this because page is a byte array
	fmt.Printf("%s", page)
}
