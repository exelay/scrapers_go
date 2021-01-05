package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type CarAd struct {
	Title string
	Price string
}

func main() {
	results := []CarAd{}
	c := colly.NewCollector()

	c.OnHTML("a[data-ftid=\"bulls-list_bull\"]", func(e *colly.HTMLElement) {
		temp := CarAd{}
		temp.Title = e.ChildText("span[data-ftid=\"bull_title\"]")
		temp.Price = e.ChildText("span[data-ftid=\"bull_price\"]")
		results = append(results, temp)
	})

	c.Visit("https://spb.drom.ru/auto/")
	fmt.Print(results)
}
