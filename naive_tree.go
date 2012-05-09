package main

type NaiveIntervalTree map[Interval]bool

func NewNaiveIntervalTree() NaiveIntervalTree {
	return NaiveIntervalTree(make(map[Interval]bool, 2))
}

func (t NaiveIntervalTree) Insert(i Interval) {
	t[i] = true
}

func (t NaiveIntervalTree) Remove(i Interval) {
	delete(t, i)
}

func (t NaiveIntervalTree) Cover(x float64) (out []Interval) {
	out = make([]Interval, 0)
	for ti, _ := range t {
		if ti.In(x) {
			out = append(out, ti)
		}
	}
	return
}

func (t NaiveIntervalTree) Len() int {
	return len(t)
}
