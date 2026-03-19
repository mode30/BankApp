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
	withdrawnAmount,err:=withdrawal("enter amount to withdraw",&totalAmount)
	if err !=nil{
		log.Fatal(err)
		// return 0,fmt.Errorf("cannot parse user input %w", err)
	}
	fmt.Print(withdrawnAmount)


	depositedAmount,err:=deposit("enter amount to deposit:")
	if err !=nil{
		log.Fatal(err)
		// return 0,fmt.Errorf("cannot parse user input %w", err)
	}
	fmt.Print(depositedAmount)

}


func withdrawal(prompt string,totalAmount *float64)(bankAccount *float64,err error){
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
	if userAmount > *totalAmount{
		state=true
		fmt.Println("amount over the total amount in accout")
	}else{
		state=false
	}

	if state{
		return  nil,errors.New("cannot withdraw amount")

	}else{
		*totalAmount-=userAmount
		return totalAmount,nil
	}

}
func deposit(prompt string)(bankAccount float64,err error){


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
	if userAmount < 0{
		state=true
		fmt.Println("amount must be greater than 0")
	}else{
		state=false
	}

	if state{
		return 0,fmt.Errorf("error parsing value less than 0:%w", err)
		// return  0,errors.New("cannot withdraw amount")

	}else{
		totalAmount+=userAmount
		// result:=totalAmount-userAmount
		return totalAmount,nil
	}

}
