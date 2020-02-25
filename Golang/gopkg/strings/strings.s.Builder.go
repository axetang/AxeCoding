/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: strings
 **Element: strings.Builder
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Builder struct {
    // contains filtered or unexported fields
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
 }
 func (b *Builder) Grow(n int)
 func (b *Builder) Len() int
 func (b *Builder) Reset()
 func (b *Builder) String() string
 func (b *Builder) Write(p []byte) (int, error)
 func (b *Builder) WriteByte(c byte) error
 func (b *Builder) WriteRune(r rune) (int, error)
 func (b *Builder) WriteString(s string) (int, error)
 ------------------------------------------------------------------------------------
 **Description:
 A Builder is used to efficiently build a string using Write methods. It minimizes
 memory copying. The zero value is ready to use. Do not copy a non-zero Builder.
 Grow grows b's capacity, if necessary, to guarantee space for another n bytes. After
 Grow(n), at least n bytes can be written to b without another allocation. If n is
 negative, Grow panics.
 Len returns the number of accumulated bytes; b.Len() == len(b.String()).
 Reset resets the Builder to be empty.
 String returns the accumulated string.
 Write appends the contents of p to b's buffer. Write always returns len(p), nil.
 WriteByte appends the byte c to b's buffer. The returned error is always nil.
 WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer. It
 returns the length of r and a nil error.
 WriteString appends the contents of s to b's buffer. It returns the length of s and
 a nil error.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Builder结构体用于通过Write方法高效构建一个字符串，它最小化了内存拷贝操作。零值可以使用，不会拷贝
 非零的Builder；
 1. Grow方法扩展Builder的容量。可以确保扩展参数n指定的字节数;
 2. Len方法返回Builder积聚的字节数，b.Len()== len(b.String())；
 3. String方法返回Builder积聚的字符串；
 4. Reset方法清空Builder；
 5. Write方法将参数p []byte内容写入到Builder缓存的尾部；
 6. WriteByte方法将参数c byte写入到Builder缓存的尾部；
 7. WriteRune方法将参数r rune写入到Builder缓存的尾部；
 8. WriteString方法将参数s string写入到Builder缓存的尾部；
 9. 感觉对Builder的理解还不够透彻，需要后续阅读其他资料后补充。
*************************************************************************************/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())

	fmt.Println("b.Len()", b.Len(), "\n", "b.String()", b.String(), len(b.String()))
	b.Reset()
	fmt.Println("b.Len()", b.Len(), "\n", "b.String()", b.String(), len(b.String()))

	b.Write([]byte("axe tang"))
	fmt.Println("b.Len()", b.Len(), "\n", "b.String()", b.String(), len(b.String()))
	b.WriteByte('.')
	fmt.Println("b.Len()", b.Len(), "\n", "b.String()", b.String(), len(b.String()))
	b.WriteRune('唐')
	fmt.Println("b.Len()", b.Len(), "\n", "b.String()", b.String(), len(b.String()))
	b.WriteString("飞 201902")
	fmt.Println("b.Len()", b.Len(), "\n", "b.String()", b.String(), len(b.String()))
}
