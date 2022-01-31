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

	ch := make(chan string)

	for _, link := range links {
		// create a new thread go routine...
		// ...and run this func with it
		go checkLink(link, ch)
	}

	// prints only one output
	//fmt.Println(<-ch)

	// prints all outputs
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-ch)
	//}

	// repeating routines (infinite loop)
	//for {
	//	go checkLink(<-ch, ch)
	//}

	// alternative loop syntax and sleeping a routine
	//for l := range ch {
	//	time.Sleep(5 * time.Second)
	//	go checkLink(l, ch)
	//}

	// function literal
	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link //"or any string"
		return
	}

	fmt.Println(link, "is up!")
	ch <- link //"or any string"
}
