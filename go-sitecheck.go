package main

/** simple depth first crawler */
import (
    "fmt"
    "log"
    "github.com/PuerkitoBio/goquery"
    "net/http"
    "strings"
    "time"
    "flag"
)

var m = make(map[string]string)
var limit time.Duration
var maxDepth int
var exceeded int
var urlMap = make(map[string]string)
var filter = make([]string, 0)

//TODO: use goroutines
func LinkScrape(url string, depth int) {
    if (depth >= maxDepth) {
        return
    }

    doc, err := goquery.NewDocument(url)
    if err != nil {
        log.Fatal(err)
    }
    var links = make([]string, 0)
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        link, exists := s.Attr("href")
        if(exists && strings.HasPrefix(link, "http") && m[link] == "" && !ignoreLink(link)) {
            links = append(links, RewriteLink(link))
            m[link] = link
        }
    })
    fmt.Printf("\n%s yielded %d links, starts crawling\n", url, len(links))
    for _,link := range links {
        Scrape(link, depth)
    }
}

func ignoreLink(link string) bool {
    for _, item := range filter {
        if strings.Contains(link, item) {
            return true
        }
    }
    return false
}

func Scrape(link string, depth int) {

    var now = time.Now()
    resp, err := http.Get(link)

    if (err != nil) {
        fmt.Print("Error: ", err)
    } else {
        var duration = time.Since(now)/time.Millisecond
        if(resp.StatusCode != 200) {
            fmt.Printf("Unexpected status code=%d for link=%s\n", resp.Status, link)
        } else if(duration > limit) {
            fmt.Printf("Loaded in %dms, url=%s\n", duration, link)
            exceeded++
        }
        depth++
        LinkScrape(link, depth)
    }
}

// Replaces potential url 
func RewriteLink(link string) string {
    for k := range urlMap {
        if(strings.Contains(link, k)) {
            return strings.Replace(link, k, urlMap[k], -1)
        }
    }
    return link
}

func main() {
    var url = flag.String("url", "http://nrk.no", "url of site to crawl")
    var threshold = flag.Int("threshold", 100, "load time threshold")
    var depth = flag.Int("depth", 1, "depth of links to crawl")
    var mappings = flag.String("map", "", "map of link replacements")
    var ignores = flag.String("ignore", "", "map of domains to ignore")

    flag.Parse()

    if(*mappings != "") {
        for _, url := range strings.Split(*mappings, ",") {
            urlMapping := strings.Split(url, ":")
            urlMap[urlMapping[0]] = urlMapping[1]
        }
    }
    if(*ignores != "") {
        for _, domain := range strings.Split(*ignores, ",") {
            filter = append(filter, domain)
        }
    }

    limit = time.Duration(*threshold)
    maxDepth = *depth
    fmt.Printf("Crawler set up with url=%s, threshold=%d, depth=%d and url map=%s\n", *url, limit, maxDepth, *mappings)
    LinkScrape(RewriteLink(*url), 0)
    fmt.Printf("Crawled %d links, %d links exceeded limit (%dms)\n", len(m), exceeded, limit)
}
