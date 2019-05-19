package main

import (
	"modules/scanner"

	"fmt"
	"time"
)

var (
	website = "http://www.lianhetech.com/"
)

func main() {
	var start = time.Now().Unix()
	scanner.Scan(website)
	var end = time.Now().Unix()
	fmt.Println("Time: ", end-start, "s")
}
