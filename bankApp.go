package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

}


func user()( string,error){
	buffer:=bufio.NewReader(os.Stdin)
	bufferString,err:=buffer.ReadString('\n')
	if err !=nil{
		return "",fmt.Errorf("%w", err)
		// log.Fatal(err)
	}
	userInputTrimmed:=strings.TrimSpace(bufferString)
	 return userInputTrimmed,nil
}
func withdrawal(prompt string)(bankAccount float64,err error){

	var state bool
	// state=false
	// var value float64
	fmt.Print(prompt)
	userInput,err:=user()
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
		result:=totalAmount -userInput
		return result,nil
	}

}
func deposit(prompt string,amount float64)(bankAccount string,err error){

}
