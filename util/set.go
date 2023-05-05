package util

import "strings"

type StringSet map[string]struct{}

func NewStringSet(ms ...string) StringSet {
	set := map[string]struct{}{}
	for _, m := range ms {
		set[m] = struct{}{}
	}
	return set
}

func (s StringSet) Add(ms ...string) StringSet {
	if s == nil {
		s = NewStringSet(ms...)
		return s
	}
	for _, m := range ms {
		s[m] = struct{}{}
	}
	return s
}

func (s StringSet) Del(member string) bool {
	_, ok := s[member]
	delete(s, member)
	return !ok
}

func (s StringSet) Diff(o StringSet) StringSet {
	for k := range o {
		delete(s, k)
	}
	return s
}

func (s StringSet) DiffWith(ms ...string) StringSet {
	other := NewStringSet(ms...)
	return s.Diff(other)
}

func (s StringSet) Union(o StringSet) StringSet {
	for k := range o {
		s[k] = struct{}{}
	}
	return s
}

func (s StringSet) UnionWith(ms ...string) StringSet {
	other := NewStringSet(ms...)
	return s.Union(other)
}

func (s StringSet) String() string {
	if s == nil {
		return "StringSet{}"
	}
	var b strings.Builder
	first := true
	b.WriteString("StringSet{")
	for k := range s {
		if first {
			first = false
		} else {
			b.WriteString(",")
		}
		b.WriteString(" " + k)
	}
	b.WriteString("}")
	return b.String()
}
