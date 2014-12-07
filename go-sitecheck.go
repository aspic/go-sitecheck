package main

import (
    "fmt"
    "log"
    "github.com/PuerkitoBio/goquery"
    "net/http"
    "strings"
    "time"
    "flag"
)

func LinkScrape(url string, thres int) {
    doc, err := goquery.NewDocument(url)
    threshold := time.Duration(thres)
    fmt.Printf("Scraper set up with url=%s and threshold=%d\n", url, thres)

    if err != nil {
        log.Fatal(err)
    }
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
    link, exists := s.Attr("href")

    if(exists && strings.HasPrefix(link, "http")) {
        now := time.Now()
        resp, err := http.Get(link)

        if (err != nil) {
            fmt.Print("Error: ", err)
        } else {
            duration := time.Since(now)/time.Millisecond
            if(resp.StatusCode != 200) {
                fmt.Printf("Unexpected status code=%d for link=%s\n", resp.Status, link)
            } else if(duration > threshold) {
                fmt.Printf("Loaded in %dms, url=%s\n", duration, link)
            }
        }
    }
  })
}

func main() {
    var url = flag.String("url", "http://nrk.no", "URL of site to check")
    var threshold = flag.Int("threshold", 100, "Load time threshold")

    flag.Parse()
    LinkScrape(*url, *threshold)
}
