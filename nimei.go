
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