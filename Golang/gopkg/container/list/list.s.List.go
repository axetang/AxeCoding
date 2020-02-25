/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: list
 **Element: list.List
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type List struct {
     // contains filtered or unexported fields
 }
 func (l *List) Back() *Element
 func (l *List) Front() *Element
 func (l *List) Init() *List
 func (l *List) InsertAfter(v interface{}, mark *Element) *Element
 func (l *List) InsertBefore(v interface{}, mark *Element) *Element
 func (l *List) Len() int
 func (l *List) MoveAfter(e, mark *Element)
 func (l *List) MoveBefore(e, mark *Element)
 func (l *List) MoveToBack(e *Element)
 func (l *List) MoveToFront(e *Element)
 func (l *List) PushBack(v interface{}) *Element
 func (l *List) PushBackList(other *List)
 func (l *List) PushFront(v interface{}) *Element
 func (l *List) PushFrontList(other *List)
 func (l *List) Remove(e *Element) interface{}
 ------------------------------------------------------------------------------------
 **Description:
 List represents a doubly linked list. The zero value for List is an empty list ready
 to use.
 Back returns the last element of list l or nil if the list is empty.
 Front returns the first element of list l or nil if the list is empty.
 Init initializes or clears list l.
 InsertAfter inserts a new element e with value v immediately after mark and returns
 e. If mark is not an element of l, the list is not modified. The mark must not be
 nil.
 InsertBefore inserts a new element e with value v immediately before mark and
 returns e. If mark is not an element of l, the list is not modified. The mark must
 not be nil.
 Len returns the number of elements of list l. The complexity is O(1).
 MoveAfter moves element e to its new position after mark. If e or mark is not an
 element of l, or e == mark, the list is not modified. The element and mark must not
 be nil.
 MoveBefore moves element e to its new position before mark. If e or mark is not an
 element of l, or e == mark, the list is not modified. The element and mark must not
 be nil.
 MoveToBack moves element e to the back of list l. If e is not an element of l, the
 list is not modified. The element must not be nil.
 MoveToFront moves element e to the front of list l. If e is not an element of l,
 the list is not modified. The element must not be nil.
 PushBack inserts a new element e with value v at the back of list l and returns e.
 PushBackList inserts a copy of an other list at the back of list l. The lists l and
 other may be the same. They must not be nil.
 PushFront inserts a new element e with value v at the front of list l and returns e.
 PushFrontList inserts a copy of an other list at the front of list l. The lists l
 and other may be the same. They must not be nil.
 Remove removes e from l if e is an element of list l. It returns the element value
 e.Value. The element must not be nil.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. List代表一个双向链表。List零值为一个空的、可用的链表;
 1. Back方法获得最后一个节点的指针，如果链表长度为0，则为`nil`；
 2. Front方法获得第一个节点的指针，如果链表长度为0，则为`nil`;
 3. Init方法初始化或者清空链表，该方法调用后，链表的长度为0；
 4. InsertAfter方法把数据`value`插入到`mark`节点的后面，并返回这个被插入的节点*Element；
 5. InsertBefork方法把数据`value`插入到`mark`节点的前面，并返回这个被插入的节点*Element；
 6. Len方法返回链表长度；
 7. MoveAfter方法把e移动到元素mark *Element之后，如果e或mask不是链表的元素，或者e=mark，
 链表不变；e和mask不能为nil；
 8. MoveBefore方法把e移动到元素mark *Element之前，如果e或mask不是链表的元素，或者e=mark，
 链表不变；e和mask不能为nil；
 9. MoveToBack方法把e移动到链表末尾；
 10. MoveToFront方法把e移动到链表开头；
 11. PushBack方法把一个对象存到链表末尾，并返回这个节点元素*Element;
 12. PushBackList方法创建参数链表other的拷贝，并将本链表l的最后一个位置连接到本链表拷贝的第
 一个位置；
 13. PushFront方法把一个对象存到链表开头，并返回这个节点元素*Element;
 14. PushFrontList方法创建参数链表的other的拷贝，并将拷贝的最后一个位置连接到本链表l的第
 一个位置；
 15. Remove方法删除一个元素e。
