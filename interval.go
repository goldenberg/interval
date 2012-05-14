package main

import (
	"fmt"
)

type Interval struct {
	Start, End float64
}

func (i *Interval) Contains(x float64) bool {
	return (x >= i.Start) && (x <= i.End)
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%v, %v]", i.Start, i.End)
}

func Intersect(x, y Interval) Interval {
	// Assume WLOG, x.Start < y.Start
	if x.Start > y.Start {
		x, y = y, x
	}

	switch {
	case x.End < y.Start:
		// x and y do not intersect.
		return Interval{}
	case x.End < y.Start && x.End < y.End:
		// x and y partially overlap.
		return Interval{x.End, y.Start}
	case x.Start < y.Start && y.End < x.End:
		// y is a subinterval of x
		return y
	}
	return Interval{}
}

type IntervalSlice []Interval

func (s IntervalSlice) Len() int { return len(s) }

// Less returns whether the element with index i should sort
// before the element with index j.
func (s IntervalSlice) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

// Swap swaps the elements with indexes i and j.
func (s IntervalSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
