package utils

import (
	"github.com/gocolly/colly"
)

func FetchMetadata(url string) (map[string]string, error) {
	c := colly.NewCollector()
	metadata := make(map[string]string)

	c.OnHTML("meta[property='og:title']", func(e *colly.HTMLElement) {
		metadata["title"] = e.Attr("content")
	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		if _, ok := metadata["title"]; !ok {
			metadata["title"] = e.Text
		}
	})
	c.OnHTML("meta[property='og:description']", func(e *colly.HTMLElement) {
		metadata["description"] = e.Attr("content")
	})
	c.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
		metadata["image"] = e.Attr("content")
	})
	c.OnHTML("meta[property='og:url']", func(e *colly.HTMLElement) {
		metadata["url"] = e.Attr("content")
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	if _, ok := metadata["url"]; !ok {
		metadata["url"] = url
	}

	return metadata, nil
}