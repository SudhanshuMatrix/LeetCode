/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    groupPrev := dummy

    for {
        kth := groupPrev
        for i := 0; i < k && kth != nil; i++ {
            kth = kth.Next
        }

        if kth == nil {
            break
        }

        groupNext := kth.Next
        groupStart := groupPrev.Next

        prev := groupNext
        curr := groupStart

        for curr != groupNext {
            next := curr.Next
            curr.Next = prev
            prev = curr
            curr = next
        }

        groupPrev.Next = kth
        groupPrev = groupStart
    }

    return dummy.Next
}