package sorter

type StringSorter []string

func (p StringSorter) Len() int {
	return len(p)
}

func (p StringSorter) Less(i, j int) bool {
	return p[j] < p[i]
}

func (p StringSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
