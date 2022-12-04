package set

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s))
	for k := range s {
		items = append(items, k)
	}

	return items
}

func NewEmptySet[T comparable]() Set[T] {
	return map[T]struct{}{}
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

func Intersection[T comparable](sets ...Set[T]) Set[T] {
	intersection := NewEmptySet[T]()

	for _, i := range sets[0].Items() {
		inAllSets := true

		for _, s := range sets[1:] {
			if !s.Contains(i) {
				inAllSets = false
				break
			}
		}

		if inAllSets {
			intersection.Add(i)
		}
	}

	return intersection
}
