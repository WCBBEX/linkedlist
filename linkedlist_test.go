package linkedlist

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	l := New[int]()
	if l.Len() != 0 {
		t.Errorf("New list length should be 0, got %d", l.Len())
	}
	if l.head != nil || l.tail != nil {
		t.Error("New list head and tail should be nil")
	}
}

func TestNewWithSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"empty slice", []int{}},
		{"single element", []int{1}},
		{"multiple elements", []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewWithSlice(tt.input)
			if l.Len() != len(tt.input) {
				t.Errorf("Expected length %d, got %d", len(tt.input), l.Len())
			}

			// Verify elements
			for i, v := range tt.input {
				val, err := l.Get(i)
				if err != nil {
					t.Errorf("Unexpected error getting element %d: %v", i, err)
				}
				if *val != v {
					t.Errorf("Expected %d at index %d, got %d", v, i, *val)
				}
			}
		})
	}
}

func TestPushBack(t *testing.T) {
	l := New[int]()
	values := []int{1, 2, 3}

	for i, v := range values {
		l.PushBack(&v)
		if l.Len() != i+1 {
			t.Errorf("Expected length %d after PushBack, got %d", i+1, l.Len())
		}
	}

	// Verify tail value
	if *l.tail.value != values[len(values)-1] {
		t.Errorf("Expected tail value %d, got %d", values[len(values)-1], *l.tail.value)
	}
}

func TestPushFront(t *testing.T) {
	l := New[int]()
	values := []int{1, 2, 3}

	for i, v := range values {
		l.PushFront(&v)
		if l.Len() != i+1 {
			t.Errorf("Expected length %d after PushFront, got %d", i+1, l.Len())
		}
	}

	// Verify head value
	if *l.head.value != values[len(values)-1] {
		t.Errorf("Expected head value %d, got %d", values[len(values)-1], *l.head.value)
	}
}

func TestPopBack(t *testing.T) {
	l := New[int]()
	values := []int{1, 2, 3}
	for _, v := range values {
		l.PushBack(&v)
	}

	for i := len(values) - 1; i >= 0; i-- {
		val, err := l.PopBack()
		if err != nil {
			t.Errorf("Unexpected error popping back: %v", err)
		}
		if *val != values[i] {
			t.Errorf("Expected %d when popping back, got %d", values[i], *val)
		}
		if l.Len() != i {
			t.Errorf("Expected length %d after PopBack, got %d", i, l.Len())
		}
	}

	// Test empty list
	_, err := l.PopBack()
	if !errors.Is(err, EmptyListErr) {
		t.Errorf("Expected ErrEmptyList when popping from empty list, got %v", err)
	}
}

func TestPopFront(t *testing.T) {
	l := New[int]()
	values := []int{1, 2, 3}
	for _, v := range values {
		l.PushBack(&v)
	}

	for i := 0; i < len(values); i++ {
		val, err := l.PopFront()
		if err != nil {
			t.Errorf("Unexpected error popping front: %v", err)
		}
		if *val != values[i] {
			t.Errorf("Expected %d when popping front, got %d", values[i], *val)
		}
		if l.Len() != len(values)-i-1 {
			t.Errorf("Expected length %d after PopFront, got %d", len(values)-i-1, l.Len())
		}
	}

	// Test empty list
	_, err := l.PopFront()
	if !errors.Is(err, EmptyListErr) {
		t.Errorf("Expected ErrEmptyList when popping from empty list, got %v", err)
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		insertAt int
		value    int
		want     []int
		wantErr  error
	}{
		{"empty list", []int{}, 0, 1, []int{1}, nil},
		{"insert at start", []int{2, 3}, 0, 1, []int{1, 2, 3}, nil},
		{"insert at end", []int{1, 2}, 2, 3, []int{1, 2, 3}, nil},
		{"insert middle", []int{1, 3}, 1, 2, []int{1, 2, 3}, nil},
		{"invalid index", []int{1}, 2, 2, []int{1}, IndexOutOfRangeErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewWithSlice(tt.initial)
			err := l.Insert(tt.insertAt, &tt.value)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}

			if err != nil {
				return
			}

			if l.Len() != len(tt.want) {
				t.Errorf("Expected length %d, got %d", len(tt.want), l.Len())
			}

			for i, v := range tt.want {
				val, err := l.Get(i)
				if err != nil {
					t.Errorf("Unexpected error getting element %d: %v", i, err)
				}
				if *val != v {
					t.Errorf("Expected %d at index %d, got %d", v, i, *val)
				}
			}
		})
	}
}

func TestGet(t *testing.T) {
	l := NewWithSlice([]int{1, 2, 3})

	tests := []struct {
		name    string
		index   int
		want    int
		wantErr error
	}{
		{"first element", 0, 1, nil},
		{"middle element", 1, 2, nil},
		{"last element", 2, 3, nil},
		{"negative index", -1, 0, IndexOutOfRangeErr},
		{"index out of bounds", 3, 0, IndexOutOfRangeErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := l.Get(tt.index)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}
			if err == nil && *val != tt.want {
				t.Errorf("Expected %d at index %d, got %d", tt.want, tt.index, *val)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		removeAt  int
		want      []int
		wantValue int
		wantErr   error
	}{
		{"remove first", []int{1, 2, 3}, 0, []int{2, 3}, 1, nil},
		{"remove middle", []int{1, 2, 3}, 1, []int{1, 3}, 2, nil},
		{"remove last", []int{1, 2, 3}, 2, []int{1, 2}, 3, nil},
		{"remove from empty", []int{}, 0, []int{}, 0, EmptyListErr},
		{"invalid index", []int{1}, 1, []int{1}, 0, IndexOutOfRangeErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewWithSlice(tt.initial)
			val, err := l.Remove(tt.removeAt)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}

			if err != nil {
				return
			}

			if *val != tt.wantValue {
				t.Errorf("Expected removed value %d, got %d", tt.wantValue, *val)
			}

			if l.Len() != len(tt.want) {
				t.Errorf("Expected length %d, got %d", len(tt.want), l.Len())
			}

			for i, v := range tt.want {
				got, err := l.Get(i)
				if err != nil {
					t.Errorf("Unexpected error getting element %d: %v", i, err)
				}
				if *got != v {
					t.Errorf("Expected %d at index %d, got %d", v, i, *got)
				}
			}
		})
	}
}

func TestIter(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	l := NewWithSlice(values)

	collected := make([]int, 0)
	for v := range l.Iter() {
		collected = append(collected, *v)
	}

	if len(collected) != len(values) {
		t.Errorf("Expected %d elements, got %d", len(values), len(collected))
	}

	for i, v := range values {
		if collected[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, collected[i])
		}
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("single element list", func(t *testing.T) {
		l := New[int]()
		v := 42
		l.PushBack(&v)

		if l.head != l.tail {
			t.Error("Head and tail should be the same for single-element list")
		}

		_, err := l.PopFront()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if l.Len() != 0 || l.head != nil || l.tail != nil {
			t.Error("List should be empty after popping last element")
		}
	})

	t.Run("insert into empty list at non-zero", func(t *testing.T) {
		l := New[int]()
		v := 42
		err := l.Insert(1, &v)
		if !errors.Is(err, IndexOutOfRangeErr) {
			t.Errorf("Expected IndexOutOfRangeErr, got %v", err)
		}
	})
}
