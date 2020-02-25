/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.CaseRange
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type CaseRange struct {
     Lo    uint32
     Hi    uint32
     Delta d
 }
 ------------------------------------------------------------------------------------
 **Description:
 CaseRange represents a range of Unicode code points for simple (one code point to
 one code point) case conversion. The range runs from Lo to Hi inclusive, with a
 fixed stride of 1. Deltas are the number to add to the code point to reach the code
 point for a different case for that character. They may be negative. If zero, it
 means the character is in the corresponding case. There is a special case
 representing sequences of alternating corresponding Upper and Lower pairs.
 It appears with a fixed Delta of
    {UpperLower, UpperLower, UpperLower}
 The constant UpperLower has an otherwise impossible delta value.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. CaseRange表示一个简单的Unicode码点范围，用于大小写转换；
 2. 在Lo和Hi范围内的码点，如果要转换成大写，只需要加上d[0]即可如果要转换为小写，只需要加上d[1]即可
 转换为大写，如果要转换为Title 格式，只需要加上d[2]即可；
 3. // CaseRanges 中 Delta 数组的索引。
 const (
	UpperCase = iota
	LowerCase
	TitleCase
	MaxCase
 )
 如果一个CaseRange中的Delta元素是UpperLower，则表示这个CaseRange是一个有着连续的大写小写大写
 小写的范围。也就是说，Lo是大写，Lo+1是小写，Lo+2是大写，Lo+3是小写 ... 一直到Hi为止。
 const (
	UpperLower = MaxRune + 1 // 不是一个有效的 Delta 元素
 )
*************************************************************************************/
package main

func main() {
}
