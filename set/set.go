package set

type Set[T comparable] struct {
	contents map[T]struct{}
}

func (s Set[T]) Add(item T) {
	s.contents[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s.contents, item)
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s.contents[item]
	return ok
}

func (s Set[T]) Items() []T {
	items := []T{}
	for k := range s.contents {
		items = append(items, k)
	}

	return items
}

func NewEmptySet[T comparable]() Set[T] {
	return Set[T]{contents: map[T]struct{}{}}
}

func NewSetFromIterable[T comparable](iterable []T) Set[T] {
	s := NewEmptySet[T]()
	for _, item := range iterable {
		s.Add(item)
	}

	return s
}

func NewSetFromString(s string) Set[string] {
	set := NewEmptySet[string]()
	for _, c := range s {
		set.Add(string(c))
	}
	return set
}

func Union[T comparable](sets ...Set[T]) Set[T] {
	union := NewEmptySet[T]()

	for _, i := range sets[0].Items() {
		inAllSets := true

		for _, s := range sets[1:] {
			if !s.Contains(i) {
				inAllSets = false
				break
			}
		}

		if inAllSets {
			union.Add(i)
		}
	}

	return union
}
