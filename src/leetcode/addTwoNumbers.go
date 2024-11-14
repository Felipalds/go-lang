package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) getSum() int {
	sum := l.Val
	for l.Next != nil {
		sum += l.Next.Val
		l = *l.Next
	}
	return sum
}

// As listas podem ter tamanhos diferentes
// Pode ter "passar um número (overflow)"
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	result := ListNode{}
	for l1.Next != nil {

	}

	return &result

}

// func (l ListNode) revert () *ListNode {
// 	// stack := make(int64[])
// }

// func main() {
// 	list1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
// 	list2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

// 	fmt.Println(list1.getSum())

// 	addTwoNumbers(list1, list2)

// }
