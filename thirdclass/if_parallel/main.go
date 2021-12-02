package main

import "fmt"

func main(){
	var point1 = [2]float64{1, 1}
	var point2 =[2]float64{2,2}
	var line1a float64
	var line1b float64
	var line1_b float64
	line1a = (point1[1]-line1_b)/point1[0]
	line1b = (point2[1]-line1_b)/point2[0]
	if line1a ==line1b{
		fmt.Println("a",line1a)
	}

	var point3 = [2]float64{1, 0}
	var point4 =[2]float64{2,1}
	var line2a float64
	var line2b float64
	var line2_b float64

	line2a = (point3[1]-line2_b)/point3[0]
	line2b = (point4[1]-line2_b)/point4[0]
	fmt.Println("b",line2a)
	fmt.Println("b",line2b)

	if line2a ==line2b{
		fmt.Println("b",line2a)
	}



}