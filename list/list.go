// Copyright 2019 Joshua J Baker. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package list a singly linked list.
package list

// T is a generic type
type T = int

// Tzero is the zero value for T
var Tzero = 0

type node struct {
	item T
	next *node
}

// List is a singly-linked list
type List struct {
	head *node
}

// Push adds an item to the front of the list
func (l *List) Push(item T) {
	l.head = &node{item: item, next: l.head}
}

// Pop remove the item from the front of the list and returns true when item
// exists or false when list is empty.
func (l *List) Pop() (T, bool) {
	if v := l.head; v != nil {
		l.head = l.head.next
		return v.item, true
	}
	return Tzero, false
}

// Range over each item in list. Return true from the iter function to continue
// iterating, or false to stop.
func (l *List) Range(iter func(item T) bool) {
	v := l.head
	for v != nil {
		if !iter(v.item) {
			return
		}
		v = v.next
	}
}
