package main

import (
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
	Interval{-1, 0.5},
	Interval{0., 1.},
	// Interval{0., 1.},
	Interval{0.75, 1.75},
	Interval{1.5, 2.},
}

var coverings1 = map[float64][]Interval{
	0.25: []Interval{
		Interval{-1., 0.5},
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
	TreeTest{intervals1, 4, coverings1},
}

func InsertIntervals(tree IntervalTreeInterface, intervals []Interval) {
	for _, i := range intervals {
		tree.Insert(i)
	}
}

func TestNaiveTree(t *testing.T) {
	tree := NewNaiveIntervalTree()
	testTree(t, tree)
}

func TestLLRBTree(t *testing.T) {
	tree := NewTree()
	testTree(t, tree)
}

func testTree(t *testing.T, tree IntervalTreeInterface) {
	for _, test := range treeTests {
		InsertIntervals(tree, intervals1)

		if tree.Len() != test.expectedLen {
			t.Errorf("Tree had Len() == %v, expected %v", tree.Len(), test.expectedLen)
		}

		for x, intervals := range test.coverings {
			coveringSlices := tree.Cover(x)
			if !unorderedSliceEqual(coveringSlices, intervals) {
				t.Errorf("For point %v, got covering slices %v, expected %v", x, coveringSlices, intervals)
			}
		}
	}
}
