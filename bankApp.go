package main

import (
	"fmt"
	"log"
	"time"
	 "excercise/bankApp/shared"
)

type Bank struct{
	accountHolder string
	holderName string
	balance float64
	createdAt time.Time
	lastTransaction time.Time

}
func main(){


	withdrawnAmount,err:=withdrawal("enter amount to withdraw",&shared.TotalAmount)
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


