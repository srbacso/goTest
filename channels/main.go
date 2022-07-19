package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://builtstream.com",
		"http://waterlooanalytics.com",
		"http://abtglobalexchange.com",
		"http://facebook.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Might be down")
		c <- "Site might be down"
		return
	}
	fmt.Println(link, " is UP")
	c <- link

}
