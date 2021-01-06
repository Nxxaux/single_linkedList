package main

import (
	"context"
	//"time"

	//"context"
	//"container/list"
	"fmt"
)

//这是一个node的 构造函数
func nodeCreater(val string) *node {
	return &node{
		val: val,
	}
}

//双向链表
type node struct {
	val  string
	prev *node
	next *node
}

//普通遍历由head 到tail 遍历
func (s *node) trans() {
	for s != nil {
		fmt.Println(*s)
		s = s.next
	}
}

//反向遍历
func (s *node) rtrans(r *node) {
	for r != nil {
		fmt.Println(*r)
		r = r.prev
	}
}

//关于正向查找和反向查找，理论上比单链表至少快一倍
//正向查找
func iteratesByNext(ctx context.Context, ret chan *node, val string, n *node) {
	for n != nil {
		select {
		case <-ctx.Done():
		default:
			if n.val == val {
				//fmt.Println("n:", n.val)
				ret <- n
				return
			}
			//fmt.Println("n:", n.val)
			n = n.next
		}
	}
}

//反向查找
func iteratesByPrev(ctx context.Context, ret chan *node, val string, p *node) {
	//t := time.NewTicker(time.Microsecond * 300).C //延时300ms 不然太快了， 来不及看结果
	for p != nil {
		select {
		case <-ctx.Done():
		// case <-t:
		// 	if p.val == val {
		// 		ret <- p
		// 		return
		// 	}
		// 	//fmt.Println("P:", p.val)
		// 	p = p.prev
		// 	//default:
		default:
			if p.val == val {
				ret <- p
				return
			}
			//fmt.Println("P:", p.val)
			p = p.prev
			//default:
		}
	}
}

//查找,双向同时查找
func twoRoutersFind(val string, head, tail *node) *node {
	ctx, cancel := context.WithCancel(context.Background())
	ret := make(chan *node)
	go iteratesByNext(ctx, ret, val, head)
	go iteratesByPrev(ctx, ret, val, tail)
	for {
		select {
		case val := <-ret:
			//fmt.Println(*val)
			cancel() //结束所有协程
			return val
		default:
		}
	}
}

func delNode(p2 *node) {
	p1 := p2.prev
	p3 := p2.next
	p3.prev = p1
	p1.next = p3
}

func insertNode(p2 *node) {
	p3 := nodeCreater("insert")
	p1 := p2.prev
	p1.next = p3
	p2.prev = p3
	p3.prev = p1
	p3.next = p2
}

func main() {
	head := &node{
		val:  "head",
		prev: nil,
	}
	tail := head
	for i := 0; i < 10; i++ {
		stu := &node{
			val: fmt.Sprintf("val%d", i),
		}
		tail.next = stu
		stu.prev = tail
		tail = stu
	}
	//普通遍历
	//head.trans()
	//反向遍历
	head.rtrans(tail)
	//双向遍历查找
	val := twoRoutersFind("val3", head, tail)
	fmt.Println("--find out this val:", val)
	// delNode(val)
	// head.trans()
	// head.rtrans(tail)

	insertNode(val)

	head.trans()
	head.rtrans(tail)
}
