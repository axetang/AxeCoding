/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.NormFloat64
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NormFloat64() float64
 ------------------------------------------------------------------------------------
 **Description:
 NormFloat64 returns a normally distributed float64 in the range [-math.MaxFloat64,
 +math.MaxFloat64] with standard normal distribution (mean = 0, stddev = 1) from the
 default Source. To produce a different normal distribution, callers can adjust the
 output using:
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NormFloat64函数返回一个区间为[-math.MaxFloat64, +math.MaxFloat64]的标准正态分布,
 2. 标准正太分布的参数(位置参数期望mean = 0, 尺度参数标准差stddev = 1).
 3. 如果想改变不同的位置参数和尺度参数,可以通过下面方法变通:
	sample = NormFloat64() * desiredStdDev + desiredMean
*************************************************************************************/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 10000
	i := 0
	sum := 0.0
	for i < n {
		x := rand.NormFloat64()
		i += 1
		sum += x
	}
	expect := sum / float64(n)
	fmt.Println(expect)
}
