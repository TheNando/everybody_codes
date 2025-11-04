package util

type Cycle[T any] struct {
	items    []T
	position int
}

func (s *Cycle[T]) leftIndex(steps int) int {
	pos := (s.position - steps) % len(s.items)
	if pos < 0 {
		pos += len(s.items)
	}
	return pos
}

func (s *Cycle[T]) rightIndex(steps int) int {
	pos := (s.position + steps) % len(s.items)
	if pos > len(s.items)-1 {
		pos -= len(s.items) - 1
	}
	return pos
}

// Move the position left by the given number of steps.
func (s *Cycle[T]) Left(steps int) {
	s.position = s.leftIndex(steps)
}

// Move the position right by the given number of steps.
func (s *Cycle[T]) Right(steps int) {
	s.position = s.rightIndex(steps)
}

func (s *Cycle[T]) Swap(dir string, steps int) {
	swapIndex := 0

	switch dir {
	case "L":
		swapIndex = s.leftIndex(steps)
	case "R":
		swapIndex = s.rightIndex(steps)
	}

	s.items[0], s.items[swapIndex] = s.items[swapIndex], s.items[0]
}

func (s *Cycle[T]) Item() T {
	return s.items[s.position]
}

func NewCycle[T any](items []T) *Cycle[T] {
	return &Cycle[T]{
		items:    items,
		position: 0,
	}
}
