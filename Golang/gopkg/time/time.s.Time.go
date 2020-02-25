/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Time
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // A Time represents an instant in time with nanosecond precision.
 //
 // Programs using times should typically store and pass them as values,
 // not pointers. That is, time variables and struct fields should be of
 // type time.Time, not *time.Time.
 //
 // A Time value can be used by multiple goroutines simultaneously except
 // that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
 // UnmarshalText are not concurrency-safe.
 //
 // Time instants can be compared using the Before, After, and Equal methods.
 // The Sub method subtracts two instants, producing a Duration.
 // The Add method adds a Time and a Duration, producing a Time.
 //
 // The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
 // As this time is unlikely to come up in practice, the IsZero method gives
 // a simple way of detecting a time that has not been initialized explicitly.
 //
 // Each Time has associated with it a Location, consulted when computing the
 // presentation form of the time, such as in the Format, Hour, and Year methods.
 // The methods Local, UTC, and In return a Time with a specific location.
 // Changing the location in this way changes only the presentation; it does not
 // change the instant in time being denoted and therefore does not affect the
 // computations described in earlier paragraphs.
 //
 // Representations of a Time value saved by the GobEncode, MarshalBinary,
 // MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
 // the location name. They therefore lose information about Daylight Saving Time.
 //
 // In addition to the required “wall clock” reading, a Time may contain an optional
 // reading of the current process's monotonic clock, to provide additional precision
 // for comparison or subtraction.
 // See the “Monotonic Clocks” section in the package documentation for details.
 //
 // Note that the Go == operator compares not just the time instant but also the
 // Location and the monotonic clock reading. Therefore, Time values should not
 // be used as map or database keys without first guaranteeing that the
 // identical Location has been set for all values, which can be achieved
 // through use of the UTC or Local method, and that the monotonic clock reading
 // has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
 // to t == u, since t.Equal uses the most accurate comparison available and
 // correctly handles the case when only one of its arguments has a monotonic
 // clock reading.
 //
 type Time struct {
	// wall and ext encode the wall time seconds, wall time nanoseconds,
	// and optional monotonic clock reading in nanoseconds.
	//
	// From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
	// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
	// The nanoseconds field is in the range [0, 999999999].
	// If the hasMonotonic bit is 0, then the 33-bit field must be zero
	// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
	// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
	// unsigned wall seconds since Jan 1 year 1885, and ext holds a
	// signed 64-bit monotonic clock reading, nanoseconds since process start.
	wall uint64
	ext  int64

	// loc specifies the Location that should be used to
	// determine the minute, hour, month, day, and year
	// that correspond to this Time.
	// The nil location means UTC.
	// All UTC times are represented with loc==nil, never loc==&utcLoc.
	loc *Location
 }

 const (
	hasMonotonic = 1 << 63
	maxWall      = wallToInternal + (1<<33 - 1) // year 2157
	minWall      = wallToInternal               // year 1885
	nsecMask     = 1<<30 - 1
	nsecShift    = 30
 )
 func (t Time) Add(d Duration) Time
 func (t Time) AddDate(years int, months int, days int) Time
 func (t Time) After(u Time) bool
 func (t Time) AppendFormat(b []byte, layout string) []byte
 func (t Time) Before(u Time) bool
 func (t Time) Clock() (hour, min, sec int)
 func (t Time) Date() (year int, month Month, day int)
 func (t Time) Day() int
 func (t Time) Equal(u Time) bool
 func (t Time) Format(layout string) string
 func (t *Time) GobDecode(data []byte) error
 func (t Time) GobEncode() ([]byte, error)
 func (t Time) Hour() int
 func (t Time) ISOWeek() (year, week int)
 func (t Time) In(loc *Location) Time
 func (t Time) IsZero() bool
 func (t Time) Local() Time
 func (t Time) Location() *Location
 func (t Time) MarshalBinary() ([]byte, error)
 func (t Time) MarshalJSON() ([]byte, error)
 func (t Time) MarshalText() ([]byte, error)
 func (t Time) Minute() int
 func (t Time) Month() Month
 func (t Time) Nanosecond() int
 func (t Time) Round(d Duration) Time
 func (t Time) Second() int
 func (t Time) String() string
 func (t Time) Sub(u Time) Duration
 func (t Time) Truncate(d Duration) Time
 func (t Time) UTC() Time
 func (t Time) Unix() int64
 func (t Time) UnixNano() int64
 func (t *Time) UnmarshalBinary(data []byte) error
 func (t *Time) UnmarshalJSON(data []byte) error
 func (t *Time) UnmarshalText(data []byte) error
 func (t Time) Weekday() Weekday
 func (t Time) Year() int
 func (t Time) YearDay() int
 func (t Time) Zone() (name string, offset int)
 ------------------------------------------------------------------------------------
 **Description:
 Add returns the time t+d.
 AddDate returns the time corresponding to adding the given number of years, months,
 and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011 returns
 March 4, 2010.
 AddDate normalizes its result in the same way that Date does, so, for example, adding
 one month to October 31 yields December 1, the normalized form for November 31.
 After reports whether the time instant t is after u.
 AppendFormat is like Format but appends the textual representation to b and returns
 the extended buffer.
 Before reports whether the time instant t is before u.
 Clock returns the hour, minute, and second within the day specified by t.
 Date returns the year, month, and day in which t occurs.
 Day returns the day of the month specified by t.
 Equal reports whether t and u represent the same time instant. Two times can be equal
 even if they are in different locations. For example, 6:00 +0200 CEST and 4:00 UTC are
 Equal. See the documentation on the Time type for the pitfalls of using == with Time
 values; most code should use Equal instead.
 Format returns a textual representation of the time value formatted according to
 layout, which defines the format by showing how the reference time, defined to be
 Mon Jan 2 15:04:05 -0700 MST 2006
 would be displayed if it were the value; it serves as an example of the desired
 output. The same display rules will then be applied to the time value.
 A fractional second is represented by adding a period and zeros to the end of the
 seconds section of layout string, as in "15:04:05.000" to format a time stamp with
 millisecond precision.
 Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard and
 convenient representations of the reference time. For more information about the
 formats and the definition of the reference time, see the documentation for ANSIC
 and the other constants defined by this package.
 GobDecode implements the gob.GobDecoder interface.
 GobEncode implements the gob.GobEncoder interface.
 Hour returns the hour within the day specified by t, in the range [0, 23].
 ISOWeek returns the ISO 8601 year and week number in which t occurs. Week ranges from
 1 to 53. Jan 01 to Jan 03 of year n might belong to week 52 or 53 of year n-1, and
 Dec 29 to Dec 31 might belong to week 1 of year n+1.
 In returns a copy of t representating the same time instant, but with the copy's
 location information set to loc for display purposes.
 In panics if loc is nil.
 IsZero reports whether t represents the zero time instant, January 1, year 1, 00:00:00
 UTC.
 Local returns t with the location set to local time.
 Location returns the time zone information associated with t.
 MarshalBinary implements the encoding.BinaryMarshaler interface.
 MarshalJSON implements the json.Marshaler interface. The time is a quoted string in
 RFC 3339 format, with sub-second precision added if present.
 MarshalText implements the encoding.TextMarshaler interface. The time is formatted in
 RFC 3339 format, with sub-second precision added if present.
 Minute returns the minute offset within the hour specified by t, in the range [0, 59].
 Month returns the month of the year specified by t.
 Nanosecond returns the nanosecond offset within the second specified by t, in the
 range [0, 999999999].
 Round returns the result of rounding t to the nearest multiple of d (since the zero
    time). The rounding behavior for halfway values is to round up. If d <= 0, Round
    returns t stripped of any monotonic clock reading but otherwise unchanged.
 Round operates on the time as an absolute duration since the zero time; it does not
 operate on the presentation form of the time. Thus, Round(Hour) may return a time
 with a non-zero minute, depending on the time's Location.
 Second returns the second offset within the minute specified by t, in the range [0, 59].
 String returns the time formatted using the format string
 "2006-01-02 15:04:05.999999999 -0700 MST"
 If the time has a monotonic clock reading, the returned string includes a final field
 "m=±<value>", where value is the monotonic clock reading formatted as a decimal number
 of seconds.
 The returned string is meant for debugging; for a stable serialized representation, use
 t.MarshalText, t.MarshalBinary, or t.Format with an explicit format string.
 Sub returns the duration t-u. If the result exceeds the maximum (or minimum) value
 that can be stored in a Duration, the maximum (or minimum) duration will be returned.
 To compute t-d for a duration d, use t.Add(-d).
 Truncate returns the result of rounding t down to a multiple of d (since the zero time).
 If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise
 unchanged.
 Truncate operates on the time as an absolute duration since the zero time; it does not
 operate on the presentation form of the time. Thus, Truncate(Hour) may return a time
 with a non-zero minute, depending on the time's Location.
 UTC returns t with the location set to UTC.
 Unix returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.
 The result does not depend on the location associated with t.
 UnixNano returns t as a Unix time, the number of nanoseconds elapsed since January 1,
 1970 UTC. The result is undefined if the Unix time in nanoseconds cannot be represented
 by an int64 (a date before the year 1678 or after 2262). Note that this means the
 result of calling UnixNano on the zero Time is undefined. The result does not depend
 on the location associated with t.
 UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
 UnmarshalJSON implements the json.Unmarshaler interface. The time is expected to be a
 quoted string in RFC 3339 format.
 UnmarshalText implements the encoding.TextUnmarshaler interface. The time is expected
 to be in RFC 3339 format.
 Weekday returns the day of the week specified by t.
 Year returns the year in which t occurs.
 YearDay returns the day of the year specified by t, in the range [1,365] for non-leap
 years, and [1,366] in leap years.
 Zone computes the time zone in effect at time t, returning the abbreviated name of the
 zone (such as "CET") and its offset in seconds east of UTC.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Time结构体表示一个带有纳秒精读的时间实例；需要注意的是，程序中使用Time结构体时，要使用其值形态
 参数而不是指针形态参数。即time.Time，而不是*time.Time。Time的零值是January 1, year 1,
 00:00:00.000000000 UTC；
 1. Add方法返回时间（t+d）
 2. AddDate方法生成由t Time加上参数中指定的年月日之后的时间，比如AddDate(-1, 2, 3) 使January 1,
 2011变为March 4, 2010。AddDate以和Date方法一样的方式标准化返回值，比如10月31日加一个月之后为11月
 31日，函数会将其标准化为12月1日；
 3. After方法判断t Time是否在参数u Time之后；
 4. AppendFormat方法把t Time代表的时间按照layout字符串要求加到参数b []byte尾部并返回b；
 5. Before方法判断t Time是否在参数u Time之前；
 6. Clock方法返回由t Time指定的时间的时、分、秒数字；
 7. Date方法返回由t Time指定的年、月、日数字；
 8. Day方法返回由t Time指定的日期数字；
 9. Format方法返回根据layout指定的格式格式化之后的字符串，layout定义了标准时间的显示格式。
 预定义的layout有ANSIC，UnixDate，RFC3339等，参见convar；
 10. GobEncode实现了gob.GobEncoder接口，具体功能待了解gob包后再更新；
 11. GobDecode实现了gob.Decode接口，具体功能待了解gob包后再更新；
 12. Hour方法返回由t Time指定的小时数字；
 13. ISOWeek方法返回时间t Time的在一年中的星期数。星期范围从1到53。每年的1月1日到3日可能会属于上一年的第52
 或者53周，每年的12月29日到31日可能属于下一年的第一周；
 14. In方法返回location设为参数loc的t Time，如果loc为nil则函数panic；
 15. IsZero方法返回时间t Time是否是Zero；
 16. Local方法返回t Time对应的本地时间；
 17. Location方法返回t Time对应的时区信息loc *Location；
 18. MarshalJSON方法实现了json.Marshaler接口；时间t Time按照RFC3339格式化；
 19. MarshalBinary方法实现了encoding.BinaryMarshaler interface；
 20. MarshalText实现了encoding.TextMarshaler interface. 时间t Time按照RFC 3339格式化；
 21. Minute方法返回t Time的分钟数字；
 22. Month方法返回t Time的月份信息Month；
 23. Round方法依据参数d Duration基准来进行时间圆整；
 24. Second方法返回t Time的秒数字；
 25. String方法返回t Time的字符串表达；
 26. Sub方法返回t-u，以Duration方式返回；
 27. Truncate方法依据参数d Duration向零方向圆整时间并返回；
 28. UTC方法返回t的UTC时间；
 29. Unix方法返回Unix时间，即从January 1, 1970 UTC起已经流逝的秒数；
 30. UnixNano方法返回Unix时间，即从January 1, 1970 UTC起已经流逝的纳秒数；
 31. UnmarshalJSON方法实现了json.Unmarshaler接口；时间t Time按照RFC3339格式化；
 32. UnmarshalBinary方法实现了encoding.BinaryUnmarshaler interface；
 33. UnmarshalText实现了encoding.TextUnmarshaler interface. 时间t Time按照RFC 3339格式化；
 34. Year方法返回t Time对应的年份；
 35. YearDay方法返回t Time指定的日期在当年的天数；
 36. Zone方法返回t Time对应的时区，以及相对UTC的offset。
 ************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----Add-----")
	now := time.Now()
	fmt.Println("now:", now)
	fmt.Println("after 3 hours", now.Add(3*time.Hour))

	fmt.Println("-----AddDate-----")
	start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)
	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)

	fmt.Println("-----After-----")
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)
	isYear3000AfterYear2000 := year3000.After(year2000) // True
	isYear2000AfterYear3000 := year2000.After(year3000) // False
	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

	fmt.Println("-----AppendFormat-----")
	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")
	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))

	fmt.Println("-----Before-----")
	year2000 = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 = time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // True
	isYear3000BeforeYear2000 := year3000.Before(year2000) // False

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)

	fmt.Println("-----Clock-----")
	now = time.Now()
	h, m, s := now.Clock()
	fmt.Println("now:", now)
	fmt.Printf("clock is %d %d %d\n", h, m, s)

	fmt.Println("-----Date-----")
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

	fmt.Println("-----Day-----")
	d = time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	day = d.Day()
	fmt.Printf("day = %v\n", day)

	fmt.Println("-----Equal-----")
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	// Unlike the equal operator, Equal is aware that d1 and d2 are the
	// same instant but in different time zones.
	d1 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	d2 := time.Date(2000, 2, 1, 20, 30, 0, 0, beijing)
	datesEqualUsingEqualOperator := d1 == d2
	datesEqualUsingFunction := d1.Equal(d2)
	fmt.Printf("datesEqualUsingEqualOperator = %v\n", datesEqualUsingEqualOperator)
	fmt.Printf("datesEqualUsingFunction = %v\n", datesEqualUsingFunction)

	fmt.Println("-----TestFormat-----")
	TestFormat()

	fmt.Println("-----Hour-----")
	d = time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	hour := d.Hour()
	fmt.Printf("hour = %v\n", hour)

	fmt.Println("-----ISOWeek-----")
	d = time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, week := d.ISOWeek()
	fmt.Printf("ISOWeek = Year %d Week %d\n", year, week)

	fmt.Println("-----In-----")
	t = time.Now()
	fmt.Println(t)
	fmt.Println(t.In(time.UTC))

	fmt.Println("-----IsZero-----")
	t = time.Date(2019, 2, 24, 21, 0, 0, 29, time.UTC)
	fmt.Printf("%s is zero time: %t \n", t, t.IsZero())

	fmt.Println("-----Local-----")
	t = time.Date(2019, 2, 24, 21, 0, 0, 29, time.UTC)
	fmt.Printf("%s local time is %s\n", t, t.Local())

	fmt.Println("-----Location-----")
	layout1 := "Mon Jan 2, 2006 at 3:04pm (MST)"
	t, _ = time.Parse(layout1, "Tue Feb 3, 2013 at 7:54pm (PST)")
	fmt.Printf("%s location: %s\n", t, t.Location())
	fmt.Println("local time is ", t.Local())

	fmt.Println("-----MarshalBinary-----")

	fmt.Println("-----MarshalText-----")

	fmt.Println("-----MarshalJSON-----")

	fmt.Println("-----Minute-----")
	t = time.Date(2019, 2, 24, 21, 9, 0, 29, time.UTC)
	fmt.Printf("%s ，Minute is %d\n", t, t.Minute())

	fmt.Println("-----Month-----")
	t = time.Date(2019, 2, 24, 21, 0, 0, 29, time.UTC)
	fmt.Printf("%s ，Month is %s\n", t, t.Month())

	fmt.Println("-----Nanosecond-----")
	t = time.Date(2019, 2, 24, 21, 0, 0, 29, time.UTC)
	fmt.Printf("%s ，NanoSecond is %d\n", t, t.Nanosecond())

	fmt.Println("-----Round-----")
	t = time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
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
	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}

	fmt.Println("-----Second-----")
	t = time.Date(2019, 2, 24, 21, 0, 10, 29, time.UTC)
	fmt.Printf("%s ，Second is %d\n", t, t.Second())

	fmt.Println("-----String-----")
	t = time.Date(2019, 2, 24, 21, 0, 0, 29, time.UTC)
	fmt.Printf("%s  String：%s\n", t, t.String())
	timeWithNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC)
	withNanoseconds := timeWithNanoseconds.String()
	timeWithoutNanoseconds := time.Date(2000, 2, 1, 12, 13, 14, 0, time.UTC)
	withoutNanoseconds := timeWithoutNanoseconds.String()
	fmt.Printf("withNanoseconds = %v\n", string(withNanoseconds))
	fmt.Printf("withoutNanoseconds = %v\n", string(withoutNanoseconds))

	fmt.Println("-----Sub-----")
	start = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	difference1 := end.Sub(start)
	difference2 := start.Sub(end)
	fmt.Printf("difference1 = %v\n", difference1)
	fmt.Printf("difference2 = %v\n", difference2)

	fmt.Println("-----Truncate-----")
	t, _ = time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
	}
	for _, d := range trunc {
		fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
	// To round to the last midnight in the local timezone, create a new Date.
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	_ = midnight

	fmt.Println("-----UTC-----")
	layout2 := "Mon Jan 2, 2006 at 3:04pm (MST)"
	t, _ = time.Parse(layout2, "Tue Feb 3, 2013 at 7:54pm (CST)")
	t = t.UTC()
	fmt.Println(t)

	fmt.Println("-----Unix&UnixNano-----")
	// 1 billion seconds of Unix, three ways.
	fmt.Println(time.Unix(1e9, 0).UTC())     // 1e9 seconds
	fmt.Println(time.Unix(0, 1e18).UTC())    // 1e18 nanoseconds
	fmt.Println(time.Unix(2e9, -1e18).UTC()) // 2e9 seconds - 1e18 nanoseconds
	t = time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
	fmt.Println(t.Unix())     // seconds since 1970
	fmt.Println(t.UnixNano()) // nanoseconds since 1970

	fmt.Println("-----Year-----")
	d = time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year = d.Year()
	fmt.Printf("year = %v\n", year)

	fmt.Println("-----YearDay-----")
	d = time.Date(2000, 12, 1, 12, 30, 0, 0, time.UTC)
	yd := d.YearDay()
	fmt.Printf("YearDay = %v\n", yd)

	fmt.Println("-----Zone-----")
	fmt.Println(time.Now().Zone())

}

