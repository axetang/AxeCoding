/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Location
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Location struct {
     // contains filtered or unexported fields
	name string
	zone []zone
	tx   []zoneTrans
	// Most lookups will be for the current time.
	// To avoid the binary search through tx, keep a
	// static one-element cache that gives the correct
	// zone for the time when the Location was created.
	// if cacheStart <= t < cacheEnd,
	// lookup can return cacheZone.
	// The units for cacheStart and cacheEnd are seconds
	// since January 1, 1970 UTC, to match the argument
	// to lookup.
	cacheStart int64
	cacheEnd   int64
	cacheZone  *zone
 }
 var Local *Location = &localLoc
 var UTC *Location = &utcLoc
 func FixedZone(name string, offset int) *Location
 func LoadLocation(name string) (*Location, error)
 func LoadLocationFromTZData(name string, data []byte) (*Location, error)
 func (l *Location) String() string
 ------------------------------------------------------------------------------------
 **Description:
 A Location maps time instants to the zone in use at that time. Typically, the Location
 represents the collection of time offsets in use in a geographical area, such as CEST
 and CET for central Europe.
 Local represents the system's local time zone.
 UTC represents Universal Coordinated Time (UTC).
 FixedZone returns a Location that always uses the given zone name and offset
 (seconds east of UTC).
 LoadLocation returns the Location with the given name.
 If the name is "" or "UTC", LoadLocation returns UTC. If the name is "Local",
 LoadLocation returns Local.
 Otherwise, the name is taken to be a location name corresponding to a file in the IANA
 Time Zone database, such as "America/New_York".
 The time zone database needed by LoadLocation may not be present on all systems,
 especially non-Unix systems. LoadLocation looks in the directory or uncompressed zip
 file named by the ZONEINFO environment variable, if any, then looks in known
 installation locations on Unix systems, and finally looks in $GOROOT/lib/time/
 zoneinfo.zip.
 LoadLocationFromTZData returns a Location with the given name initialized from the
 IANA Time Zone database-formatted data. The data should be in the format of a standard
 IANA time zone file (for example, the content of /etc/localtime on Unix systems).
 String returns a descriptive name for the time zone information, corresponding to the
 name argument to LoadLocation or FixedZone.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Location结构体所有成员都是非导出的；
 1. 注意几个变量定义关系：
  		var UTC *Location = &utcLoc
		var Local *Location = &localLoc
 		var utcLoc = Location{name: "UTC"}
		var localLoc Location
 2. FixedZone方法返回由参数name string和offset int（秒数）指定的位置信息*Location；
 3. LoadLocation返回name对应的Location。如果name是""或“UTC”,LoadLocation返回UTC，如果name
 是“Local”， LoadLocation返回Local。否则，name取为IANA时区数据库中的一个文件对应的位置名称，
 如“America/New_York”。LoadLocation使用的时区数据库可能不会被所有的系统所提供，尤其是非Unix系统。
 LoadLocation在ZONEINFO环境变量指定的目录或者解压的zip文件中寻找，如果没有，在Unix系统中已知的
 安装位置寻找，最后在 $GOROOT/lib/time/zoneinfo.zip中寻找;
 4. LoadLocationFromTZData函数的功能和使用方法需要后续查阅源代码后更新；
 5. String方法返回Location的字符串表达，它对应于LoadLocation和FixedZone方法的name string参数。
************************************************************************************/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----var UTC-----")
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// If the system has a timezone database present, it's possible to load a location
	// from that, e.g.:
	//    newYork, err := time.LoadLocation("America/New_York")
	// Creating a time requires a location. Common locations are time.Local and time.UTC.
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
	// 8 hours ahead so the two dates actually represent the same instant.
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	fmt.Println("-----FixedZone-----")
	loc := time.FixedZone("UTC-8", -8*60*60)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t.Format(time.RFC822))

	fmt.Println("-----LoadLocation-----")
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Println(loc)
	loc, _ = time.LoadLocation("Local")
	fmt.Println(loc)
	loc, _ = time.LoadLocation("UTC")
	fmt.Println(loc)

	fmt.Println("-----LoadLocationFromTZData-----")

	fmt.Println("-----String-----")
	loc = time.FixedZone("UTC-8", -8*60*60)
	fmt.Println(loc.String())
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Println(loc.String())

}
