/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	cur := head
	temp := &Node{}
	tCur := temp
	visited := map[*Node]*Node{}
	// copy list first
	for cur != nil {
		n := &Node{Val: cur.Val}
		visited[cur] = n
		tCur.Next = n
		tCur = tCur.Next
		cur = cur.Next
	}
	// copy random pointers
	cur, tCur = head, temp
	for cur != nil {
		tCur.Next.Random = visited[cur.Random]
		tCur = tCur.Next
		cur = cur.Next
	}
	return temp.Next
}