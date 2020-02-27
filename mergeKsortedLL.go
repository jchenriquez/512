package mergeksortedll

// ListNode is a simple linkedlist node structure.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeKLists will merge are sorted lists into a sorted list
func MergeKLists(lists []*ListNode) *ListNode {

	voidHead := &ListNode{-1, nil}
	curr := voidHead

	for {
		var smallest *ListNode
		var smallestIndex int

		for index, list := range lists {
			if list != nil {
				if smallest == nil {
					smallest = list
					smallestIndex = index
				} else if list.Val < smallest.Val {
					smallest = list
					smallestIndex = index
				}
			}
		}

		if smallest == nil {
			break
		}

		nNode := ListNode{smallest.Val, nil}
		curr.Next = &nNode
		curr = curr.Next
		lists[smallestIndex] = lists[smallestIndex].Next
	}

	return voidHead.Next
}
