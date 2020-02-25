/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Date
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Parse(layout, value string) (Time, error)
 ------------------------------------------------------------------------------------
 **Description:
 Parse parses a formatted string and returns the time value it represents. The layout
 defines the format by showing how the reference time, defined to be
 Mon Jan 2 15:04:05 -0700 MST 2006
 would be interpreted if it were the value; it serves as an example of the input format.
 The same interpretation will then be made to the input string.
 Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard and convenient
 representations of the reference time. For more information about the formats and the
 definition of the reference time, see the documentation for ANSIC and the other
 constants defined by this package. Also, the executable example for Time.Format
 demonstrates the working of the layout string in detail and is a good reference.
 Elements omitted from the value are assumed to be zero or, when zero is impossible,
 one, so parsing "3:04pm" returns the time corresponding to Jan 1, year 0, 15:04:00
 UTC (note that because the year is 0, this time is before the zero Time). Years must
 be in the range 0000..9999. The day of the week is checked for syntax but it is
 otherwise ignored.
 In the absence of a time zone indicator, Parse returns a time in UTC.
 When parsing a time with a zone offset like -0700, if the offset corresponds to a
 time zone used by the current location (Local), then Parse uses that location and
 zone in the returned time. Otherwise it records the time as being in a fabricated
 location with time fixed at the given zone offset.
 When parsing a time with a zone abbreviation like MST, if the zone abbreviation has
 a defined offset in the current location, then that offset is used. The zone
 abbreviation "UTC" is recognized as UTC regardless of location. If the zone
 abbreviation is unknown, Parse records the time as being in a fabricated location
 with the given zone abbreviation and a zero offset. This choice means that such a
 time can be parsed and reformatted with the same layout losslessly, but the exact
 instant used in the representation will differ by the actual zone offset. To avoid
 such problems, prefer time layouts that use a numeric zone offset, or use
 ParseInLocation.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Parse函数根据格式化的字符串layout分析字符串value，然后返回value所代表的时间值的结构体Time；
 2. layout定义了标准时间的显示格式（Mon Jan 2 15:04:05 -0700 MST 2006），还用来描述待分析的
 字符串。预定义的layout有ANSIC，UnixDate等，具体参见time.convar.go中的const;
 3. value中省略的元素默认为0，或者不可能出现0时为1。比如转换“3：04pm”返回的时间为Jan 1, year 0,
 15:04:00 UTC。年必须在0000..9999之间。星期会做语法检查，但是值会被忽略。
*************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.
	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Mon Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Tue Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println("longForm:", t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println("ShortForm", t)
	println()

	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error:", err) // Returns an error as the layout is not a valid time value

}
