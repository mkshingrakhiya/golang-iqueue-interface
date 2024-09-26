package main

import (
	"errors"
	"fmt"
)

// IQueue interface defines the methods for an integer queue
type IQueue interface {
	Push(item int)
	Pop() (int, error)
	IsQueueEmpty() bool
}

// IntQueue is a concrete implementation of the IQueue interface
type IntQueue struct {
	items []int // Slice to hold queue items
}

// NewIntQueue creates a new IntQueue
func NewIntQueue() *IntQueue {
	return &IntQueue{
		items: []int{},
	}
}

// Push adds an item to the end of the queue
func (q *IntQueue) Push(item int) {
	q.items = append(q.items, item)
}

// Pop removes and returns the item at the front of the queue
func (q *IntQueue) Pop() (int, error) {
	if q.IsQueueEmpty() {
		return 0, errors.New("queue is empty")
	}
	item := q.items[0]    // Get the first item
	q.items = q.items[1:] // Remove the first item from the queue
	return item, nil
}

// IsQueueEmpty checks if the queue is empty
func (q *IntQueue) IsQueueEmpty() bool {
	return len(q.items) == 0
}

func main() {
	var queue IQueue = NewIntQueue()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if item, err := queue.Pop(); err == nil {
		fmt.Println("Popped item:", item) // Should print 1
	} else {
		fmt.Println(err)
	}

	fmt.Println("Is queue empty?", queue.IsQueueEmpty()) // Should print false

	// Pop remaining items
	for !queue.IsQueueEmpty() {
		if item, err := queue.Pop(); err == nil {
			fmt.Println("Popped item:", item) // Should print remaining items in FIFO order
		}
	}

	fmt.Println("Is queue empty?", queue.IsQueueEmpty()) // Should print true
}
