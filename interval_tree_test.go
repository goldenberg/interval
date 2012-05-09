package main

import (
	"sort"
	"testing"
)

type TreeTest struct {
	intervals   []Interval
	expectedLen int
	coverings   map[float64][]Interval
}

type Covering struct {
	x         float64
	intervals []Interval
}

var intervals1 = []Interval{
	Interval{0., 1.},
	Interval{0., 1.},
	Interval{0.75, 1.75},
	Interval{1.5, 2.},
}

var coverings1 = map[float64][]Interval{
	0.25: []Interval{
		Interval{0., 1.},
	},
	0.8: []Interval{
		Interval{0., 1.},
		Interval{0.75, 1.75},
	},
	1.1: []Interval{
		Interval{0.75, 1.75},
	},
	3: []Interval{},
}

var treeTests = []TreeTest{
	TreeTest{intervals1, 3, coverings1},
}

func InsertIntervals(tree IntervalTreeInterface, intervals []Interval) {
	for _, i := range intervals {
		tree.Insert(i)
	}
}

func sliceEqual(x, y []Interval) bool {
	if len(x) != len(y) {
		return false
	}
	for i, _ := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func unorderedSliceEqual(x, y []Interval) bool {
	orderedX := IntervalSlice(x)
	orderedY := IntervalSlice(y)
	sort.Sort(orderedX)
	sort.Sort(orderedY)

	return sliceEqual(orderedX, orderedY)
}

func TestLen(t *testing.T) {
	for _, test := range treeTests {
		tree := NewNaiveIntervalTree()
		InsertIntervals(tree, intervals1)

		if tree.Len() != test.expectedLen {
			t.Errorf("Tree had Len() == %v, expected %v", tree.Len(), test.expectedLen)
		}

		for x, intervals := range test.coverings {
			coveringSlices := tree.Cover(x)
			if !unorderedSliceEqual(coveringSlices, intervals) {
				t.Errorf("Got covering slices %v, expected %v", coveringSlices, intervals)
			}
		}
	}
}
