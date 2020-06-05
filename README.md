# GoLang

### 1. GoLang
1. 켄톰슨(B 언어), 롭파이크(UTF8
2. C언어와 비슷, 파워플.
3. 성능이 좋고, 안정적이다.
4. 빠른 생산성
5. http://golang.org
6. go version
7. go env
8. go path 경로 확인하고, 폴더 만들어 줍니다.

### 2. hello go
1. package, 묶음, 모듈, 라이브러리
  - 라이브러리 : 기능들을 묶어 놓은 것
  - 모듈 : 기능들을 조금 더 특정해서 묶음
  - 프레임웍 : 기능묶음 + 절차(규정)
  - 엔진 : 기능묶음 + 다른 툴
  - package main
```go
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
```

### 변수(변하는 수)
1. 변수의 속성
  - 이름
  - 값
  - type : 정수(int), 실수(float), 문자열(string)