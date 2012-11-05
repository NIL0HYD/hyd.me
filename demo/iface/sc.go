package iface

type Sc struct {
	c int
}

func (sc Sc) Get() int {
	return sc.c
}

func (sc Sc) Put(v int) {
	sc.c = v
}
