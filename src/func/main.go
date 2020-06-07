package main

import (
	"strings"
	"fmt"
)

func multiply(a ,b int) int{
	return a *b
}
func lanAndUpper(name string)(int, string){
	return len(name), strings.ToUpper(name)
}
func repeatMe(words ...string) {
	fmt.Println(words)
}

//naked return
func notReturn(name string)(length int, uppercase string){
	length = len(name)
  uppercase = strings.ToUpper(name)
	return 
}

//defer 함수가 끝날 때 추가적으로 동작하게 한다.
func add(a int, b int) int{
	defer fmt.Println("I'm Done");

	return a + b
}
func main(){
	fmt.Println(multiply(2,3))
	totalLength, upperName := lanAndUpper("Geonil")
	fmt.Println(totalLength, upperName)
	repeatMe("gi","nico","cola","may")
	len, upName := notReturn("nick")
	fmt.Println(len, upName)
	fmt.Println(add(10,20))
}
