// Copyright 2019 Joshua J Baker. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package queue a priority queue.
package queue

// T is a generic type
type T = int

// Tzero is the zero default for T
var Tzero = 0

// Tlt is the less-than comparator function for T
func Tlt(a, b T) bool { return a < b }

// Tgt is the greater-than comparator function for T
func Tgt(a, b T) bool { return a > b }

// Teq is the equal-to comparator function for T
func Teq(a, b T) bool { return a == b }

// Queue is a simple queue
type Queue struct {
	nodes []T
	len   int
	size  int
}

// Push pushes item onto queue
func (q *Queue) Push(item T) {
	if q.nodes == nil {
		q.nodes = make([]T, 2)
	} else {
		q.nodes = append(q.nodes, Tzero)
	}
	i := q.len + 1
	j := i / 2
	for i > 1 && Tgt(q.nodes[j], item) {
		q.nodes[i] = q.nodes[j]
		i = j
		j = j / 2
	}
	q.nodes[i] = item
	q.len++
}

// Peek looks at item at top of queue
func (q *Queue) Peek() (T, bool) {
	if q.len == 0 {
		return Tzero, false
	}
	return q.nodes[1], true
}

// Pop removes and returns the item at the top of queue
func (q *Queue) Pop() (T, bool) {
	if q.len == 0 {
		return Tzero, false
	}
	n := q.nodes[1]
	q.nodes[1] = q.nodes[q.len]
	q.len--
	var j, k int
	i := 1
	for i != q.len+1 {
		k = q.len + 1
		j = 2 * i
		if j <= q.len && Tlt(q.nodes[j], q.nodes[k]) {
			k = j
		}
		if j+1 <= q.len && Tlt(q.nodes[j+1], q.nodes[k]) {
			k = j + 1
		}
		q.nodes[i] = q.nodes[k]
		i = k
	}
	return n, true
}

// Len returns the number of items in queue
func (q *Queue) Len() int {
	return q.len
}
