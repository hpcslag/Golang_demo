package main

import "fmt"

var str string = "STRING VARIABLE"

var (
	Int int = 123
	Bool bool = true
	mStr,mInt,mBool = "123",123,false
)

func main(){
	AdvancedStr := "AdvancedStr!!!!!!!!!"
	fmt.Printf("AdvancedStr must be used! %s",AdvancedStr)
	fmt.Printf("string: %s",str);
}