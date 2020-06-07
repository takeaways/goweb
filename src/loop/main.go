package main

import (
	"fmt"
)

func superAdd(numbers ...int) int{
	//forEach 같이 사용할 수 있다
	total :=0;
	//ignore 시키기
	for _, number := range numbers{
		total += number
	}
	for i :=0 ; i < len(numbers); i++ {
		// fmt.Println(numbers[i])
	}
	return total
}

func canIDrink(age int) bool {

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
		return true
	
	// switch  {
	// 	case age < 18:
	// 		return false
	// 	case age == 18:
	// 		return true
		
	// }


}

func main(){
	total := superAdd(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("result",total)
	fmt.Println(canIDrink(14))
}