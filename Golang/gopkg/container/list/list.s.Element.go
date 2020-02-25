/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: list
 **Element: list.Element
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}
 func (e *Element) Next() *Element
 func (e *Element) Prev() *Element
 ------------------------------------------------------------------------------------
 **Description:
 Element is an element of a linked list.
 Next returns the next list element or nil.
 Prev returns the previous list element or nil.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Element类型代表是双向链表的一个元素，包含四个成员：next，prev，list，Value;
 2. Next返回链表的后一个元素或者nil；
 3. Prev返回链表的前一个元素或者nil。
 4. 程序可参见结构体List的实例，这里不再赘述。
*************************************************************************************/
package main

func main() {
}
