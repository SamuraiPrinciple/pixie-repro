package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		(func() {
			resp, err := http.Get("https://example.com")
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()
			fmt.Println("Response status:", resp.Status)
		})()
		time.Sleep(10 * time.Second)
	}
}
