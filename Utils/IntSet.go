package utils

type IntSet map[int]bool

func (s IntSet) Add(item int) {
	s[item] = true
}

func (s IntSet) Intersect(s2 map[int]bool) map[int]bool {
	s1 := s
	intersection := map[int]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter IntSet
	}
	for k, _ := range s1 {
		if s2[k] {
			intersection[k] = true
		}
	}
	return intersection
}
