package main

import(
	"fmt"
	"excercise/bankApp/userInput"
	// "excercise/bankApp/shared"
	"strconv"
	"errors"
	"time"
	// "log"
)


type Bank struct{
	accountHolder string
	// holderName string
	balance float64
	createdAt time.Time
	lastTransaction time.Time

}
func(totalAmount *Bank)withdrawal(prompt string)(bankAccount *Bank,err error){
// func withdrawal(prompt string,totalAmount *float64)(bankAccount *float64,err error){
	var state bool
	// state=false
	// var value float64
	fmt.Print(prompt)
	userInput,err:=userInput.User()
	if err!=nil{
		return nil,fmt.Errorf("cannot read user input %w",err)
		// log.Fatal(err)
	}
	userAmount,err:=strconv.ParseFloat(userInput,64)
	if err !=nil{
		return nil,fmt.Errorf("cannot get user input:%w", err)
		// log.Fatal(err)
	}
	if userAmount > totalAmount.balance{
		state=true
		// fmt.Println("amount over the total amount in accout")
	}else{
		state=false
	}

	if state{
		return  nil,errors.New("cannot withdraw amount")

	}else{
		totalAmount.balance-=userAmount
		return totalAmount,nil
	}

}
func (totalAmount *Bank)deposit(prompt string)(bankAccount *Bank,err error){


	var state bool
	// state=false
	// var value float64
	fmt.Print(prompt)
	userInput,err:=userInput.User()
	if err!=nil{
		return nil,fmt.Errorf("cannot read user input", err)
		// log.Fatal(err)
	}
	userAmount,err:=strconv.ParseFloat(userInput,64)
	if err !=nil{
		return nil,fmt.Errorf("cannot read user input", err)
		// log.Fatal(err)
	}
	if userAmount < 0{
		state=true
		// fmt.Println("amount must be greater than 0")
	}else{
		state=false
	}

	if state{
		return nil,errors.New("error parsing value less than 0")
		// return  0,errors.New("cannot withdraw amount")

	}else{
		totalAmount.balance+=userAmount
		// result:=totalAmount-userAmount
		return totalAmount,nil
	}

}


func (bankInfo *Bank)getBalance()(accountBalance *Bank,err error){
	personAccount:=Bank{
		accountHolder:bankInfo.accountHolder,
		balance:bankInfo.balance,
		createdAt:bankInfo.createdAt,
		lastTransaction:bankInfo.lastTransaction,
	}
	return &personAccount,nil

}




func (newBankAccount *Bank)newAccount()(bankAccount *Bank,err error){
	if newBankAccount.accountHolder=="" || newBankAccount.balance <= 0 {
		return nil,errors.New("first name,last name,balance greater than 0 has to be provided")

	}
	return &Bank{

		accountHolder:newBankAccount.accountHolder,
		balance :newBankAccount.balance,
		createdAt:newBankAccount.createdAt,
		lastTransaction:newBankAccount.lastTransaction,

	},nil
}
