package userInput


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func User()( string,error){
	buffer:=bufio.NewReader(os.Stdin)
	bufferString,err:=buffer.ReadString('\n')
	if err !=nil{
		return "",fmt.Errorf("%w", err)
		// log.Fatal(err)
	}
	userInputTrimmed:=strings.TrimSpace(bufferString)
	 return userInputTrimmed,nil
}
