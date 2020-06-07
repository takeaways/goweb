package main
//이 파일은 package <name> 패키지 이다.
//프로그램 실행 파일 -> HDD -> 메모리 적제 -> IP 부터 시작
//그래서 Go 에서는 package main으로 실행 IP(Instructor Point)

import (
	"fmt"
)
//가져온다. package를 가져온다.
//기능의 묶음을 가져온다.
//표준패키지 입니다.

func main(){
	//main은 함수 입력 -> 연산 -> 출력
	fmt.Print("Hello Go")
	//fmt.(안에있는)Print()함수를 사용하겠다.
}