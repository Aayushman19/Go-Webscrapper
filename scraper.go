package main

import (
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

func main(){
	args := os.Args
	url := args[1]
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error occured ", err)
	})

	collector.OnHTML("h1", func(h *colly.HTMLElement) {
		fmt.Println(h.DOM.Html())
	})

	collector.Visit(url)
}