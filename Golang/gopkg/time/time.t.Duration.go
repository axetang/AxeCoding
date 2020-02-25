/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Duration
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Duration int64
 func (d Duration) Hours() float64
 func (d Duration) Minutes() float64
 func (d Duration) Nanoseconds() int64
 func (d Duration) Round(m Duration) Duration
 func (d Duration) Seconds() float64
 func (d Duration) String() string
 func (d Duration) Truncate(m Duration) Duration

 const (
     Nanosecond  Duration = 1
     Microsecond          = 1000 * Nanosecond
     Millisecond          = 1000 * Microsecond
     Second               = 1000 * Millisecond
     Minute               = 60 * Second
     Hour                 = 60 * Minute
 )
 ------------------------------------------------------------------------------------
 **Description:
 A Duration represents the elapsed time between two instants as an int64 nanosecond
 count. The representation limits the largest representable duration to approximately
 290 years.
 Hours returns the duration as a floating point number of hours.
 Minutes returns the duration as a floating point number of minutes.
 Nanoseconds returns the duration as an integer nanosecond count.
 Round returns the result of rounding d to the nearest multiple of m. The rounding
 behavior for halfway values is to round away from zero. If the result exceeds the
 maximum (or minimum) value that can be stored in a Duration, Round returns the
 maximum (or minimum) duration. If m <= 0, Round returns d unchanged.
 Seconds returns the duration as a floating point number of seconds.
 String returns a string representing the duration in the form "72h3m0.5s". Leading
 zero units are omitted. As a special case, durations less than one second format use
 a smaller unit (milli-, micro-, or nanoseconds) to ensure that the leading digit is
 non-zero. The zero duration formats as 0s.
 Truncate returns the result of rounding d toward zero to a multiple of m. If m <= 0,
 Truncate returns d unchanged.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. time包使用Duration定义了时间的常用单位常数，最小的时间单位NanoSecond定义为1；
 1. Duration表示在两个时间实例之间逝去的时间段，用int64纳秒计数来表示；
 2. Hours方法返回时间d Duration以float64表示的小时数；
 3. Minutes方法返回时间d Duration以float64表示的分钟数；
 4. NanoSeconds方法返回时间d Duration以float64表示的纳秒数；
 5. Round方法返回时间d Duration以参数m Duration为园整单位标准的园整结果，并返回这个结果的duration；
 3. Seconds方法返回时间d Duration以float64表示的秒数；
 6. String方法根据时间返回一个"72h3m0.5s"形式的字符串，前导0会被省略。当时间小于1秒时，会用更小的
 单位（毫秒，微秒，纳秒）来保证前导数字不为0。时间就为0时则仅返回0，没有单位；
 7. Truncate方法以m Duration为基准向零方向园整d并返回，Truncate方法和Round方法执行结果略有差异。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----Hours-----")
	d, _ := time.ParseDuration("3h4m2s")
	fmt.Println(d.Hours())

	fmt.Println("-----TestRound&Truncate-----")
	TestRoundTruncate()
	fmt.Println("-----TestString-----")
	TestString()
}
func TestString() {
	//Println方法会自动调用d Duration的String方法
	d, _ := time.ParseDuration("1h2m3.5s")
	fmt.Println(d)
	d, _ = time.ParseDuration("3ms4ns")
	fmt.Println(d)
	d, _ = time.ParseDuration("0")
	fmt.Println(d)
}

func TestRoundTruncate() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s，d.Truncate= %s\n",
			r, d.Round(r).String(), d.Truncate(r).String())
	}

}
