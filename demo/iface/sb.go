package iface

type Sb struct {
	x int
}

func (sb *Sb) Get() int {
	return sb.x
}

func (sb *Sb) Put(v int) {
	sb.x = v
}
