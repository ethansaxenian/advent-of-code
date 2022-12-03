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

func NewSetFromIterable[T comparable](iterable []T) Set[T] {
	s := Set[T]{contents: map[T]struct{}{}}
	for _, item := range iterable {
		s.Add(item)
	}

	return s
}

func NewSetFromString(s string) Set[string] {
	set := Set[string]{contents: map[string]struct{}{}}
	for _, c := range s {
		set.Add(string(c))
	}
	return set
}
