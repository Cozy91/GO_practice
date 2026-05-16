package main
import (
	"fmt"
	"errors"
)
func main(){
var	number1,number2,input,result int
var err error
fmt.Printf("\n enter two numbers: ")
fmt.Scan(&number1,&number2)

fmt.Println("\nWelcome to calculator!")
fmt.Println("\nPress 1 for addition")
fmt.Println("\nPress 2 for subtraction")
fmt.Println("\nPress 3 for multi")
fmt.Println("\nPress 4 for division")
fmt.Scan(&input)
switch input{
case 1:
	result =add(number1,number2)
case 2:
	result,err =sub(number1,number2)
  if err !=nil{
		fmt.Println("Error: ",err)
	}
case 3:
	result =mult(number1,number2)
case 4:
	result,err =divide(number1,number2)
	if err!=nil{
		fmt.Println("Error: ",err)
		return
	}
}
fmt.Printf("\n Your result is : %v",result)

}

func add(number1 int,number2 int) int{
	return number1+number2
}
func sub(number1 int ,number2 int ) (int,error) {
	if number2>number1{
    return 0, errors.New("Invalid subtraction")
	}
	return (number1-number2),nil
}
func mult(number1 int,number2 int )int{
	return number1*number2
}
func divide(number1 int,number2 int)(int,error){
	if number2 == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return number1/number2,nil
}
