package main

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// func (i *TreePair) Generate(rand *rand.Rand, size int) reflect.Value {
// 	mock := NewNaiveIntervalTree()
// 	real := NewTree()
// 	for i := 0; i < size; i++ {
// 		interval := Interval{rand.Float64(), rand.Float64()}
// 		mock.Insert(interval)
// 		real.Insert(interval)
// 	}
// 	return reflect.ValueOf(&TreePair{mock, real})
// }

func (i Interval) Generate(rand *rand.Rand, size int) reflect.Value {
	a := rand.Float64()
	b := rand.Float64()
	return reflect.ValueOf(Interval{math.Min(a, b), math.Max(a, b)})
}

func (s IntervalSlice) Generate(rand *rand.Rand, size int) reflect.Value {
	out := make([]Interval, size)
	for i := 0; i < size; i++ {

		a := rand.Float64()
		b := rand.Float64()
		out[i] = Interval{math.Min(a, b), math.Max(a, b)}
	}

	return reflect.ValueOf(out)
}

func TestCoverings(t *testing.T) {
	f := func(intervals IntervalSlice, x float64) bool {
		mock := NewNaiveIntervalTree()
		real := NewTree()
		for _, i := range intervals {
			mock.Insert(i)
			real.Insert(i)
		}
		expected := mock.Cover(x)
		actual := real.Cover(x)
		if !unorderedSliceEqual(expected, actual) {
			lOnly, rOnly := diff(expected, actual)
			t.Errorf("For x = %v, expected: %v, got: %v", x, expected, actual)
			t.Errorf("Expected, but didn't get: %v", lOnly)
			t.Errorf("Got, but didn't expect: %v", rOnly)
			return false
		} else {
			t.Logf("Successfuly got: %v", expected)
		}
		return true
	}
	config := &quick.Config{MaxCount: 5}
	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}
