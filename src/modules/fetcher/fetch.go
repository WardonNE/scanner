package fetcher

import (
	"github.com/PuerkitoBio/goquery"

	"fmt"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

type Fetcher struct {
	urls               []string
	contentTree        map[string][]content
	ignoreurls         []string
	targetElement      string
	targetElementIndex int
	targetElementsTree map[string]string
	sync.RWMutex
}

func NewFetcher(urls []string) *Fetcher {
	this := new(Fetcher)
	this.urls = urls
	this.targetElementIndex = 0
	return this
}

// func NewFetcherWithConfigFile(configfile string) *Fetcher {

// }

func (this *Fetcher) SetIgnoreUrls(urls []string) *Fetcher {
	this.ignoreurls = urls
	return this
}

func (this *Fetcher) GetIgnoreUrls() []string {
	return this.ignoreurls
}

func (this *Fetcher) SetTargetElement(ele string, index int) *Fetcher {
	this.targetElement = ele
	this.targetElementIndex = index
	this.contentTree = make(map[string][]content)
	return this
}

func (this *Fetcher) GetTargetElement() string {
	return fmt.Sprintf("%s[%d]", this.targetElement, this.targetElementIndex)
}

func (this *Fetcher) SetTargetElementsTree(eleTree map[string]string) *Fetcher {
	this.targetElementsTree = eleTree
	return this
}

func (this *Fetcher) getTargetElementsTree() map[string]string {
	return this.targetElementsTree
}

func (this *Fetcher) ListContentTree() map[string][]content {
	return this.contentTree
}

func (this *Fetcher) FetchAll() *Fetcher {
	for _, url := range this.urls {
		waitgroup.Add(1)
		go this.Fetch(url)
	}
	waitgroup.Wait()
	return this
}

func (this *Fetcher) Fetch(url string) *Fetcher {
	fmt.Println("Start Time: ", time.Now(), " Url: ", url)
	this.Lock()
	defer this.Unlock()
	defer waitgroup.Done()
	if _, ok := this.contentTree[url]; ok {
		return this
	}
	dom, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	var title = dom.Find("title").Text()
	targetElement, ok := this.targetElementsTree[url]
	if !ok {
		if this.targetElementIndex == -1 {
			dom.Find(this.targetElement).Each(func(i int, s *goquery.Selection) {
				cont := new(content)
				cont.link = url
				cont.title = title
				html, err := goquery.OuterHtml(s)
				if err != nil {
					panic(err)
				}
				cont.html = html
				this.contentTree[url] = append(this.contentTree[url], *cont)
			})
		} else {
			cont := new(content)
			cont.link = url
			cont.title = title
			html, err := goquery.OuterHtml(dom.Find(this.targetElement).Eq(this.targetElementIndex))
			if err != nil {
				panic(err)
			}
			cont.html = html
			// fmt.Println(html)
			this.contentTree[url] = append(this.contentTree[url], *cont)
		}
	} else {
		if this.targetElementIndex == -1 {
			dom.Find(targetElement).Each(func(i int, s *goquery.Selection) {
				cont := new(content)
				cont.link = url
				cont.title = title
				html, err := goquery.OuterHtml(s)
				if err != nil {
					panic(err)
				}
				cont.html = html
				this.contentTree[url] = append(this.contentTree[url], *cont)
			})
		} else {
			cont := new(content)
			cont.link = url
			cont.title = title
			html, err := goquery.OuterHtml(dom.Find(targetElement).Eq(this.targetElementIndex))
			if err != nil {
				panic(err)
			}
			cont.html = html
			this.contentTree[url] = append(this.contentTree[url], *cont)
		}
	}
	return this
}
