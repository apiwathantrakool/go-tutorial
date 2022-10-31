package main

import "fmt"

func main(){
	text := getText();
	fmt.Println(text)
}

func getText() string {
	return "Hello World";
}