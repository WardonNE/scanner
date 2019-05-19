package main

import (
	"fmt"
	"modules/fetcher"
	"modules/scanner"
	// "time"
)

var (
	website = "http://www.lianhetech.com/"
)

func main() {
	urls := scanner.Scan(website)

	contentTree := fetcher.NewFetcher(urls).SetTargetElement(".innercontent", 0).FetchAll().ListContentTree()
	fmt.Println(contentTree)
}
