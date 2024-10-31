package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type op func() (string, error)
type wrappedOp func(*Stack) op

var (
	forward = func(s *Stack) op {
		return s.Forward
	}
	back = func(s *Stack) op {
		return s.Back
	}
	push = func(branch string) func(s *Stack) op {
		return func(s *Stack) op {
			return func() (string, error) {
				s.Push(branch)
				return "", nil
			}
		}
	}
)

func TestStack(t *testing.T) {
	cases := []*struct {
		name           string
		stack          *Stack
		ops            []wrappedOp
		expectedStack  []string
		expectedCursor int
		expectedErr    error
	}{
		{
			name: "multiple backward",
			stack: &Stack{
				Stack:  []string{"main", "foo", "bar", "baz"},
				Cursor: 2,
			},
			ops:            []wrappedOp{back, back},
			expectedStack:  []string{"main", "foo", "bar", "baz"},
			expectedCursor: 0,
		},
		{
			name: "single forward",
			stack: &Stack{
				Stack:  []string{"main", "foo", "bar", "baz"},
				Cursor: 2,
			},
			ops:            []wrappedOp{forward},
			expectedStack:  []string{"main", "foo", "bar", "baz"},
			expectedCursor: 3,
		},
		{
			name: "multiple forward",
			stack: &Stack{
				Stack:  []string{"main", "foo", "bar", "baz"},
				Cursor: 0,
			},
			ops:            []wrappedOp{forward, forward, forward},
			expectedStack:  []string{"main", "foo", "bar", "baz"},
			expectedCursor: 3,
		},
		{
			name: "front to back noop",
			stack: &Stack{
				Stack:  []string{"a", "b", "c"},
				Cursor: 0,
			},
			ops:            []wrappedOp{forward, forward, back, back},
			expectedStack:  []string{"a", "b", "c"},
			expectedCursor: 0,
		},
		{
			name: "pushing midstack",
			stack: &Stack{
				Stack:  []string{"a", "b", "c", "d"},
				Cursor: 2,
			},
			ops:            []wrappedOp{push("d")},
			expectedStack:  []string{"a", "b", "c", "d"},
			expectedCursor: 3,
		},
		{
			name: "pushing dupe at mid stack",
			stack: &Stack{
				Stack:  []string{"a", "b", "c"},
				Cursor: 1,
			},
			ops:            []wrappedOp{push("a")},
			expectedStack:  []string{"b", "a"},
			expectedCursor: 1,
		},
		{
			name: "pushing dupe at back of stack",
			stack: &Stack{
				Stack:  []string{"a", "b", "c"},
				Cursor: 0,
			},
			expectedStack:  []string{"a"},
			ops:            []wrappedOp{push("a")},
			expectedCursor: 0,
		},
		{
			name: "pushing dupe at front of stack",
			stack: &Stack{
				Stack:  []string{"a", "b", "c"},
				Cursor: 2,
			},
			expectedStack:  []string{"b", "c", "a"},
			ops:            []wrappedOp{push("a")},
			expectedCursor: 2,
		},
		{
			name: "empty push",
			stack: &Stack{
				Stack:  []string{},
				Cursor: 0,
			},
			expectedStack:  []string{"a"},
			ops:            []wrappedOp{push("a")},
			expectedCursor: 0,
		},
	}

	for _, c := range cases {
		for _, op := range c.ops {
			f := op(c.stack)
			_, err := f()
			if c.expectedErr != nil {
				require.Errorf(t, err, "%s: expected an error", c.name)
			} else {
				require.NoErrorf(t, err, "%s: unexpected error: %s", c.name, err)
			}
		}

		if fmt.Sprintf("%v", c.expectedStack) != fmt.Sprintf("%v", c.stack.Stack) {
			t.Fatalf("%s: expected Stack %v, got %v", c.name, c.expectedStack, c.stack.Stack)
		}

		if fmt.Sprintf("%v", c.expectedCursor) != fmt.Sprintf("%v", c.stack.Cursor) {
			t.Fatalf("%s: expected Cursor %v, got %v", c.name, c.expectedCursor, c.stack.Cursor)
		}
	}
}
