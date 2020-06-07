package main

import "fmt"

func main(){
	a :=2
	b :=&a
	*b = 20 // 주소의 값을 업데이트 하기
	fmt.Println(a) //&a 주소야 *안을 들여다 볼게

}