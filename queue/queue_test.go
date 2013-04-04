// See LICENSE file for copyright and license details.

package queue_test

import (
	"github.com/ozkriff/marauder/queue"
	"testing"
)

func TestBasicUsage(t *testing.T) {
	q := queue.Queue{}
	if !q.IsEmpty() {
		t.Errorf("New queue isn't empty")
	}
	q.Push(1)
	if q.IsEmpty() {
		t.Errorf("Queue shouldn't be empty")
	}
	if q.Pop().(int) != 1 {
		t.Errorf("Bad value poped from queue")
	}
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty")
	}
}
