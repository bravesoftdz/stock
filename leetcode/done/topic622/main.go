package main

//MyCircularQueue 1
type MyCircularQueue struct {
	capacity int
	queue    []int
}

//MyCircularQueue
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{capacity: k, queue: make([]int, 0, k*2)}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.queue) == this.capacity {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if len(this.queue) == 0 {
		return false
	}

	this.queue = this.queue[1:]
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.queue) == 0
}

//IsFull 1
func (this *MyCircularQueue) IsFull() bool {
	return len(this.queue) == this.capacity
}

func main() {

}
