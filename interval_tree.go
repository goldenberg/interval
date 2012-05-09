package main

import (
	"github.com/petar/GoLLRB/llrb"
	"math"
)

type IntervalNode struct {
	*Interval
	maxEnd float64
}

func (i *IntervalNode) Notify(n *llrb.Node) {
	i.maxEnd = math.Max(i.maxEnd, math.Max(n.Left.Item.(IntervalNode).maxEnd,
		n.Right.Item.(IntervalNode).maxEnd))
}

func lessFunc(x, y interface{}) bool {
	return x.(IntervalNode).Start < y.(IntervalNode).Start
}

type IntervalTreeInterface interface {
	Insert(i Interval)
	Remove(i Interval)
	Cover(x float64) []Interval
	Len() int
}

type Tree llrb.Tree
