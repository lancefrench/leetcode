/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	current int
	data    []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	var res PeekingIterator
	for iter.hasNext() {
		res.data = append(res.data, iter.next())
	}
	res.current = 0
	return &res
}

func (this *PeekingIterator) hasNext() bool {
	return this.current < len(this.data)
}

func (this *PeekingIterator) next() int {
	this.current++
	return this.data[this.current-1]
}

func (this *PeekingIterator) peek() int {
	return this.data[this.current]
}