package main

import "fmt"

func testFat(name string,sex string, bmi float64, age int){
	var sexWeight float64
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*sexWeight)/100
	fmt.Println("姓名：", name)
	fmt.Printf("体脂率是：%.2f\n",fatRate)

	if sex == "男" {
		if age >=18 && age <=39{
			if fatRate <=0.1{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.1&&fatRate<=0.16{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.16&&fatRate<=0.21 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.21&&fatRate<=0.26{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age >=40 &&age <=59{
			if fatRate <=0.11{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.11&&fatRate<=0.17{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.17&&fatRate<=0.22 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.22&&fatRate<=0.27{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age>=60{
			if fatRate <=0.13{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.13&&fatRate<=0.19{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.19&&fatRate<=0.24 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.24&&fatRate<=0.29{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else{
			fmt.Println("暂时不支持计算未成年人的体脂率～")
		}
	} else {
		if age >=18 && age <=39{
			if fatRate <=0.2{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.2&&fatRate<=0.27{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.27&&fatRate<=0.34 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.34&&fatRate<=0.39{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age >=40 &&age <=59{
			if fatRate <=0.21{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.21&&fatRate<=0.28{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.28&&fatRate<=0.35 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.35&&fatRate<=0.41{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age>=60{
			if fatRate <=0.22{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.22&&fatRate<=0.29{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.29&&fatRate<=0.36 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.36&&fatRate<=0.41{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else{
			fmt.Println("暂时不支持计算未成年人的体脂率～")
		}
	}
}
func main() {
	var (
		names [3] string
		name1 string
		name2 string
		name3 string
		weight [3] float64
		weight1 float64
		weight2 float64
		weight3 float64
		tall [3] float64
		tall1 float64
		tall2 float64
		tall3 float64
		age [3]int
		age1 int
		age2 int
		age3 int
		sex [3]string
		sex1 string
		sex2 string
		sex3 string

	)

	fmt.Print("请输入要计算的姓名（以空格分隔）：")
	fmt.Scanln(&name1, &name2,&name3)
	names[0] =name1
	names[1] =name2
	names[2] =name3
	fmt.Print("请输入要计算的体重（以空格分隔）：")
	fmt.Scanln(&weight1, &weight2,&weight3)
	weight[0] =weight1
	weight[1] =weight2
	weight[2] =weight3


	fmt.Print("请输入要计算的身高（以空格分隔）：")
	fmt.Scanln(&tall1, &tall2,&tall3)
	tall[0]=tall1
	tall[1]=tall2
	tall[2]=tall3

	fmt.Print("请输入要计算的年龄（以空格分隔）：")
	fmt.Scanln(&age1, &age2,&age3)
	age[0]=age1
	age[1]=age2
	age[2]=age3

	fmt.Print("请输入要计算的性别（以空格分隔）：")
	fmt.Scanln(&sex1, &sex2,&sex3)
	sex[0]=sex1
	sex[1]=sex2
	sex[2]=sex3

	for i:=0;i<3;i++{
		var bmi float64 = weight[i] / (tall[i] * tall[i])
		testFat(names[i],sex[i],bmi,age[i])
	}





}
