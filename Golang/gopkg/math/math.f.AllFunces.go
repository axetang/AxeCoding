/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: math
 **Element: math.AllFunces
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Abs(x float64) float64
 Abs returns the absolute value of x.
 Special cases are:
 Abs(±Inf) = +Inf
 Abs(NaN) = NaN
 func Acos(x float64) float64
 Acos returns the arccosine, in radians, of x.
 Special case is:
 Acos(x) = NaN if x < -1 or x > 1
 func Acosh(x float64) float64
 Acosh returns the inverse hyperbolic cosine of x.
 Special cases are:
 Acosh(+Inf) = +Inf
 Acosh(x) = NaN if x < 1
 Acosh(NaN) = NaN
 func Asin(x float64) float64
 Asin returns the arcsine, in radians, of x.
 Special cases are:
 Asin(±0) = ±0
 Asin(x) = NaN if x < -1 or x > 1
 func Asinh(x float64) float64
 Asinh returns the inverse hyperbolic sine of x.
 Special cases are:
 Asinh(±0) = ±0
 Asinh(±Inf) = ±Inf
 Asinh(NaN) = NaN
 func Atan(x float64) float64
 Atan returns the arctangent, in radians, of x.
 Special cases are:
 Atan(±0) = ±0
 Atan(±Inf) = ±Pi/2
 func Atan2(y, x float64) float64
 Atan2 returns the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value.
 Special cases are (in order):
 Atan2(y, NaN) = NaN
 Atan2(NaN, x) = NaN
 Atan2(+0, x>=0) = +0
 Atan2(-0, x>=0) = -0
 Atan2(+0, x<=-0) = +Pi
 Atan2(-0, x<=-0) = -Pi
 Atan2(y>0, 0) = +Pi/2
 Atan2(y<0, 0) = -Pi/2
 Atan2(+Inf, +Inf) = +Pi/4
 Atan2(-Inf, +Inf) = -Pi/4
 Atan2(+Inf, -Inf) = 3Pi/4
 Atan2(-Inf, -Inf) = -3Pi/4
 Atan2(y, +Inf) = 0
 Atan2(y>0, -Inf) = +Pi
 Atan2(y<0, -Inf) = -Pi
 Atan2(+Inf, x) = +Pi/2
 Atan2(-Inf, x) = -Pi/2
 func Atanh(x float64) float64
 Atanh returns the inverse hyperbolic tangent of x.
 Special cases are:
 Atanh(1) = +Inf
 Atanh(±0) = ±0
 Atanh(-1) = -Inf
 Atanh(x) = NaN if x < -1 or x > 1
 Atanh(NaN) = NaN
 func Cbrt(x float64) float64
 Cbrt returns the cube root of x.
 Special cases are:
 Cbrt(±0) = ±0
 Cbrt(±Inf) = ±Inf
 Cbrt(NaN) = NaN

 func Ceil(x float64) float64
 Ceil returns the least integer value greater than or equal to x.
 Special cases are:
 Ceil(±0) = ±0
 Ceil(±Inf) = ±Inf
 Ceil(NaN) = NaN

 func Copysign(x, y float64) float64
 Copysign returns a value with the magnitude of x and the sign of y.

 func Cos(x float64) float64
 Cos returns the cosine of the radian argument x.
 Special cases are:
 Cos(±Inf) = NaN
 Cos(NaN) = NaN

 func Cosh(x float64) float64
 Cosh returns the hyperbolic cosine of x.
 Special cases are:
 Cosh(±0) = 1
 Cosh(±Inf) = +Inf
 Cosh(NaN) = NaN

 func Dim(x, y float64) float64
 Dim returns the maximum of x-y or 0.
 Special cases are:
 Dim(+Inf, +Inf) = NaN
 Dim(-Inf, -Inf) = NaN
 Dim(x, NaN) = Dim(NaN, x) = NaN

 func Erf(x float64) float64
 Erf returns the error function of x.
 Special cases are:
 Erf(+Inf) = 1
 Erf(-Inf) = -1
 Erf(NaN) = NaN

 func Erfc(x float64) float64
 Erfc returns the complementary error function of x.
 Special cases are:
 Erfc(+Inf) = 0
 Erfc(-Inf) = 2
 Erfc(NaN) = NaN

 func Erfcinv(x float64) float64
 Erfcinv returns the inverse of Erfc(x).
 Special cases are:
 Erfcinv(0) = +Inf
 Erfcinv(2) = -Inf
 Erfcinv(x) = NaN if x < 0 or x > 2
 Erfcinv(NaN) = NaN

 func Erfinv(x float64) float64
 Erfinv returns the inverse error function of x.
 Special cases are:
 Erfinv(1) = +Inf
 Erfinv(-1) = -Inf
 Erfinv(x) = NaN if x < -1 or x > 1
 Erfinv(NaN) = NaN

 func Exp(x float64) float64
 Exp returns e**x, the base-e exponential of x.
 Special cases are:
 Exp(+Inf) = +Inf
 Exp(NaN) = NaN
 Very large values overflow to 0 or +Inf. Very small values underflow to 1.

 func Exp2(x float64) float64
 Exp2 returns 2**x, the base-2 exponential of x.
 Special cases are the same as Exp.

 func Expm1(x float64) float64
 Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.
 Special cases are:
 Expm1(+Inf) = +Inf
 Expm1(-Inf) = -1
 Expm1(NaN) = NaN
 Very large values overflow to -1 or +Inf.

 func Float32bits(f float32) uint32
 Float32bits returns the IEEE 754 binary representation of f.
 func Float32frombits(b uint32) float32
 Float32frombits returns the floating point number corresponding to the IEEE 754 binary representation b.

 func Float64bits(f float64) uint64
 Float64bits returns the IEEE 754 binary representation of f.
 func Float64frombits(b uint64) float64
 Float64frombits returns the floating point number corresponding the IEEE 754 binary representation b.

 func Floor(x float64) float64
 Floor returns the greatest integer value less than or equal to x.
 Special cases are:
 Floor(±0) = ±0
 Floor(±Inf) = ±Inf
 Floor(NaN) = NaN

 func Frexp(f float64) (frac float64, exp int)
 Frexp breaks f into a normalized fraction and an integral power of two. It returns frac and exp satisfying f == frac × 2**exp, with the absolute value of frac in the interval [½, 1).
 Special cases are:
 Frexp(±0) = ±0, 0
 Frexp(±Inf) = ±Inf, 0
 Frexp(NaN) = NaN, 0

 func Gamma(x float64) float64
 Gamma returns the Gamma function of x.
 Special cases are:
 Gamma(+Inf) = +Inf
 Gamma(+0) = +Inf
 Gamma(-0) = -Inf
 Gamma(x) = NaN for integer x < 0
 Gamma(-Inf) = NaN
 Gamma(NaN) = NaN

 func Hypot(p, q float64) float64
 Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
 Special cases are:
 Hypot(±Inf, q) = +Inf
 Hypot(p, ±Inf) = +Inf
 Hypot(NaN, q) = NaN
 Hypot(p, NaN) = NaN

 func Ilogb(x float64) int
 Ilogb returns the binary exponent of x as an integer.
 Special cases are:
 Ilogb(±Inf) = MaxInt32
 Ilogb(0) = MinInt32
 Ilogb(NaN) = MaxInt32

 func Inf(sign int) float64
 Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
 func IsInf(f float64, sign int) bool
 IsInf reports whether f is an infinity, according to sign. If sign > 0, IsInf reports whether f is positive infinity. If sign < 0, IsInf reports whether f is negative infinity. If sign == 0, IsInf reports whether f is either infinity.

 func IsNaN(f float64) (is bool)
 IsNaN reports whether f is an IEEE 754 “not-a-number” value.

 func J0(x float64) float64
 J0 returns the order-zero Bessel function of the first kind.
 Special cases are:
 J0(±Inf) = 0
 J0(0) = 1
 J0(NaN) = NaN

 func J1(x float64) float64
 J1 returns the order-one Bessel function of the first kind.
 Special cases are:
 J1(±Inf) = 0
 J1(NaN) = NaN

 func Jn(n int, x float64) float64
 Jn returns the order-n Bessel function of the first kind.
 Special cases are:
 Jn(n, ±Inf) = 0
 Jn(n, NaN) = NaN

 func Ldexp(frac float64, exp int) float64
 Ldexp is the inverse of Frexp. It returns frac × 2**exp.
 Special cases are:
 Ldexp(±0, exp) = ±0
 Ldexp(±Inf, exp) = ±Inf
 Ldexp(NaN, exp) = NaN

 func Lgamma(x float64) (lgamma float64, sign int)
 Lgamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).
 Special cases are:
 Lgamma(+Inf) = +Inf
 Lgamma(0) = +Inf
 Lgamma(-integer) = +Inf
 Lgamma(-Inf) = -Inf
 Lgamma(NaN) = NaN

 func Log(x float64) float64
 Log returns the natural logarithm of x.
 Special cases are:
 Log(+Inf) = +Inf
 Log(0) = -Inf
 Log(x < 0) = NaN
 Log(NaN) = NaN

 func Log10(x float64) float64
 Log10 returns the decimal logarithm of x. The special cases are the same as for Log.
 func Log1p(x float64) float64
 Log1p returns the natural logarithm of 1 plus its argument x. It is more accurate than Log(1 + x) when x is near zero.
 Special cases are:
 Log1p(+Inf) = +Inf
 Log1p(±0) = ±0
 Log1p(-1) = -Inf
 Log1p(x < -1) = NaN
 Log1p(NaN) = NaN

 func Log2(x float64) float64
 Log2 returns the binary logarithm of x. The special cases are the same as for Log.
 func Logb(x float64) float64
 Logb returns the binary exponent of x.
 Special cases are:
 Logb(±Inf) = +Inf
 Logb(0) = -Inf
 Logb(NaN) = NaN

 func Max(x, y float64) float64
 Max returns the larger of x or y.
 Special cases are:
 Max(x, +Inf) = Max(+Inf, x) = +Inf
 Max(x, NaN) = Max(NaN, x) = NaN
 Max(+0, ±0) = Max(±0, +0) = +0
 Max(-0, -0) = -0

 func Min(x, y float64) float64
 Min returns the smaller of x or y.
 Special cases are:
 Min(x, -Inf) = Min(-Inf, x) = -Inf
 Min(x, NaN) = Min(NaN, x) = NaN
 Min(-0, ±0) = Min(±0, -0) = -0

 func Mod(x, y float64) float64
 Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.
 Special cases are:
 Mod(±Inf, y) = NaN
 Mod(NaN, y) = NaN
 Mod(x, 0) = NaN
 Mod(x, ±Inf) = x
 Mod(x, NaN) = NaN

 func Modf(f float64) (int float64, frac float64)
 Modf returns integer and fractional floating-point numbers that sum to f. Both values have the same sign as f.
 Special cases are:
 Modf(±Inf) = ±Inf, NaN
 Modf(NaN) = NaN, NaN

 func NaN() float64
 NaN returns an IEEE 754 “not-a-number” value.

 func Nextafter(x, y float64) (r float64)
 Nextafter returns the next representable float64 value after x towards y.
 Special cases are:
 Nextafter(x, x)   = x
 Nextafter(NaN, y) = NaN
 Nextafter(x, NaN) = NaN

 func Nextafter32(x, y float32) (r float32)
 Nextafter32 returns the next representable float32 value after x towards y.
 Special cases are:
 Nextafter32(x, x)   = x
 Nextafter32(NaN, y) = NaN
 Nextafter32(x, NaN) = NaN

 func Pow(x, y float64) float64
 Pow returns x**y, the base-x exponential of y.
 Special cases are (in order):
 Pow(x, ±0) = 1 for any x
 Pow(1, y) = 1 for any y
 Pow(x, 1) = x for any x
 Pow(NaN, y) = NaN
 Pow(x, NaN) = NaN
 Pow(±0, y) = ±Inf for y an odd integer < 0
 Pow(±0, -Inf) = +Inf
 Pow(±0, +Inf) = +0
 Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
 Pow(±0, y) = ±0 for y an odd integer > 0
 Pow(±0, y) = +0 for finite y > 0 and not an odd integer
 Pow(-1, ±Inf) = 1
 Pow(x, +Inf) = +Inf for |x| > 1
 Pow(x, -Inf) = +0 for |x| > 1
 Pow(x, +Inf) = +0 for |x| < 1
 Pow(x, -Inf) = +Inf for |x| < 1
 Pow(+Inf, y) = +Inf for y > 0
 Pow(+Inf, y) = +0 for y < 0
 Pow(-Inf, y) = Pow(-0, -y)
 Pow(x, y) = NaN for finite x < 0 and finite non-integer y

 func Pow10(n int) float64
 Pow10 returns 10**n, the base-10 exponential of n.
 Special cases are:
 Pow10(n) =    0 for n < -323
 Pow10(n) = +Inf for n > 308

 func Remainder(x, y float64) float64
 Remainder returns the IEEE 754 floating-point remainder of x/y.
 Special cases are:
 Remainder(±Inf, y) = NaN
 Remainder(NaN, y) = NaN
 Remainder(x, 0) = NaN
 Remainder(x, ±Inf) = x
 Remainder(x, NaN) = NaN

 func Round(x float64) float64
 Round returns the nearest integer, rounding half away from zero.
 Special cases are:
 Round(±0) = ±0
 Round(±Inf) = ±Inf
 Round(NaN) = NaN

 func RoundToEven(x float64) float64
 RoundToEven returns the nearest integer, rounding ties to even.
 Special cases are:
 RoundToEven(±0) = ±0
 RoundToEven(±Inf) = ±Inf
 RoundToEven(NaN) = NaN

 func Signbit(x float64) bool
 Signbit returns true if x is negative or negative zero.
 func Sin(x float64) float64
 Sin returns the sine of the radian argument x.
 Special cases are:
 Sin(±0) = ±0
 Sin(±Inf) = NaN
 Sin(NaN) = NaN

 func Sincos(x float64) (sin, cos float64)
 Sincos returns Sin(x), Cos(x).
 Special cases are:
 Sincos(±0) = ±0, 1
 Sincos(±Inf) = NaN, NaN
 Sincos(NaN) = NaN, NaN

 func Sinh(x float64) float64
 Sinh returns the hyperbolic sine of x.
 Special cases are:
 Sinh(±0) = ±0
 Sinh(±Inf) = ±Inf
 Sinh(NaN) = NaN

 func Sqrt(x float64) float64
 Sqrt returns the square root of x.
 Special cases are:
 Sqrt(+Inf) = +Inf
 Sqrt(±0) = ±0
 Sqrt(x < 0) = NaN
 Sqrt(NaN) = NaN

 func Tan(x float64) float64
 Tan returns the tangent of the radian argument x.
 Special cases are:
 Tan(±0) = ±0
 Tan(±Inf) = NaN
 Tan(NaN) = NaN

 func Tanh(x float64) float64
 Tanh returns the hyperbolic tangent of x.
 Special cases are:
 Tanh(±0) = ±0
 Tanh(±Inf) = ±1
 Tanh(NaN) = NaN

 func Trunc(x float64) float64
 Trunc returns the integer value of x.
 Special cases are:
 Trunc(±0) = ±0
 Trunc(±Inf) = ±Inf
 Trunc(NaN) = NaN

 func Y0(x float64) float64
 Y0 returns the order-zero Bessel function of the second kind.
 Special cases are:
 Y0(+Inf) = 0
 Y0(0) = -Inf
 Y0(x < 0) = NaN
 Y0(NaN) = NaN

 func Y1(x float64) float64
 Y1 returns the order-one Bessel function of the second kind.
 Special cases are:
 Y1(+Inf) = 0
 Y1(0) = -Inf
 Y1(x < 0) = NaN
 Y1(NaN) = NaN

 func Yn(n int, x float64) float64
 Yn returns the order-n Bessel function of the second kind.
 Special cases are:
 Yn(n, +Inf) = 0
 Yn(n ≥ 0, 0) = -Inf
 Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
 Yn(n, x < 0) = NaN
 Yn(n, NaN) = NaN
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 见程序中注释。
*************************************************************************************/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Abs(1-99.99999)", math.Abs(1-99.99999))
	fmt.Println("Abs(NaN)", math.Abs(math.NaN()))

	fmt.Printf("%.2f\n", math.Acos(1))
	fmt.Printf("%.2f\n", math.Acos(100))

	fmt.Printf("%.2f", math.Acosh(1))
	fmt.Printf("%.2f", math.Asin(0))
	fmt.Printf("%.2f", math.Asinh(0))

	fmt.Printf("%.2f", math.Atan(0))

	fmt.Printf("%.2f", math.Atan2(0, 0))
	fmt.Printf("%.2f", math.Atanh(0))

	//cube root立方根
	fmt.Println("Cbrt(5)=", math.Cbrt(5))

	//向上圆整的最小整数，结果为2.0
	fmt.Printf("math.Ceil(1.49)=%.1f\n", math.Ceil(1.49))

	//取第一个参数的值，第二个参数的符号
	fmt.Printf("math.Copysign(100.501，-25.02)=%f\n", math.Copysign(100.501, -25.02))

	fmt.Printf("math.Cos(math.Pi/2)=%.2f\n", math.Cos(math.Pi/2))

	fmt.Printf("math.Cosh(0))=%.2f\n", math.Cosh(0))

	//Dim返回max{x-y,0}
	fmt.Println("Dim(100，-99.5）=", math.Dim(100, -99.5))

	//下面三个函数功能不太明白，需要查询erf.go源代码后更新
	fmt.Println("Erf(100.5）=", math.Erf(100.5))
	fmt.Println("Erfinv(100.5）=", math.Erfinv(100.5))
	fmt.Println("Erfc(100.5）=", math.Erfc(100.5))
	fmt.Println("Erfcinv(100.5）=", math.Erfcinv(100.5))

	//Exp函数返回e^x结果
	fmt.Println("Exp(100)=", math.Exp(100))

	//Exp2函数返回2^x结果
	fmt.Println("Exp2(10)=", math.Exp2(10))

	//Expm1函数返回e^x - 1结果
	fmt.Println("Expm1(1)=", math.Expm1(1))

	//以下四个函数功能还没彻底弄清楚，后续更新
	fmt.Println("Float32bits(12.121)=", math.Float32bits(12.121))
	fmt.Println("Float64bits(12.121)=", math.Float64bits(12.121))
	fmt.Println("Float32frombits(1094840222)=", math.Float32frombits(1094840222))
	fmt.Println("Float64bits(4623013134440178123)=", math.Float64frombits(4623013134440178123))

	//向下圆整到最大的整数
	fmt.Printf("Floor(1.51)=%.1f\n", math.Floor(1.51))

	//Frexp分解参数f，满足f == frac × 2**exp
	//frac的取值范围在interval [½, 1)，且为绝对值
	frac, exp := math.Frexp(100.100)
	fmt.Printf("Frexp(100.100) returns frac=%f,exp=%d\n", frac, exp)
	frac, exp = math.Frexp(-100.100)
	fmt.Printf("Frexp(-100.100) returns frac=%f,exp=%d\n", frac, exp)

	//
	fmt.Println("Gamma(100)=", math.Gamma(100))

	//Hypot返回Sqrt(p*p + q*q)
	fmt.Println("Hypot(10,20)=", math.Hypot(10, 20))

	//Ilogb返回the binary exponent of x as an integer.
	fmt.Println("Ilogb(100)=", math.Ilogb(100))

	//Inf返回正负无穷大
	fmt.Println("Inf(111)=", math.Inf(111))
	fmt.Println("Inf(-91)=", math.Inf(-91))

	//IsInf判断是否是正负无穷大，正负有第二个参数指定
	fmt.Println("IsInf(1000,-1)", math.IsInf(1000, -1))
	fmt.Println("IsInf(1000,1)", math.IsInf(1000, 1))

	//IsNaN判断是否是一个数字：Not a Number——NaN
	fmt.Println("IsNaN(1000):", math.IsNaN(1000))
	fmt.Println("IsNaN(math.NaN()):", math.IsNaN(math.NaN()))
	fmt.Println("IsNaN(math.Inf(-1)):", math.IsNaN(math.Inf(-1)))

	//J0 returns the order-zero Bessel function of the first kind.
	fmt.Println("J0(100.1)=", math.J0(100.1))

	//J1 returns the order-one Bessel function of the first kind.
	fmt.Println("J1(100.1)=", math.J1(100.1))

	//Jn returns the order-n Bessel function of the first kind.
	fmt.Println("Jn(2,100.1)=", math.Jn(2, 100.1))

	//Ldexp is the inverse of Frexp. It returns frac × 2**exp.
	fmt.Println("Ldexp(100.11,2)=", math.Ldexp(100.11, 2))

	//Lgamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).
	lga, sign := math.Lgamma(100.11)
	fmt.Printf("Lgamma(100.11) return lgamma %f,sign %d\n", lga, sign)

	//Log返回参数x的自然对数值,Log10返回参数的log10的对数值
	fmt.Println("Log(10)=", math.Log(10))
	fmt.Println("Log10(10)=", math.Log10(10))

	//Log1p returns the natural logarithm of 1 plus its argument x.
	//It is more accurate than Log(1 + x) when x is near zero.
	fmt.Println("Log1p(10)=", math.Log1p(10))

	//Log2返回以2为底对数值
	fmt.Println("Log2(10)=", math.Log2(10))

	//Logb返回2进展的x的幂？？？
	fmt.Println("Logb(10)=", math.Logb(10))

	fmt.Println("Max(1,2)=", math.Max(1, 2))
	fmt.Println("Min(1,2)=", math.Min(1, 2))

	//Mod returns the floating-point remainder of x/y. The magnitude of the result
	//is less than y and its sign agrees with that of x.
	fmt.Println("Mod(98.1,33.99)=", math.Mod(98.1, 32.99))

	//Modf returns integer and fractional floating-point numbers that sum to f.
	//Both values have the same sign as f.
	in, fl := math.Modf(100.9)
	fmt.Printf("Modf(100.9) returns int %f,float %f\n", in, fl)

	//NaN返回一个Not a Number的值
	fmt.Println("NaN()=", math.NaN())

	//Nextafter returns the next representable float64 value after x towards y.
	fmt.Println("Nextafter(10.1,1000)", math.Nextafter(10.1, 1000))

	//Nextafter32 returns the next representable float32 value after x towards y.
	fmt.Println("Nextafter32(10.1,1000)", math.Nextafter32(10.1, 1000))

	//Pow returns x**y, the base-x exponential of y.
	fmt.Println("Pow(2.2,10.1)=", math.Pow(2.2, 10.1))

	//Pow10 returns 10**n, the base-10 exponential of n.
	fmt.Println("Pow10(2)=", math.Pow10(2))

	//Remainder returns the IEEE 754 floating-point remainder of x/y.
	fmt.Println("Remainder(122.4,33.9)=", math.Remainder(122.4, 33.9))

	//Round returns the nearest integer, rounding half away from zero.
	fmt.Println("Round(-9.999)=", math.Round(-9.999))
	fmt.Println("Round(9.999)=", math.Round(9.999))
	fmt.Println("Round(-4.499)=", math.Round(-4.999))
	fmt.Println("Round(4.499)=", math.Round(4.999))

	//RoundToEven returns the nearest integer, rounding ties to even.
	fmt.Println("RoundToEven(-9.999)=", math.RoundToEven(-9.999))
	fmt.Println("RoundToEven(9.999)=", math.RoundToEven(9.999))
	fmt.Println("RoundToEven(-4.499)=", math.RoundToEven(-4.999))
	fmt.Println("RoundToEven(4.499)=", math.RoundToEven(4.999))

	//Signbit returns true if x is negative or negative zero.
	fmt.Println("Signbit(-0.000001) is", math.Signbit(-0.000001))
	fmt.Println("Signbit(0.000001) is", math.Signbit(0.000001))

	fmt.Println("Sin(100)=", math.Sin(100))

	//Sincos returns Sin(x), Cos(x).
	s, c := math.Sincos(100)
	fmt.Printf("Sincos(100) return sin(100)=%f, cos(100)=%f\n", s, c)

	//Sinh returns the hyperbolic sine of x.
	fmt.Println("Sinh(100)=", math.Sinh(100))

	fmt.Println("Sqrt(100)=", math.Sqrt(100))
	fmt.Println("Tan(100)=", math.Tan(100))
	fmt.Println("Tanh(100)=", math.Tanh(100))

	//Trunc returns the integer value of x.
	fmt.Println("Trunc(101.934)=", math.Trunc(101.934))
	fmt.Println("Trunc(-101.934)=", math.Trunc(-101.934))

	//Y0 returns the order-zero Bessel function of the second kind.
	//Y1 returns the order-one Bessel function of the second kind.
	//Yn returns the order-n Bessel function of the second kind.
	fmt.Println("Y0(100.1)=", math.Y0(100.1))
	fmt.Println("Y1(100.1)=", math.Y1(100.1))
	fmt.Println("Yn(2,100.1)=", math.Yn(2, 100.1))
}
