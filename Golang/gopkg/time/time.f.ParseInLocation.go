/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.ParseInLocation
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseInLocation(layout, value string, loc *Location) (Time, error)
 func Unix(sec int64, nsec int64) Time
 ------------------------------------------------------------------------------------
 **Description:
 ParseInLocation is like Parse but differs in two important ways. First, in the
 absence of time zone information, Parse interprets a time as UTC; ParseInLocation
 interprets the time as in the given location. Second, when given a zone offset or
 abbreviation, Parse tries to match it against the Local location; ParseInLocation
 uses the given location.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseInLocation和Parse函数类似，但在两个方面不同，如下；
 2. 第一，如果layout和value字符串中缺少时区信息，则Parse将value被解释为UTC时间，但ParseInLocation将使用
 参数loc *Location指定的时区；
 3. 第二，如果layout和value字符串中给定了时区偏移或者缩写，Parse会使用这个时区信息而不是本地时区；
 但ParseInLocation会使用给定的时区信息。
 4. 注意：经实测，还存在layout与value中只有一个提供了时区信息的情况，结果也是不同的。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	//loc, _ := time.LoadLocation("Europe/Berlin")
	loc, _ := time.LoadLocation("Asia/Shanghai")

	//const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	const longForm = "Jan 2, 2006 at 3:04pm"
	//t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	//t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (MST)", loc)
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am", loc)
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)
}
