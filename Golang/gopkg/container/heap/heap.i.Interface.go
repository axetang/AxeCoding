/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: heap
 **Element: heap.Interface
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Interface interface {
     sort.Interface
     Push(x interface{}) // add x as element Len()
     Pop() interface{}   // remove and return element Len() - 1.
 }
 ------------------------------------------------------------------------------------
 **Description:
 The Interface type describes the requirements for a type using the routines in this
 package. Any type that implements it may be used as a min-heap with the following
 invariants (established after Init has been called or if the data is empty or
    sorted):
 !h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
 Note that Push and Pop in this interface are for package heap's implementation to
 call. To add and remove things from the heap, use heap.Push and heap.Pop.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. 任何实现了本接口的类型都可以用于构建最小堆。最小堆可以通过heap.Init建立，数据是递增顺序
 或者空的话也是最小堆。最小堆的约束条件是：
    !h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
 2. 注意接口的Push和Pop方法是供heap包调用的，请使用heap.Push和heap.Pop来向一个堆添加或者
 删除元素。
 3. sort.Interface接口定义如下：
  type Interface interface {
     // Len is the number of elements in the collection.
     Len() int
     // Less reports whether the element with
     // index i should sort before the element with index j.
     Less(i, j int) bool
     // Swap swaps the elements with indexes i and j.
     Swap(i, j int)
 }
 4. 这是堆的接口，`heap`包里面的方法只是提供堆的一些堆算法操作，要想使用这些算法操作，就必须
 实现这些接口，每个接口方法都有具体的含义，堆本身的数据结构由这个接口的具体实现决定，可以是数组、
 列表;
 5. 把元素x存放到索引号为Len()的位置上，比如，一个列表中元素有7个，索引号从0开始，那么x将被存放到
 索引号为7的位置上，即最末尾;
 6. 把索引号为Len()-1的元素移除并返回这个被移除的元素。比如，如果这个堆是一个数组，那么才操作就是
 把数组的最后一个元素移除并返回。
*************************************************************************************/
package main

type myHeap []int // 定义一个堆，存储结构为数组

func main() {

}

//TestPop
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

//TestPush

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
