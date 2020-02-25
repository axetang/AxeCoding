/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.ExpFloat64
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func (r *Rand) ExpFloat64() float64
 ------------------------------------------------------------------------------------
 **Description:
 ExpFloat64 returns an exponentially distributed float64 in the range
 (0, +math.MaxFloat64] with an exponential distribution whose rate parameter (lambda)
 is 1 and whose mean is 1/lambda (1). To produce a distribution with a different rate
 parameter, callers can adjust the output using:
		sample = ExpFloat64() / desiredRateParameter
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 这个函数主要实现了指数分布公式;
 2. 函数定义域在(0, +∞);在此用math.MaxFloat64表示正无穷, 率参数λ归元为1,函数的期望是1/λ(λ=1),
 返回值计算如下:
	f(x,1)
 3. 如果想要λ是别的值怎么办呢?可以通过调整分母来达到目的.
	sample = ExpFloat64() / desiredRateParameter
*************************************************************************************/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := 0.0
	n := 10
	i := 0
	for i < n {
		x := rand.ExpFloat64()
		sum += x
		i += 1
		fmt.Printf("i=%d,x=rand.ExpFloat64()=%f,sum=%f\n", i, x, sum)
	}
	expect := sum / (float64)(n)
	fmt.Println(expect)
}
