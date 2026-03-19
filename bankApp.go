package main

import (
	// "bufio"
	"errors"
	"fmt"
	"log"
	// "os"
	"strconv"
	// "strings"
	"time"
	"excercise/bankApp/userInput"
)

type Bank struct{
	accountHolder string
	holderName string
	balance float64
	createdAt time.Time
	lastTransaction time.Time

}
var totalAmount float64=10000
func main(){
	result,err:=withdrawal("enter amount to withdraw")
	if err !=nil{
		log.Fatal(err)
		// return 0,fmt.Errorf("cannot parse user input %w", err)
	}

}


func withdrawal(prompt string)(bankAccount float64,err error){
	var state bool
	// state=false
	// var value float64
	fmt.Print(prompt)
	userInput,err:=userInput.User()
	if err!=nil{
		log.Fatal(err)
	}
	userAmount,err:=strconv.ParseFloat(userInput,64)
	if err !=nil{
		log.Fatal(err)
	}
	if userAmount > totalAmount{
		state=true
		fmt.Println("amount over the total amount in accout")
	}else{
		state=false
	}

	if state{
		return  0,errors.New("cannot withdraw amount")

	}else{
		result:=totalAmount-userAmount
		return result,nil
	}

}
// func deposit(prompt string,amount float64)(bankAccount string,err error){

// }
