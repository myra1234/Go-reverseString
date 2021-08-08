package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("Program to reverse a string that user inputs!")

	fmt.Print("Enter string to reverse, hit enter when string is complete: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err!=nil{
		fmt.Println("An error occured while reading your input. Please try again", err)
		return
	}

	fmt.Println("You entered - ", input)

	fmt.Println("Reverse string is", reverseString(input))
}

func reverseString(input string)(result string){
	for _, v:= range input{
		result = string(v) +result
	}
	return
}