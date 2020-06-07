package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}
func main(){
	favFood := []string{"c","a"}
	gi := person{name:"gi", age:29, favFood:favFood}

	fmt.Println(gi.name)
}