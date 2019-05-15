package main

import (
	"./modules/scanner"
)

var (
	website = "http://www.stec.net/"
)

func main() {
	scanner.Scan(website)
}
