package main

import "fmt"
func main(){
	gi := map[string]string{"name":"gi","age":"29"}
	fmt.Println(gi["name"])
	for key, value := range gi{
		fmt.Println(key, value)
	}
}