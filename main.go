package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func main() {
	fName := "data.csv"

	file, err := os.Create(fName)

	if err != nil {
		log.Fatalf("Could not create file, err : %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	c.OnHTML(".internship_meta", func(h *colly.HTMLElement) {
		writer.Write([]string{
			h.ChildText("a"),
			h.ChildText("span"),
		})
	})

	for i := 0; i < 190; i++ {
		fmt.Printf("Scraping page : %d\n", i)

		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping Completed\n")
	log.Println(c)

}
