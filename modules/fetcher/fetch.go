package scanner

import (
	"github.com/PuerkitoBio/goquery"
	// "strings"
)

func FetchContentByUrls(urls []string, element string, index int) []map[string][]string {
	var result []map[string][]string
	for _, url := range urls {
		url_content := make(map[string][]string)
		url_content = FetchContent(url, element, index)
		result = append(result, url_content)
	}
	return result
}

func FetchContent(url string, element string, index int) map[string][]string {
	var result = make(map[string][]string)
	if dom, err := goquery.NewDocument(url); err == nil {
		if index == -1 {
			content := dom.Find(element).Eq(index).Text()
			result[url] = append(result[url], content)
		} else {
			dom.Find(element).Each(func(i int, s *goquery.Selection) {
				content := s.Text()
				result[url] = append(result[url], content)
			})
		}
	}
	return result
}
