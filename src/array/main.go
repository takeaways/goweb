package main

import "fmt"

func main(){
	// withouth length [] 제한이 없다
	names := []string{"gi","am","cc"}
	names = append(names, "GeonilJang")
	names[3] = "H"
	// names[4] = "O"
	fmt.Println(names)
}