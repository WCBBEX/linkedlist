package linkedlist

import (
	"errors"
	"fmt"
)

var (
	EmptyListErr       = errors.New("empty list")
	IndexOutOfRangeErr = errors.New("index out of range")
)

func checkIndex(index, len int) error {
	if index < 0 || index >= len {
		return fmt.Errorf("%w, index: %d length: %d", IndexOutOfRangeErr, index, len)
	}
	return nil
}

func checkEmpty(len int) error {
	if len == 0 {
		return fmt.Errorf("%w", EmptyListErr)
	}
	return nil
}
