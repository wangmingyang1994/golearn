package main

import "fmt"

func add(a, b int)int{
	return a+b
}
func sub(a, b int)int{
	return a-b
}

type calc func(a,b int) int

func computer(a,b int,c  calc)int{
	return c(a,b)
}
func main() {
	result:=computer(3,2,add)
	fmt.Println(result)
}
