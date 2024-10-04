package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const OUTPUT_FILE = "pokemon.json"

// Pokemon holds the statistics of a Pokemon
type Pokemon struct {
	Name           string   `json:"name"`
	Number         string   `json:"number"`
	Classification string   `json:"classification"`
	Height         []string `json:"height"`
	Weight         []string `json:"weight"`
	HitPoints      int      `json:"hit_points"`
	Attack         int      `json:"attack"`
	Defense        int      `json:"defense"`
	Special        int      `json:"special"`
	Speed          int      `json:"speed"`
}

func main() {
	save := flag.Bool("save", false, "save the output to JSON")
	first := flag.Int("first", 1, "the ID of the first Pokémon to retrieve")
	last := flag.Int("last", 1, "the ID of the last Pokémon to retrieve")
	verbose := flag.Bool("verbose", false, "print the Pokémon's statistics to console")
	flag.Parse()

	log.Println("Extracting data from Serebii.net")
	firstID, lastID := validateInput(*first, *last)

	dataList := []Pokemon{}
	for pokeID := firstID; pokeID <= lastID; pokeID++ {
		data, err := extractStatistics(pokeID)
		if err != nil {
			log.Println(err)
			continue
		}
		dataList = append(dataList, data)

		if *verbose || !*save {
			displayFormatted(data)
		} else {
			log.Printf("Scraped %s %s", data.Number, data.Name)
		}
	}

	if *save {
		log.Printf("Saving to %s", OUTPUT_FILE)
		saveToJSON(dataList)
	} else {
		log.Println("All Pokémon retrieved! To save to JSON, use the --save flag")
	}
}

func extractStatistics(pokeID int) (Pokemon, error) {
	url := fmt.Sprintf("https://serebii.net/pokedex-swsh/%03d.shtml", pokeID)
	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to fetch data from %s: %v", url, err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to parse document: %v", err)
	}

	centerDivs := doc.Find("div[align='center']").Nodes
	if len(centerDivs) < 2 {
		return Pokemon{}, fmt.Errorf("failed to find expected HTML elements in %s", url)
	}

	centerPanelInfo := doc.Find("td.fooinfo")
	if centerPanelInfo.Length() < 8 {
		return Pokemon{}, fmt.Errorf("failed to parse Pokémon details from %s", url)
	}

	// Clean up height and weight
	height := strings.TrimSpace(centerPanelInfo.Eq(6).Text())
	height = strings.ReplaceAll(height, "\n\t\t\t", " ")

	weight := strings.TrimSpace(centerPanelInfo.Eq(7).Text())
	weight = strings.ReplaceAll(weight, "\n\t\t\t", " ")

	baseStats := doc.Find("td:contains('Base Stats - Total')").NextAll()
	hp, _ := strconv.Atoi(baseStats.Eq(0).Text())
	attack, _ := strconv.Atoi(baseStats.Eq(1).Text())
	defense, _ := strconv.Atoi(baseStats.Eq(2).Text())
	special, _ := strconv.Atoi(baseStats.Eq(3).Text())
	speed, _ := strconv.Atoi(baseStats.Eq(4).Text())

	return Pokemon{
		Name:           centerPanelInfo.Eq(1).Text(),
		Number:         fmt.Sprintf("#%03d", pokeID),
		Classification: centerPanelInfo.Eq(5).Text(),
		Height:         []string{height},
		Weight:         []string{weight},
		HitPoints:      hp,
		Attack:         attack,
		Defense:        defense,
		Special:        special,
		Speed:          speed,
	}, nil
}

func displayFormatted(p Pokemon) {
	fmt.Printf("Name\t\t %s\n", p.Name)
	fmt.Printf("Number\t\t %s\n", p.Number)
	fmt.Printf("Classification\t %s\n", p.Classification)
	fmt.Printf("Height\t\t %s\n", strings.Join(p.Height, " "))
	fmt.Printf("Weight\t\t %s\n", strings.Join(p.Weight, " "))
	fmt.Printf("HP\t\t %d\n", p.HitPoints)
	fmt.Printf("Attack\t\t %d\n", p.Attack)
	fmt.Printf("Defense\t\t %d\n", p.Defense)
	fmt.Printf("Special\t\t %d\n", p.Special)
	fmt.Printf("Speed\t\t %d\n", p.Speed)
	fmt.Println(strings.Repeat("-", 20))
}

func saveToJSON(data []Pokemon) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal data: %v", err)
	}

	err = ioutil.WriteFile(OUTPUT_FILE, file, 0644)
	if err != nil {
		log.Fatalf("failed to write data to file: %v", err)
	}
}

func validateInput(firstID, lastID int) (int, int) {
	if firstID >= 906 || lastID >= 906 {
		log.Fatalf("Error: This Pokémon is not yet supported!")
	}
	if lastID < firstID {
		lastID = firstID
	}
	return firstID, lastID
}
