package main

import "fmt"

func testFat(name string, sex string, bmi float64, age int) (fatRate float64) {

	if sex == "男" {
		fatRate := getFatRate(bmi, age, getmanFatRate)
		manSuggestion(age, fatRate)
		return fprintfateRate(name, fatRate)
	} else {
		fatRate := getFatRate(bmi, age, getwomanFatRate)
		womanSuggestion(age, fatRate)
		return fprintfateRate(name, fatRate)
	}
}
func fprintfateRate(name string, fatRate float64) float64 {
	fmt.Println("姓名：", name)
	fmt.Printf("体脂率是：%.2f\n", fatRate)
	return fatRate
}
func getmanFatRate(bmi float64, age int)(fatRate float64){
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*1) / 100
	return
}
func getwomanFatRate(bmi float64, age int)(fatRate float64){
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*0) / 100
	return
}
type FatRate func(bmi float64, age int)(fatRate float64)
func getFatRate(bmi float64, age int, c FatRate)(fatRate float64){
	return c(bmi, age)
}
func womanSuggestion(age int, fatRate float64) {
	if age >= 18 && age <= 39 {
		if fatRate <= 0.2 {
			fmt.Println("目前偏瘦，多吃多锻炼")
		} else if fatRate > 0.2 && fatRate <= 0.27 {
			fmt.Println("目前标准，太棒了，要保持哦")
		} else if fatRate > 0.27 && fatRate <= 0.34 {
			fmt.Println("目前偏重，吃完饭多散散步，消化消化")
		} else if fatRate > 0.34 && fatRate <= 0.39 {
			fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
		} else {
			fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.21 {
			fmt.Println("目前偏瘦，多吃多锻炼")
		} else if fatRate > 0.21 && fatRate <= 0.28 {
			fmt.Println("目前标准，太棒了，要保持哦")
		} else if fatRate > 0.28 && fatRate <= 0.35 {
			fmt.Println("目前偏重，吃完饭多散散步，消化消化")
		} else if fatRate > 0.35 && fatRate <= 0.41 {
			fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
		} else {
			fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
		}
	} else if age >= 60 {
		if fatRate <= 0.22 {
			fmt.Println("目前偏瘦，多吃多锻炼")
		} else if fatRate > 0.22 && fatRate <= 0.29 {
			fmt.Println("目前标准，太棒了，要保持哦")
		} else if fatRate > 0.29 && fatRate <= 0.36 {
			fmt.Println("目前偏重，吃完饭多散散步，消化消化")
		} else if fatRate > 0.36 && fatRate <= 0.41 {
			fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
		} else {
			fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
		}
	} else {
		fmt.Println("暂时不支持计算未成年人的体脂率～")
	}
}
func manSuggestion(age int, fatRate float64) {
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("目前偏瘦，多吃多锻炼")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("目前标准，太棒了，要保持哦")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("目前偏重，吃完饭多散散步，消化消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
		} else {
			fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.11 {
			fmt.Println("目前偏瘦，多吃多锻炼")
		} else if fatRate > 0.11 && fatRate <= 0.17 {
			fmt.Println("目前标准，太棒了，要保持哦")
		} else if fatRate > 0.17 && fatRate <= 0.22 {
			fmt.Println("目前偏重，吃完饭多散散步，消化消化")
		} else if fatRate > 0.22 && fatRate <= 0.27 {
			fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
		} else {
			fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
		}
	} else if age >= 60 {
		if fatRate <= 0.13 {
			fmt.Println("目前偏瘦，多吃多锻炼")
		} else if fatRate > 0.13 && fatRate <= 0.19 {
			fmt.Println("目前标准，太棒了，要保持哦")
		} else if fatRate > 0.19 && fatRate <= 0.24 {
			fmt.Println("目前偏重，吃完饭多散散步，消化消化")
		} else if fatRate > 0.24 && fatRate <= 0.29 {
			fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
		} else {
			fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
		}
	} else {
		fmt.Println("暂时不支持计算未成年人的体脂率～")
	}
}
func inputAndCalcBmi(count int) []float64 {
	//定义存储容器
	var fatRates []float64
	for i := 0; i < count; i++ {
		var (
			name1   string
			weight1 float64
			tall1   float64
			age1    int
			sex1    string
		)
		//终端获取需要的参数
		fmt.Printf("请输入第%d个要计算的姓名：", i+1)
		fmt.Scanln(&name1)
		fmt.Printf("请输入第%d个要计算的体重(kg)：", i+1)
		fmt.Scanln(&weight1)
		fmt.Printf("请输入第%d个要计算的身高(m)：", i+1)
		fmt.Scanln(&tall1)
		fmt.Printf("请输入第%d个要计算的年龄：", i+1)
		fmt.Scanln(&age1)
		fmt.Printf("请输入第%d个要计算的性别(男/女)：", i+1)
		fmt.Scanln(&sex1)
		//计算bmi值
		var bmi float64 = weight1 / (tall1 * tall1)
		//计算出体脂率，并存储到slice
		fatRate := testFat(name1, sex1, bmi, age1)
		fatRates = append(fatRates, fatRate)
	}
	return fatRates
}
func calcAvgFatRate(fatRates []float64, sum_fat float64, count int) {
	for _, v := range fatRates {
		sum_fat += v
	}
	//计算平均体脂率，输出
	avg_fatRate := sum_fat / float64(count)
	fmt.Printf("本次共计算%d个人，平均体脂率是%.2f\n", count, avg_fatRate)
}
func main() {
	//获取要计算的人数
	var count int
	fmt.Print("请输入要计算的人数：")
	fmt.Scanln(&count)
	//终端持续录入
	for {
		if count != 0 {
			fatRates := inputAndCalcBmi(count)
			var sum_fat float64
			//遍历体脂率求和
			calcAvgFatRate(fatRates, sum_fat, count)
			return
		} else {
			fmt.Print("欢迎下次使用体脂计算器～")
			return
		}
	}

}



