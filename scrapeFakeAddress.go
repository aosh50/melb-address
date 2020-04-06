package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func scrape() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.fakeaddressgenerator.com/AU_Real_Random_Address/index/city/Melbourne")
}
