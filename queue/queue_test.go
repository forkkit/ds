// Copyright 2019 Joshua J Baker. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package queue

import (
	"sort"
	"testing"
)

// Tfrom is utility funtion that will attempt to automatically convert any
// value from is original type to the user-defined value of type T.
// This function is primarily helpful for _test.go files.
//
// The templated signature should *always* be the following.
func Tfrom(x T) T { return x }

func TestQueue(t *testing.T) {
	var q Queue
	if _, ok := q.Peek(); ok {
		t.Fatal("expected false")
	}
	if _, ok := q.Pop(); ok {
		t.Fatal("expected false")
	}
	var items = []int{
		10, 2, 30, 12, 1, 99, 16,
	}
	for _, x := range items {
		q.Push(Tfrom(x))
	}
	sort.Ints(items)
	if q.Len() != len(items) {
		t.Fatalf("expected '%v', got '%v'", len(items), q.Len())
	}
	for len(items) > 0 {
		x, ok := q.Peek()
		if !ok {
			t.Fatal("expected true")
		}
		if x != Tfrom(items[0]) {
			t.Fatalf("expecting '%v', got '%v'", Tfrom(items[0]), x)
		}
		x, ok = q.Pop()
		if !ok {
			t.Fatal("expected true")
		}
		if x != Tfrom(items[0]) {
			t.Fatalf("expecting '%v', got '%v'", Tfrom(items[0]), x)
		}
		items = items[1:]
		if q.Len() != len(items) {
			t.Fatalf("expected '%v', got '%v'", len(items), q.Len())
		}
	}
	if _, ok := q.Peek(); ok {
		t.Fatal("expected false")
	}
	if _, ok := q.Pop(); ok {
		t.Fatal("expected false")
	}
}
