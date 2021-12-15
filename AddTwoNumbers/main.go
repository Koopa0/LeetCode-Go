package AddTwoNumbers

type ListNode struct {
	val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{val: 0}

	n1, n2, carry, current := 0, 0, 0, first

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.val
			l2 = l2.Next
		}
		current.Next = &ListNode{val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return first.Next
}
