//go:build !no_gotd_slices
// +build !no_gotd_slices

// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// BotMenuButtonClassArray is adapter for slice of BotMenuButtonClass.
type BotMenuButtonClassArray []BotMenuButtonClass

// Sort sorts slice of BotMenuButtonClass.
func (s BotMenuButtonClassArray) Sort(less func(a, b BotMenuButtonClass) bool) BotMenuButtonClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotMenuButtonClass.
func (s BotMenuButtonClassArray) SortStable(less func(a, b BotMenuButtonClass) bool) BotMenuButtonClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotMenuButtonClass.
func (s BotMenuButtonClassArray) Retain(keep func(x BotMenuButtonClass) bool) BotMenuButtonClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s BotMenuButtonClassArray) First() (v BotMenuButtonClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotMenuButtonClassArray) Last() (v BotMenuButtonClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotMenuButtonClassArray) PopFirst() (v BotMenuButtonClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotMenuButtonClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotMenuButtonClassArray) Pop() (v BotMenuButtonClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsBotMenuButton returns copy with only BotMenuButton constructors.
func (s BotMenuButtonClassArray) AsBotMenuButton() (to BotMenuButtonArray) {
	for _, elem := range s {
		value, ok := elem.(*BotMenuButton)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// BotMenuButtonArray is adapter for slice of BotMenuButton.
type BotMenuButtonArray []BotMenuButton

// Sort sorts slice of BotMenuButton.
func (s BotMenuButtonArray) Sort(less func(a, b BotMenuButton) bool) BotMenuButtonArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotMenuButton.
func (s BotMenuButtonArray) SortStable(less func(a, b BotMenuButton) bool) BotMenuButtonArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotMenuButton.
func (s BotMenuButtonArray) Retain(keep func(x BotMenuButton) bool) BotMenuButtonArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s BotMenuButtonArray) First() (v BotMenuButton, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotMenuButtonArray) Last() (v BotMenuButton, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotMenuButtonArray) PopFirst() (v BotMenuButton, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotMenuButton
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotMenuButtonArray) Pop() (v BotMenuButton, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}