package fetcher

import (
	"regexp"
	"strings"
)

type content struct {
	html  string
	link  string
	title string
}

func (this *content) ToHtml() string {
	return this.html
}

func (this *content) ToText() string {
	return trimHtml(this.html)
}

// func (this *content) ToFile(filepath string) {

// }

func (this *content) GetTitle() string {
	return this.title
}

func trimHtml(html string) string {
	if lowerStr, err := regexp.Compile("\\<[\\S\\s]+?\\>"); err == nil {
		html = lowerStr.ReplaceAllStringFunc(html, strings.ToLower)
		if noStyleStr, err := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>"); err == nil {
			html = noStyleStr.ReplaceAllString(html, "")
			if noJsStr, err := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>"); err == nil {
				html = noJsStr.ReplaceAllString(html, "")
				if noTagStr, err := regexp.Compile("\\<[\\S\\s]+?\\>"); err != nil {
					noTagStr.ReplaceAllString(html, "\n")
					if noLineStr, err := regexp.Compile("\\s{2,}"); err != nil {
						noLineStr.ReplaceAllString(html, "\n")
						return strings.TrimSpace(html)
					} else {
						panic(err)
					}
				} else {
					panic(err)
				}
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
