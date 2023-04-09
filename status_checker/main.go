package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, site := range links {
		go checkSiteConnectivity(site, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkSiteConnectivity(link, c)
		}(l)
	}

}

func checkSiteConnectivity(siteLink string, c chan string) {
	resp, err := http.Get(siteLink)
	if err != nil {
		fmt.Printf("Site %s is down with err:%s and status:%d \n", siteLink, err.Error(), resp.StatusCode)
		c <- siteLink
		return
	}

	fmt.Printf("Site %s is up status:%d \n", siteLink, resp.StatusCode)
	c <- siteLink
}