*************************************************************************************/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	fmt.Println(l.Len()) // 输出：3

	fmt.Println("----Back------")
	e := l.Back()
	fmt.Println(e.Value)
	l.PushBack("d")
	e = l.Back()
	fmt.Println(e.Value, l.Len()) // 输出：a
	fmt.Println("-----List Content------")
	for el, i := l.Front(), 0; i < l.Len(); i++ {
		fmt.Print(el.Value, " ")
		el = el.Next()
	}
	fmt.Println()

	fmt.Println("----Front------")
	e = l.Front()
	fmt.Println(e.Value)
	l.PushFront("c")
	e = l.Front()
	fmt.Println(e.Value) // 输出：a
	fmt.Println(l.Len())
	fmt.Println("-----List Content------")
	for el, i := l.Front(), 0; i < l.Len(); i++ {
		fmt.Print(el.Value, " ")
		el = el.Next()
	}
	fmt.Println()

	fmt.Println("-----InsertAfter&InsertBefore--------")
	l.InsertAfter("bb", e)
	fmt.Println(e.Next().Value) // 输出：b
	l.InsertBefore("xx", e)
	fmt.Println(e.Prev().Value) // 输出：b

	fmt.Println("-----List Content------")
	for el, i := l.Front(), 0; i < l.Len(); i++ {
		fmt.Print(el.Value, " ")
		el = el.Next()
	}
	fmt.Println()

	fmt.Println("--------MoveAfter&MoveBefore&Remove")
	fmt.Println("Before move, l.len()=", l.Len())
	e = l.Front()
	mark := l.Back()
	l.MoveBefore(e, mark)
	l.MoveAfter(e, mark)
	fmt.Println("after move, l.len()=", l.Len())

	fmt.Println("-----List Content------")
	for el, i := l.Front(), 0; i < l.Len(); i++ {
		fmt.Print(el.Value, " ")
		el = el.Next()
	}
	fmt.Println()

	e = l.Back()
	l.Remove(e)

	fmt.Println("----Init----")
	l.Init()
	fmt.Println(l.Len()) // 输出：0

	fmt.Println("----PushBackList-----")
	TestPushBackList()
	fmt.Println("----PushFrontList-----")
	TestPushFrontList()
}

func TestPushFrontList() {
	l1 := list.New()
	l1.PushFront("a")
	l1.PushFront("b")
	fmt.Println(l1.Len()) // 输出：2

	l2 := list.New()
	l2.PushFront("c")
	l2.PushFront("d")

	l1.PushFrontList(l2)  // l2中所有节点的list字段都是l2，在l2的所有节点都加到l1的开头后，list字段编程了l1
	fmt.Println(l1.Len()) // 输出：4
	elementFromL1 := l1.Front()
	fmt.Println(elementFromL1.Value) // 输出：d
	l1.Remove(elementFromL1)         // 删除开头的元素
	fmt.Println(l1.Len())            // 输出：3
	fmt.Println(l1.Front().Value)    // 输出：c
}
func TestPushBackList() {
	l1 := list.New()
	l1.PushBack("a")
	l1.PushBack("b")
	fmt.Println(l1.Len()) // 输出：2

	l2 := list.New()
	l2.PushBack("c")
	l2.PushBack("d")

	l1.PushBackList(l2)   // l2中所有节点的list字段都是l2，在l2的所有节点都加到l1的末尾后，list字段编程了l1
	fmt.Println(l1.Len()) // 输出：4
	elementFromL1 := l1.Back()
	fmt.Println(elementFromL1.Value) // 输出：d
	l1.Remove(elementFromL1)         // 删除末尾的元素
	fmt.Println(l1.Len())            // 输出：3
	fmt.Println(l1.Back().Value)     // 输出：c
}
