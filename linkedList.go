package main

import (
	"fmt"
)

type student struct {
	val  string
	next *student
}

type stuCreater func(string) *student

//打印
func (s *student) trans() {
	for s != nil {
		fmt.Println(*s)
		s = s.next
	}
}

func thisCreater(val string) *student {
	return &student{
		val: val,
	}
}

//链表反转
func (s *student) reverse() *student {
	if s == nil || s.next == nil {
		return nil
	}
	head := s
	p1 := s
	p2 := s.next
	for p1 != nil {
		if p2 == nil {
			//s = p1 //获取反转链表
			return p1
		}
		p3 := p2.next
		p2.next = p1
		if p1 == head {
			p1.next = nil
		}
		p1 = p2
		p2 = p3
	}
	return nil
}

//删个节点
func (s *student) delNode(val string) *student {
	//如果删除头节点
	if s.val == val {
		head := s
		s = head.next
		return s
	}
	p1 := s
	p2 := s.next
	for p1 != nil {
		if p2.val == val {
			p1.next = p2.next
			p2 = nil
			return nil
		}
		p1 = p1.next
		p2 = p1.next
	}
	return nil
}

//插入节点
func (s *student) insertNode(val string, sc stuCreater) {
	if val == "" {
		return //传进来的不能是空
	}
	p1 := s
	p2 := s.next
	for p1.val != val {
		p1 = p1.next
		p2 = p2.next
	}
	i := sc("insert")
	p1.next = i
	i.next = p2
}

func main() {
	head := &student{
		val: "head",
	}

	tail := head
	for i := 0; i < 10; i++ {
		stu := &student{
			val: fmt.Sprintf("val%d", i),
		}
		tail.next = stu
		tail = stu
	}
	head.trans()

	//测试反转
	//rev := head.reverse()
	//head = rev
	//head.trans()

	//测试删除
	// isNewHead := head.delNode("val8")
	// if isNewHead != nil {
	// 	head = isNewHead
	// }

	//测试插入
	head.insertNode("val1", thisCreater)

	//执行增删查改后的打印
	head.trans()

	// fmt.Println("------trans----")
	// for rev != nil {
	// 	fmt.Println(*rev)
	// 	rev = rev.next
	// }

}
