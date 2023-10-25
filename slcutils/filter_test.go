package slcutils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name       string
		items      []int
		expKept    []int
		expRemoved []int
		removeFn   func(i int) bool
	}{
		{name: "remove nothing", items: []int{1, 1, 1}, expKept: []int{1, 1, 1}, expRemoved: []int{}, removeFn: func(i int) bool {
			return i == 2
		}},
		{name: "remove all", items: []int{1, 1, 1}, expKept: []int{}, expRemoved: []int{1, 1, 1}, removeFn: func(i int) bool {
			return i == 1
		}},
		{name: "remove half", items: []int{1, 2, 1, 2}, expKept: []int{1, 1}, expRemoved: []int{2, 2}, removeFn: func(i int) bool {
			return i == 2
		}},
		{name: "do nothing", items: []int{}, expKept: []int{}, expRemoved: []int{}, removeFn: func(i int) bool {
			return i == 2
		}},
		{name: "do nothing nil", items: nil, expKept: []int{}, expRemoved: []int{}, removeFn: func(i int) bool {
			return i == 2
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removed, kept := Remove(tt.items, tt.removeFn)
			require.Equal(t, tt.expKept, kept)
			require.Equal(t, tt.expRemoved, removed)
		})
	}
}

func TestKeep(t *testing.T) {
	tests := []struct {
		name       string
		items      []int
		expRemoved []int
		expKept    []int
		keepFn     func(i int) bool
	}{
		{name: "keep nothing", items: []int{1, 1, 1}, expRemoved: []int{1, 1, 1}, expKept: []int{}, keepFn: func(i int) bool {
			return i == 2
		}},
		{name: "keep all", items: []int{1, 1, 1}, expRemoved: []int{}, expKept: []int{1, 1, 1}, keepFn: func(i int) bool {
			return i == 1
		}},
		{name: "keep half", items: []int{1, 2, 1, 2}, expRemoved: []int{1, 1}, expKept: []int{2, 2}, keepFn: func(i int) bool {
			return i == 2
		}},
		{name: "do nothing", items: []int{}, expRemoved: []int{}, expKept: []int{}, keepFn: func(i int) bool {
			return i == 2
		}},
		{name: "do nothing nil", items: nil, expRemoved: []int{}, expKept: []int{}, keepFn: func(i int) bool {
			return i == 2
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kept, removed := Keep(tt.items, tt.keepFn)
			require.Equal(t, tt.expKept, kept)
			require.Equal(t, tt.expRemoved, removed)
		})
	}
}
