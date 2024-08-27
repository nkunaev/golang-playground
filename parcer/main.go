package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

// Defile logging params
	f, err := os.OpenFile("/tmp/parcer-app.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error openning log file: %v", err)
	}
	defer f.Close()
	writeLog := io.MultiWriter(os.Stdout, f)

// Web block
	c := colly.NewCollector( func(c *colly.Collector) {
		colly.AllowedDomains("parsemachine.com")
	}	)

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			log.SetOutput(writeLog)
			log.Printf("Site parsing error: status code: %v, error: %v", r.StatusCode, err)
		}
	})

	c.OnHTML(".product-card", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		title := e.ChildText(".card-title")

		fmt.Println(link, title)
	})

	c.Visit("https://parsemachine.com/sandbox/catalog/?page=100")
}