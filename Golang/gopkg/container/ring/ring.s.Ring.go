/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: ring
 **Element: ring.Ring
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Ring struct {
     Value interface{} // for use by client; untouched by this library
     // contains filtered or unexported fields
 }
 func (r *Ring) Do(f func(interface{}))
 func (r *Ring) Len() int
 func (r *Ring) Link(s *Ring) *Ring
 func (r *Ring) Move(n int) *Ring
 func (r *Ring) Next() *Ring
 func (r *Ring) Prev() *Ring
 func (r *Ring) Unlink(n int) *Ring
 ------------------------------------------------------------------------------------
 **Description:
 A Ring is an element of a circular list, or ring. Rings do not have a beginning or
 end; a pointer to any ring element serves as reference to the entire ring. Empty
 rings are represented as nil Ring pointers. The zero value for a Ring is a
 one-element ring with a nil Value.
 Do calls function f on each element of the ring, in forward order. The behavior of
 Do is undefined if f changes *r.
 Len computes the number of elements in ring r. It executes in time proportional to
 the number of elements.
 Link connects ring r with ring s such that r.Next() becomes s and returns the
 original value for r.Next(). r must not be empty.
 If r and s point to the same ring, linking them removes the elements between r and
 s from the ring. The removed elements form a subring and the result is a reference
 to that subring (if no elements were removed, the result is still the original
	value for r.Next(), and not nil).
 If r and s point to different rings, linking them creates a single ring with the
 elements of s inserted after r. The result points to the element following the last
 element of s after insertion.
 Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0) in the ring
 and returns that ring element. r must not be empty.
 Next returns the next ring element. r must not be empty.
 Prev returns the previous ring element. r must not be empty.
 Unlink removes n % r.Len() elements from the ring r, starting at r.Next(). If n %
 r.Len() == 0, r remains unchanged. The result is the removed subring. r must not
 be empty.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Ring结构体代表环形链表的一个元素，同时也代表链表本身。环形链表没有头尾；指向环形链表任一
 元素的指针都可以作为整个环形链表看待。Ring零值是具有一个（Value字段为nil的）元素的链表；
 1. Do方法对链表的每一个元素都执行f（正向顺序），注意如果f改变了*r，Do的行为是未定义的;
 2. Len方法返回环形链表中的元素个数，复杂度O(n);
 3. Link方法Link连接r和s，并返回r原本的后继元素r.Next()。r不能为空。如果r和s指向同一个环形
 链表，则会删除掉r和s之间的元素，删掉的元素构成一个子链表，返回指向该子链表的指针
 （r的原后继元素）；如果没有删除元素，则仍然返回r的原后继元素，而不是nil。如果r和s指向不同的
 链表，将创建一个单独的链表，将s指向的链表插入r后面，返回s原最后一个元素后面的元素（即r的原
 后继元素）。
 4. Move方法返回移动n个位置（n>=0向前移动，n<0向后移动）后的元素，r不能为空；
 5. Next方法返回后一个元素，r不能为空;
 4. Prev方法返回前一个元素，r不能为空；
 6. Unlink方法删除链表中n % r.Len()个元素，从r.Next()开始删除。如果n % r.Len() == 0，
 不修改r。返回删除的元素构成的链表，r不能为空。
*************************************************************************************/
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()
	fmt.Println("r.Len()=", n)

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	println("----Do------")
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	println("-----NextPrev------")
	fmt.Println(r.Value)
	fmt.Println(r.Next().Value)
	fmt.Println(r.Prev().Value)
	r = r.Next()
	fmt.Println("---")
	fmt.Println(r.Next().Value)
	fmt.Println(r.Prev().Value)

	println("------Move------")
	fmt.Println(r.Value)
	fmt.Println(r.Move(-1).Value)
	fmt.Println(r.Move(0).Value)
	fmt.Println(r.Move(1).Value)

	fmt.Println("-----TestLink-----")
	TestLink()
	fmt.Println("-----TestUnlink-----")
	TestUnlink()
}

func TestLink() {
	// Create two rings, r and s, of size 2
	r := ring.New(5)
	s := ring.New(5)

	// Get the length of the ring
	lr := r.Len()
	ls := s.Len()

	// Initialize r with 0s
	for i := 0; i < lr; i++ {
		r.Value = i + 1
		r = r.Next()
	}

	// Initialize s with 1s
	for j := 0; j < ls; j++ {
		s.Value = (j + 1) * 10
		s = s.Next()
	}

	fmt.Println("r.Value=", r.Value, "s.Value=", s.Value)
	// Link ring r and ring s
	rs := r.Link(s)

	// Iterate through the combined ring and print its contents
	rs.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}

func TestUnlink() {
	// Create a new ring of size 6
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Unlink three elements from r, starting from r.Next()
	fmt.Println("3%r.Len()=", 3%r.Len())
	fmt.Println("90%100=", 90%100)
	fmt.Println("111%100=", 111%100)
	r.Unlink(3)

	// Iterate through the remaining ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}
