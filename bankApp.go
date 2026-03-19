package main

import (
	"fmt"
	"log"
	"time"
	 "excercise/bankApp/shared"
)

func main(){


	var account Bank

		account.accountHolder="benjamin"
		account.balance=shared.TotalAmount
		account.createdAt=time.Now()

	// }
	withdrawnAmount,err:=withdrawal("enter amount to withdraw",&account)
	// withdrawnAmount,err:=withdrawal("enter amount to withdraw",&shared.TotalAmount)
	if err !=nil{
		log.Fatal(err)
		// return 0,fmt.Errorf("cannot parse user input %w", err)
	}
	fmt.Print(withdrawnAmount)


	depositedAmount,err:=deposit("enter amount to deposit:",&account)
	if err !=nil{
		log.Fatal(err)
		// return 0,fmt.Errorf("cannot parse user input %w", err)
	}
	fmt.Print(depositedAmount)

}


