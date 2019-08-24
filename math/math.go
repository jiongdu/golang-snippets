package math

import (
	"fmt"
	"math"
)

func mathOperator() {
	i := 1
	fmt.Println(math.Abs(float64(i)))		//绝对值
	fmt.Println(math.Ceil(float64(i)))		//向上取整
	fmt.Println(math.Floor(float64(i)))		//向下取整
	fmt.Println(math.Mod(11, 3))		//取余数
	fmt.Println(math.Pow(3, 2))		//x的y次方
	fmt.Println(math.Sqrt(8))			//开平方
	fmt.Println(math.Cbrt(8))			//开立方
}