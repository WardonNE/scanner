package scanner

import (
	"github.com/PuerkitoBio/goquery"

	"fmt"
	"strings"
)

var (
	wait_scan_urls  []string
	been_found_urls map[string]int
	total_urls      []string
	wait_fetch_urls []string
	root_url        string
)

func Scan(root string) []string {
	root_url = strings.TrimRight(root, "/")
	been_found_urls = make(map[string]int)
	total_urls = []string{root_url}
	wait_fetch_urls = []string{root_url}
	pushUrl2WaitScan(total_urls)
	for {
		if 0 >= len(wait_scan_urls) {
			break
		}
		url := wait_scan_urls[0]
		wait_scan_urls = wait_scan_urls[1:]
		urls := findAllHrefs(url)
		pushUrl2WaitScan(urls)
	}
	return wait_fetch_urls
}

func pushUrl2WaitScan(urls []string) {
	for _, url := range urls {
		url = filterAnchorPoint(url)
		ok := isFoundUrl(url)
		if !ok {
			fmt.Printf("new url:(%s)\n", url)
			wait_scan_urls = append(wait_scan_urls, url)
			been_found_urls[url] = 1
			wait_fetch_urls = append(wait_fetch_urls, url)
		} else {
			been_found_urls[url]++
		}
	}
}

func isFoundUrl(url string) bool {
	_, ok := been_found_urls[strings.TrimRight(url, "/")]
	return ok
}

func findAllHrefs(url string) []string {
	var result []string
	if dom, err := goquery.NewDocument(url); err == nil {
		dom.Find("a").Each(func(i int, s *goquery.Selection) {
			if href, exists := s.Attr("href"); exists == true {
				if valid := checkUrl(href); valid == true {
					if strings.Contains(href, root_url) {
						result = append(result, href)
					} else {
						result = append(result, root_url+href)
					}
				}
			}
		})
	}
	return result
}

func checkUrl(url string) bool {
	var valid bool = true
	for _, r := range equal {
		if url == r {
			valid = false
			return false
		}
	}
	for _, r := range contains {
		if strings.Contains(url, r) {
			valid = false
			return valid
		}
	}
	for _, r := range internal {
		if strings.Contains(url, r) && !strings.Contains(url, root_url) {
			valid = false
			return valid
		}
	}
	return valid
}

func filterAnchorPoint(url string) string {
	urlparts := strings.Split(url, "#")
	return urlparts[0]
}