func TestFormat() {
	// Parse a time value from a string in the standard Unix format.
	t, err := time.Parse(time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	// time.Time's Stringer method is useful without any format.
	fmt.Println("default format:", t)

	// Predefined constants in the package implement common layouts.
	fmt.Println("Unix format:", t.Format(time.UnixDate))

	// The time zone attached to the time value affects its output.
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

	// The rest of this function demonstrates the properties of the
	// layout string used in the format.

	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	// There are some wrinkles illustrated below.

	// Most uses of Format and Parse use constant layout strings such as
	// the ones defined in this package, but the interface is flexible,
	// as these examples show.

	// Define a helper function to make the examples' output look nice.
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-15s %q gives %q\n", name, layout, got)
	}

	// Print a header in our output.
	fmt.Printf("\nFormats:\n\n")

	// A simple starter example.
	do("Basic", "Mon Jan 2 15:04:05 MST 2006", "Sat Mar 7 11:06:39 PST 2015")

	// For fixed-width printing of values, such as the date, that may be one or
	// two characters (7 vs. 07), use an _ instead of a space in the layout string.
	// Here we print just the day, which is 2 in our layout string and 7 in our
	// value.
	do("No pad", "<2>", "<7>")

	// An underscore represents a space pad, if the date only has one digit.
	do("Spaces", "<_2>", "< 7>")

	// A "0" indicates zero padding for single-digit values.
	do("Zeros", "<02>", "<07>")

	// If the value is already the right width, padding is not used.
	// For instance, the second (05 in the reference time) in our value is 39,
	// so it doesn't need padding, but the minutes (04, 06) does.
	do("Suppressed pad", "04:05", "06:39")

	// The predefined constant Unix uses an underscore to pad the day.
	// Compare with our simple starter example.
	do("Unix", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")

	// The hour of the reference time is 15, or 3PM. The layout can express
	// it either way, and since our value is the morning we should see it as
	// an AM time. We show both in one format string. Lower case too.
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	// When parsing, if the seconds value is followed by a decimal point
	// and some digits, that is taken as a fraction of a second even if
	// the layout string does not represent the fractional second.
	// Here we add a fractional second to our time value used above.
	t, err = time.Parse(time.UnixDate, "Sat Mar  7 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}
	// It does not appear in the output if the layout string does not contain
	// a representation of the fractional second.
	do("No fraction", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")

	// Fractional seconds can be printed by adding a run of 0s or 9s after
	// a decimal point in the seconds value in the layout string.
	// If the layout digits are 0s, the fractional second is of the specified
	// width. Note that the output has a trailing zero.
	do("0s for fraction", "15:04:05.00000", "11:06:39.12340")

	// If the fraction in the layout is 9s, trailing zeros are dropped.
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")

}
