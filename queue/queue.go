// See LICENSE file for copyright and license details.

// Package queue implements extremle inefficiemnt queue.
//
package queue

import (
	"container/list"
)

type Queue struct {
	l list.List
}

func (self *Queue) Push(value interface{}) {
	self.l.PushBack(value)
}

func (self *Queue) Pop() interface{} {
	tmp := self.l.Front()
	self.l.Remove(tmp)
	return tmp.Value
}

func (self *Queue) IsEmpty() bool {
	return self.l.Len() == 0
}
