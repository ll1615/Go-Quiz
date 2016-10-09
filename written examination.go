// 1.有如下代码
//   msg := []int{1,2,3,4,5,6,7,8,9};
//   sli1 := msg[2:3:4]
//   sli2 := msg[2:3]
//   请说明两个切片有什么不同?

// 两者的长度相同,都为1(3-2).
// sli1显示声明了容量,为首位下标只差:2(4-2)
// sli2容量默认为起始下标至原切片末节点:7(9-2)

package main

import "fmt"

func main() {
	msg := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli1 := msg[2:3:4]
	sli2 := msg[2:3]
	fmt.Println("len(sli1)", len(sli1), "cap(sli1)", cap(sli1))
	fmt.Println("len(sli2)", len(sli2), "cap(sli2)", cap(sli2))
}

// len(sli1)=1, cap(sli1)=2
// len(sli2)=1, cap(sli2)=7

//********************************************************************
