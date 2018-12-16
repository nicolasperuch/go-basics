package main

import(
	"fmt"
	"net/http"
	"os"
)

var TARGET_URL = "http://google.com"

func main() {
	resp, err := http.Get(TARGET_URL)
	if nil != err {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}