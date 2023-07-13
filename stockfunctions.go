package main

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

func checkAllStock() {
	// check for 8GB Pi 4
	for _, link := range pi8links {
		website := determineLink(link)
		switch website {
		case "www.adafruit.com":
			// check stock, if returns true, send post request to discord webhook
			if checkStockAdafruit(link) {
				// send post request to discord webhook
				sendDCmesage("Adafruit Pi 6", link)
			}

		case "vilros.com":
			if checkStockVilros(link) {
				// send post request to discord webhook
				sendDCmesage("Vilros Pi 6", link)
			}
		case "www.pishop.us":
			if checkStockPishop(link) {
				// send post request to discord webhook
				sendDCmesage("PiShop Pi 6", link)
			}
		case "www.sparkfun.com":
			if checkStockSparkfun(link) {
				// send post request to discord webhook
				sendDCmesage("Sparkfun Pi 6", link)
			}
		}
	}
	// check for 4GB Pi 4
	for _, link := range pi4links {
		website := determineLink(link)
		switch website {
		case "www.adafruit.com":
			if checkStockAdafruit(link) {
				// send post request to discord webhook
				sendDCmesage("Adafruit Pi 4", link)
			}
		case "vilros.com":
			if checkStockVilros(link) {
				// send post request to discord webhook
				sendDCmesage("Vilros Pi 4", link)
			}
		case "www.pishop.us":
			if checkStockPishop(link) {
				// send post request to discord webhook
				sendDCmesage("PiShop Pi 4", link)
			}
		case "www.sparkfun.com":
			if checkStockSparkfun(link) {
				// send post request to discord webhook
				sendDCmesage("Sparkfun Pi 4", link)
			}
		}
	}
}

func determineLink(link string) string {
	re := regexp.MustCompile(`(http(s)?:\/\/)|(\/.*){1}`)
	website := re.ReplaceAllString(link, "")
	return website
}

func checkStockAdafruit(link string) bool {
	var inStock bool = false
	// regex to get the final part of the url, minus the .html
	re := regexp.MustCompile(`[0-9]+`)
	sku := re.FindString(link)
	c := colly.NewCollector()
	c.OnHTML("ol.meta_pid_boxes", func(e *colly.HTMLElement) {
		e.ForEach("li[data-part-id]", func(_ int, el *colly.HTMLElement) {
			if el.Attr("data-part-id") == sku {
				if el.ChildText("span.meta_pid_box_status") != "Out of stock" {
					inStock = true
				}
			}
		})
	})
	c.Visit(link)
	if inStock {
		log.Println("Adafruit: ", sku, " is in stock")
	} else {
		log.Println("Adafruit: ", sku, " is out of stock")

	}
	return inStock
}

func checkStockVilros(link string) bool {
	var inStock bool = false
	re := regexp.MustCompile(`(?m)[^/]+$`)
	sku := re.FindString(link)
	c := colly.NewCollector()
	c.OnHTML("div.payment-buttons", func(e *colly.HTMLElement) {
		// if button text is not Sold Out, then in stock
		if e.ChildText("button") != "Sold Out" {
			inStock = true
		}
	})
	c.Visit(link)
	if inStock {
		log.Println("Vilros: ", sku, " is in stock")
	} else {
		log.Println("Vilros: ", sku, " is out of stock")
	}
	return inStock
}

func checkStockPishop(link string) bool {
	inStock := false
	// regex is everything after the last /
	re := regexp.MustCompile(`(?m)[^/]+$`)
	sku := re.FindString(link)
	c := colly.NewCollector()
	// on load of input id form-action-addToCart
	c.OnHTML("input#form-action-addToCart", func(e *colly.HTMLElement) {
		// if button text is not "Out of stock", then in stock
		if e.Attr("value") != "Out of stock" {
			inStock = true
		}
	})
	c.Visit(link)
	if inStock {
		log.Println("Pishop: ", sku, " is in stock")
	} else {
		log.Println("Pishop: ", sku, " is out of stock")
	}
	return inStock
}

func checkStockSparkfun(link string) bool {
	inStock := false
	re := regexp.MustCompile(`(?m)[^/]+$`)
	sku := re.FindString(link)
	c := colly.NewCollector()
	c.OnHTML("p.add-buttons", func(e *colly.HTMLElement) {
		// if the value of the input is "Add to Cart" then in stock
		if e.ChildAttr("input", "value") == "Add to Cart" {
			inStock = true
		}
	})
	c.Visit(link)
	if inStock {
		log.Println("Sparkfun: ", sku, " is in stock")
	} else {
		log.Println("Sparkfun: ", sku, " is out of stock")
	}
	return inStock
}
