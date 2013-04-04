// See LICENSE file for copyright and license details.

// Package queue реализует чертовски неэффективную очередь.
//
package queue

import (
	"container/list"
)

// Queue это простая очередь
type Queue struct {
	l list.List
}

// Push заносит элемент в очередь
func (self *Queue) Push(value interface{}) {
	self.l.PushBack(value)
}

// Pop выталкиевает первый элемент из очереди и возвращает его
func (self *Queue) Pop() interface{} {
	tmp := self.l.Front()
	self.l.Remove(tmp)
	return tmp.Value
}

// IsEmpty показывает пуста ли очередь
func (self *Queue) IsEmpty() bool {
	return self.l.Len() == 0
}
