package main

import (
	"errors"
	"net/http"
	"fmt"
)

var errRequestFailed = errors.New("Request Failed")




func main(){
	results :=  map[string]string{}

	urls := []string{"https://www.naver.com/","https://google.com"}

	results["hello"] = "Hello"
	for _, url := range urls{
		result := "OK"
		err := hitUrl(url)
		if err != nil{
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results{
		fmt.Println(url, result)
	}

}

func hitUrl(url string) error{
	fmt.Println("Checking Url ", url)
	resp, error := http.Get(url)
	// fmt.Println(resp);
	// fmt.Println(error);
	if error != nil || resp.StatusCode >= 400{
		return errRequestFailed
	}
	return nil
}