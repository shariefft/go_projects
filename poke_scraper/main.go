package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a new Colly collector
	c := colly.NewCollector()

	// Set custom HTTP transport to skip TLS verification
	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	// Log when a request is made
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	// Define what to do when an HTML element is found
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract data from the element (e.g., href attribute)
		link := e.Attr("href")
		fmt.Println("Found link:", link)
	})

	// Log any errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Start scraping by visiting a URL
	err := c.Visit("https://www.serebii.net/")
	if err != nil {
		log.Fatal(err)
	}
}

