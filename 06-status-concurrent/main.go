package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, url := range urls {
		go checkUrl(url, channel)
	}

	for {
		go func() {
			time.Sleep(time.Second * 5)
			checkUrl(<-channel, channel)
		}()
	}
}

func checkUrl(url string, channel chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "FAIL")
		channel <- url
		return
	}

	fmt.Println(url, "OK")
	channel <- url
}
