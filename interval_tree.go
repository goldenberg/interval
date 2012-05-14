package main

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
	"math"
)

type IntervalItem struct {
	*Interval
	maxEnd float64
}

func max(a ...float64) (m float64) {
	m = math.Inf(-1)
	for _, x := range a {
		m = math.Max(m, x)
	}
	return
}

func (i *IntervalItem) Notify(n *llrb.Node) {
	fmt.Println("notified of a change!")
	leftMax := math.Inf(-1)
	rightMax := math.Inf(-1)
	if n.Left != nil {
		leftMax = n.Left.Item.(IntervalItem).maxEnd
	}
	if n.Right != nil {
		rightMax = n.Right.Item.(IntervalItem).maxEnd
	}

	i.maxEnd = math.Max(leftMax, rightMax)
}

func lessFunc(x, y interface{}) bool {
	return x.(IntervalItem).Start < y.(IntervalItem).Start
}

type IntervalTreeInterface interface {
	Insert(i Interval)
	Remove(i Interval)
	Cover(x float64) []Interval
	Len() int
}

type Tree struct {
	*llrb.Tree
}

func NewTree() *Tree {
	return &Tree{llrb.New(lessFunc)}
}

func (t *Tree) Insert(i Interval) {
	item := IntervalItem{&i, i.End}
	t.InsertNoReplace(item)
	DoOnNodes(t.Root(), updateMaxEnd)
}

func updateMaxEnd(n *llrb.Node) {
	item := n.Item.(IntervalItem)
	item.Notify(n)
}

func DoOnNodes(n *llrb.Node, f func(n *llrb.Node)) {
	f(n)
	if n.Left != nil {
		DoOnNodes(n.Left, updateMaxEnd)
	}

	if n.Right != nil {
		DoOnNodes(n.Right, updateMaxEnd)
	}
}

func (t *Tree) Remove(i Interval) {
	t.Delete(i)
}

func (t *Tree) Cover(x float64) (out []Interval) {
	c := make(chan Interval, 100)
	out = make([]Interval, 0)
	go func() {
		t.search(t.Root(), x, c)
		close(c)
	}()
	for x := range c {
		fmt.Println("appending", x)
		out = append(out, x)
	}
	fmt.Println("final out", out)
	return out
}

func (t *Tree) search(n *llrb.Node, x float64, out chan Interval) {
	fmt.Println("search(", n.Item, x, out, ")")
	interval := n.Item.(IntervalItem)

	fmt.Println("interval:", interval, "maxEnd:", interval.maxEnd)

	// If x is past the end, then there won't be any matches.
	if x > interval.maxEnd {
		fmt.Println("past the end")
		return
	}

	// Search left children
	if n.Left != nil {
		fmt.Println("searching left")
		t.search(n.Left, x, out)
	}

	// Check this interval
	if interval.Contains(x) {
		fmt.Println("got one!", interval.Interval)
		out <- *interval.Interval
		fmt.Println("out:", out)
	}

	// If x is to the left of the start, then it can't be in any child on the
	// right.
	if x < interval.Start {
		fmt.Println("we're to the left of the start")
		return
	}

	// Otherwise search the right children.
	if n.Right != nil {
		fmt.Println("searching right")
		t.search(n.Right, x, out)
	}
}

// func (t *Tree) Len() int {
// 	return t.Len()
// }
