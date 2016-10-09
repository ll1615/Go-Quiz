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

//************************************************************************************************************************************
// 2.方法中,方法的接受者类型,带*与不带*有什么区别?

//   带*:传入的是接受者类型的指针,可使用指针直接对原调用对象进行修改
// 不带*:传入的是接受者类型的实例,即原调用对象的拷贝,无法对原调用对象进行操作
package main

import "fmt"

func main() {
	smp := Sample{ID: 2}
	fmt.Println("raw ID:", smp.ID)
	smp.DoSomethingElse()
	fmt.Println("ID after calling function DoSomethingElse:", smp.ID)
	smp.DoSomething()
	fmt.Println("ID after calling DoSomething:", smp.ID)
}

type Sample struct {
	ID int
}

func (s *Sample) DoSomething() {
	fmt.Println("ID in function DoSomething:", s.ID)
	s.ID *= 10
}

func (s Sample) DoSomethingElse() {
	fmt.Println("ID in function DoSomethingElse:", s.ID)
	s.ID *= 10
}

// raw ID: 2
// ID in function DoSomethingElse: 2
// ID after calling function DoSomethingElse: 2
// ID in function DoSomething: 2
// ID after calling DoSomething: 20

//************************************************************************************************************************************
// 3.写一个简单的实例方法,实现面向对象的多态

