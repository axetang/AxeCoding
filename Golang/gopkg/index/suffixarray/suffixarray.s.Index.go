/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: suffixarray
 **Element: suffixarray.Index
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Index struct {
    // contains filtered or unexported fields
	data []byte
	sa   []int // suffix array for data; len(sa) == len(data)
 }
 func (x *Index) Bytes() []byte
 func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)
 func (x *Index) Lookup(s []byte, n int) (result []int)
 func (x *Index) Read(r io.Reader) error
 func (x *Index) Write(w io.Writer) error
 ------------------------------------------------------------------------------------
 **Description:
 Index implements a suffix array for fast substring search.
 Bytes returns the data over which the index was created. It must not be modified.
 FindAllIndex returns a sorted list of non-overlapping matches of the regular
 expression r, where a match is a pair of indices specifying the matched slice
 of x.Bytes(). If n < 0, all matches are returned in successive order. Otherwise,
 at most n matches are returned and they may not be successive. The result is nil if
 there are no matches, or if n == 0.
 Lookup returns an unsorted list of at most n indices where the byte string s occurs
 in the indexed data. If n < 0, all occurrences are returned. The result is nil if s
 is empty, s is not found, or n == 0. Lookup time is O(log(N)*len(s) + len(result))
 where N is the size of the indexed data.
 Read reads the index from r into x; x must not be nil.
 Write writes the index x to w.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. suffixarrayb包通过使用内存中的后缀树实现了对数级时间消耗的子字符串搜索;
 1. Bytes方法返回创建x *Index时提供的[]byte数据，注意不能修改返回值；
 2. Read方法从一个字符流r io.Reader中读取数据到索引；
 3. Lookup方法根据输入参数s []byte数组查找索引Index，查找的次数由参数n int指定，返回匹配结果在
 数据中的出现的位置，结果未排序。如果s为空，则结果为nil，如果n<0则查找全部出现的位置并返回；
 4. 返回一个正则表达式r的不重叠的匹配的经过排序的列表，一个匹配表示为一对指定了匹配结果的切片的索引
 （相对于x.Bytes())。如果n<0，返回全部匹配；如果n==0或匹配失败，返回nil；否则n为result的最大长度；
 5. 待学习了Regexp后再更新FindAllIndex方法功能。
*************************************************************************************/
package main

import (
	"bytes"
	"fmt"
	"index/suffixarray"
	"regexp"
)

func main() {

	data := []byte("abcd")
	index := suffixarray.New(data)
	fmt.Println(index.Bytes())         //[97,97,97,97]
	fmt.Println(string(index.Bytes())) //[97,97,97,97]

	fmt.Println("------Read&Write-------")
	TestReadWrite()
	fmt.Println("------Lookup-------")
	TestLookup()
	fmt.Println("------FindAllIndex-------")
	TestFindAllIndex()
}

func TestReadWrite() {
	data := []byte("ab")
	x := suffixarray.New(data)
	var buf bytes.Buffer
	if err := x.Write(&buf); err != nil {
		fmt.Println("failed writing index %s", err)
	}
	size := buf.Len()
	var y suffixarray.Index
	if err := y.Read(&buf); err != nil {
		fmt.Println("failed reading index %s ", err)
	}
	fmt.Println(string(y.Bytes())) //[97 98]
	size++
}

func TestLookup() {
	data := []byte("aabaacadaa")
	index := suffixarray.New(data)
	str := []byte("aa")
	res := index.Lookup(str, 1)
	fmt.Println(res) //[6]
	res = index.Lookup(str, 3)
	fmt.Println(res) //[6 5 4]
	res = index.Lookup(str, -1)
	fmt.Println(res) //[6 5 4 3 2 1 0]
}

func TestFindAllIndex() {
	data := []byte("aaaaaaa")
	index := suffixarray.New(data)
	str := "[a+]"
	r, _ := regexp.Compile(str)
	res := index.FindAllIndex(r, 1)
	fmt.Println(res) //[[0 1]]
	res = index.FindAllIndex(r, 3)
	fmt.Println(res) //[[0 1][1 2][2 3]]
}
