package main

type StringIntMap struct {
	m map[string]int
}

func (s *StringIntMap) Add(key string, value int) {
	s.m[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.m, key)
}

func (s StringIntMap) Copy() map[string]int {
	return s.m
}

func (s *StringIntMap) Exists(key string) bool {
	if _, ok := s.m[key]; !ok {
		return false
	} else {
		return true
	}
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.m[key]
	if !ok {
		return 0, false
	} else {
		return v, true
	}
}
