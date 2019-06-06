// Copyright 2019 Joshua J Baker. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package list

import (
	"testing"
)

// Tfrom is utility funtion that will attempt to automatically convert any
// value from is original type to the user-defined value of type T.
// This function is primarily helpful for _test.go files.
//
// The templated signature should *always* be the following.
func Tfrom(x T) T { return x }

func TestList(t *testing.T) {
	var l List
	x, ok := l.Pop()
	if ok {
		t.Fatal("expected false")
	}
	if x != Tzero {
		t.Fatalf("expected '%v', got '%v'", Tzero, x)
	}
	for i := 1; i <= 100; i++ {
		l.Push(Tfrom(i))
	}
	x, ok = l.Pop()
	if !ok {
		t.Fatal("expected true")
	}
	if x != Tfrom(100) {
		t.Fatalf("expected '%v', got '%v'", Tfrom(100), x)
	}
	start := 99
	var count int
	l.Range(func(item T) bool {
		if Tfrom(start) != item {
			t.Fatalf("expected '%v', got '%v", item, Tfrom(start))
		}
		count++
		start--
		return true
	})
	if count != 99 {
		t.Fatalf("expected '%v', got '%v'", 99, count)
	}
	count = 0
	l.Range(func(item T) bool {
		if count == 10 {
			return false
		}
		count++
		return true
	})
	if count != 10 {
		t.Fatalf("expected '%v', got '%v'", 10, count)
	}
}
