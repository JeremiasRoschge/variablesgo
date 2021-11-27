package main

import (
	"fmt"
	"strconv"
)

func main()  {
	edad := 15

	edad_str := strconv.Itoa(edad)
	fmt.Println("I have "+ edad_str + "y/o")

	num1 := "10"

	num1_int,_ := strconv.Atoi(num1)

	fmt.Println(num1_int + 10)

}