package main

type SSlice []string

//implementation of sort interface
func (t SSlice) Len() int {
	return len(t)
}

func (t SSlice) Less(i, j int) bool {
	return t[i] > t[j]
}

func (t SSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
