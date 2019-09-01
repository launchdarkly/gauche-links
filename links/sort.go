package links

import "strings"

type LinkSliceByPath []Link

func (s LinkSliceByPath) Len() int {
	return len(s)
}

func (s LinkSliceByPath) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s LinkSliceByPath) Less(i, j int) bool {
	return strings.Compare(s[i].LiteralPrefix(), s[j].LiteralPrefix()) < 0 ||
		strings.Compare(s[i].PathString(), s[j].PathString()) < 0
}

type LinkSliceByPathLength []Link

func (s LinkSliceByPathLength) Len() int {
	return len(s)
}
func (s LinkSliceByPathLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s LinkSliceByPathLength) Less(i, j int) bool {
	return s[i].Len() < s[j].Len()
}
