package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"net/http"
	"crypto/tls"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id:", err)
			return
		}

		factDesc := element.Text
		if factDesc == "" {
			log.Println("Empty fact description")
			return
		}

		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
		log.Println("Added fact:", fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")
	writeJSON(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file:", err)
		return
	}

	err = ioutil.WriteFile("rhinofacts.json", file, 0644)
	if err != nil {
		log.Println("Unable to write json file:", err)
	}
}
