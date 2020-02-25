/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: builtin
 **Element: builtin.make
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func make(t Type, size ...IntegerType) Type
 ------------------------------------------------------------------------------------
 **Description:
 The make built-in function allocates and initializes an object of type slice, map, or
 chan (only). Like new, the first argument is a type, not a value. Unlike new, make's
 return type is the same as the type of its argument, not a pointer to it. The
 specification of the result depends on the type:
	Slice: The size specifies the length. The capacity of the slice is equal to its
	length. A second integer argument may be provided to specify a different capacity;
	it must be no smaller than the length. For example, make([]int, 0, 10) allocates
	an underlying array of size 10 and returns a slice of length 0 and capacity 10
	that is backed by this underlying array.
	Map: An empty map is allocated with enough space to hold the specified number of
	elements. The size may be omitted, in which case a small starting size is allocated.
	Channel: The channel's buffer is initialized with the specified buffer capacity. If
	zero, or the size is omitted, the channel is unbuffered.
 ------------------------------------------------------------------------------------
 **要点总结:
 make内建函数分配并初始化一个类型为(SMC)slice、map、chan及interface的对象。与new相同的是，其第一个实参为类型，
 而非值。不同的是，make的返回类型与其参数相同，而非指向它的指针，其具体结果取决于具体的类型。
 	slice：
		size指定了切片长度, 该切片的容量等于其长度;
		第二个整数实参可用来指定不同的容量；
		它必须不小于其长度，因此 make([]int, 0, 10)；
		会分配一个长度为0，容量为10的切片。
 	map：
		初始分配取决于size，但产生的 map 长度为0;
		size可以省略，这种情况下就会分配一个最小的起始容量。
 	chan：
		信道的缓存根据指定的缓存容量初始化;
		若size为零或被省略，该信道即为无缓存的。
*************************************************************************************/
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
