package sorter

type IntSorter []int

func (p IntSorter) Len() int {
	return len(p)
}

func (p IntSorter) Less(i, j int) bool {
	return p[j] < p[i]
}

func (p IntSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
