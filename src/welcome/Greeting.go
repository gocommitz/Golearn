package main

import (
	"fmt"
	"time"
  "errors"
  "os"
)

func main() {
	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)
	fmt.Println(greeting)
  if(err != nil){
    fmt.Println(err)
    os.Exit(1)
  }
}

func getGreeting(hour int) (string,error) {
 var message string
  if hour < 12 {

    if hour < 6 {
      	error:=errors.New("Too early")
        return message,error
    }else {
      	message="Good Morning"
    }
	} else {
		message="Good Night"
	}
 return message,nil
}
