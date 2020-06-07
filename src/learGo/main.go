package main

import (
	"learGo/mydic"
	// "log"
	// "learGo/banking"
	"fmt"
)

func main(){
	// account := banking.NewAccount("gi");
	// account.Deposit(100)
	// err := account.Withdraw(1000)
	// if err != nil{
	// 	// log.Fatalln(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance())
	// fmt.Println(account.Balance(), account.Owner())
	// account.ChangeOwner("Geonil Jang");
	// fmt.Println(account.Owner())
	dictionary := mydic.Dictionary{"first":"First Words"}
	// fmt.Println(dictionary["first"])
	// definition ,err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(definition)
	// }
	word := "s"
	definition := "S is s...."
	addError := dictionary.Add(word, definition);
	if(addError != nil){
		fmt.Println(addError)
	}
	// definition2, err := dictionary.Search(word)
	// fmt.Println(err)
	// if err != nil{
	// 	fmt.Println(definition2)
	// }
	err := dictionary.Update(word, "Update Dic-----")
	if err != nil {
		fmt.Println(err)
	}
	base, _ := dictionary.Search(word)
	fmt.Println(base)
	
	
}