package main

import (
	"fmt"
)

func main() {
  //Array
  var langs [3]string
	langs[0] = "java"
	langs[1] = "python"
	langs[2] = "go"
	fmt.Println(langs)
  //Slices
	moreLangs:=[]string{"c","c++","Fortran"}
  moreLangs=append(moreLangs,"Ruby")
  moreLangs=append(moreLangs,"JS")
	fmt.Println(moreLangs)
  //Looping on
  for i,element:=range(moreLangs){
    fmt.Print(i)
    fmt.Println(element)
  }

}
