package main

import (
	"log"
	"net/http"
	"strings"
)

func sendDCmesage(site string, link string) {
	// send post request to discord webhook.
	// the post should have a body of {"msg": "This is my message"}

	// create message in the form of {"msg": "RaspberryPi Tracker🥧\n\Site is in stock!\nLink: link"}
	msg := "{\"msg\": \"RaspberryPi Tracker🥧\\n\\n" + site + " is in stock!\\n" + link + "\"}"
	// create the post request
	req, err := http.NewRequest("POST", "LINK", strings.NewReader(msg))
	if err != nil {
		log.Fatal(err)
	}
	// set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// send the post request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}
