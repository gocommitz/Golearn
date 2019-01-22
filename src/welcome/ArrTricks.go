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
  //Looping on, '_' denotes a value which will not be referenced again to avoid unused variables
  for _,element:=range(moreLangs){
    fmt.Println(element)
  }

}
