package _2_add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	offset := 0
	var currNode *ListNode
	var node *ListNode
	for l1 != nil && l2 != nil {
		value := offset
		if l1 != nil {
			value += l1.Val
		}
		if l2 != nil {
			value += l2.Val
		}

		offset = value / 10
		value %= 10
		if currNode == nil {
			currNode = &ListNode{
				Val:  value,
				Next: nil,
			}
			node = currNode
		} else {
			currNode.Next = &ListNode{
				Val:  value,
				Next: nil,
			}
			currNode = currNode.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if offset != 0 {
		currNode.Next = &ListNode{
			Val:  offset,
			Next: nil,
		}
	}

	return node
}
