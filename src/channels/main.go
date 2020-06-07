package main

import (
	"errors"
	"net/http"
	"fmt"
)

var errRequestFailed = errors.New("Request Failed")

type result struct{
	url string
	status string
}



func main(){
	c := make(chan result)
	results :=  map[string]string{}

	urls := []string{"https://www.naver.com/","https://google.com"}
	for _, url := range urls{
		go hitUrl(url, c)
	}

	for i:= 0 ; i < len(urls) ; i++ {
		res := <-c
		results[res.url] = res.status
	}

	for url, status := range results{
		fmt.Println(url, status)
	}

}

func hitUrl(url string, c chan result){
	resp, error := http.Get(url)
	if error != nil || resp.StatusCode >= 400{
		c <- result{url:url,status:"FAILED"}
	}else{
		c <- result{url:url, status:"OK"}
	}
	
}