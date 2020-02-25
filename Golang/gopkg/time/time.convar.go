/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
 )
 These are predefined layouts for use in Time.Format and time.Parse. The reference
 time used in the layouts is the specific time:

 Mon Jan 2 15:04:05 MST 2006
 which is Unix time 1136239445. Since MST is GMT-0700, the reference time can be
 thought of as

 01/02 03:04:05PM '06 -0700
 To define your own format, write down what the reference time would look like
 formatted your way; see the values of constants like ANSIC, StampMicro or Kitchen for
 examples. The model is to demonstrate what the reference time looks like so that the
 Format and Parse methods can apply the same transformation to a general time value.

 Some valid layouts are invalid time values for time.Parse, due to formats such as _
 for space padding and Z for zone information.

 Within the format string, an underscore _ represents a space that may be replaced by a
 digit if the following number (a day) has two digits; for compatibility with
 fixed-width Unix time formats.

 A decimal point followed by one or more zeros represents a fractional second, printed
 to the given number of decimal places. A decimal point followed by one or more nines
 represents a fractional second, printed to the given number of decimal places, with
 trailing zeros removed. When parsing (only), the input may contain a fractional second
 field immediately after the seconds field, even if the layout does not signify its
 presence. In that case a decimal point followed by a maximal series of digits is
 parsed as a fractional second.

 Numeric time zone offsets format as follows:
 -0700  ±hhmm
 -07:00 ±hh:mm
 -07    ±hh
 Replacing the sign in the format with a Z triggers the ISO 8601 behavior of printing Z
 instead of an offset for the UTC zone. Thus:

 Z0700  Z or ±hhmm
 Z07:00 Z or ±hh:mm
 Z07    Z or ±hh
 The recognized day of week formats are "Mon" and "Monday". The recognized month formats
 are "Jan" and "January".

 Text in the format string that is not recognized as part of the reference time is
 echoed verbatim during Format and expected to appear verbatim in the input to Parse.

 The executable example for Time.Format demonstrates the working of the layout string
 in detail and is a good reference.

 Note that the RFC822, RFC850, and RFC1123 formats should be applied only to local times.
 Applying them to UTC times will use "UTC" as the time zone abbreviation, while strictly
 speaking those RFCs require the use of "GMT" in that case. In general RFC1123Z should
 be used instead of RFC1123 for servers that insist on that format, and RFC3339 should
 be preferred for new protocols. RFC3339, RFC822, RFC822Z, RFC1123, and RFC1123Z are
 useful for formatting; when used with time.Parse they do not accept all the time
 formats permitted by the RFCs. The RFC3339Nano format removes trailing zeros from the
 seconds field and thus may not sort correctly once formatted.

 const (
     Nanosecond  Duration = 1
     Microsecond          = 1000 * Nanosecond
     Millisecond          = 1000 * Microsecond
     Second               = 1000 * Millisecond
     Minute               = 60 * Second
     Hour                 = 60 * Minute
 )
 Common durations. There is no definition for units of Day or larger to avoid confusion
 across daylight savings time zone transitions.

 To count the number of units in a Duration, divide:

 second := time.Second
 fmt.Print(int64(second/time.Millisecond)) // prints 1000
 To convert an integer number of units to a Duration, multiply:

 seconds := 10
 fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
