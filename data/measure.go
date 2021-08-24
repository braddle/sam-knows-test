package data

import (
	"math"
	"sort"
	"time"
)

type Measurements struct {
	M []Measure
}

func (m Measurements) GetStartDate() time.Time {
	var s time.Time
	var nilTime time.Time

	for _, d := range m.M {
		if s == nilTime || d.Time.Before(s) {
			s = d.Time
		}
	}

	return s
}

func (m Measurements) GetEndDate() time.Time {
	var e time.Time
	var nilTime time.Time

	for _, d := range m.M {
		if e == nilTime || d.Time.After(e) {
			e = d.Time
		}
	}

	return e
}

func (m Measurements) GetMinimumInMBPS() float64 {
	var f float64

	for _, d := range m.M {
		if f == 0 || d.Metric < f {
			f = d.Metric
		}
	}

	return f
}

func (m Measurements) GetMaximumInMBPS() float64 {
	var f float64

	for _, d := range m.M {
		if f == 0 || d.Metric > f {
			f = d.Metric
		}
	}

	return f
}

func (m Measurements) GetAverageInMBPS() float64 {
	var a float64

	for _, d := range m.M {
		a += d.Metric
	}

	return a / float64(len(m.M))
}

func (m Measurements) GetMedianInMBPS() interface{} {
	v := []float64{}

	for _, d := range m.M {
		v = append(v, d.Metric)
	}
	sort.Float64s(v)

	len := len(v)
	mid := int(math.Ceil(float64(len))) / 2
	var median float64

	if len%2 == 0 {
		upper := v[mid]
		lower := v[mid - 1]

		diff := upper - lower
		median = lower + (diff / 2)
	} else {
		median = v[mid]
	}

	return median
}

type Measure struct {
	Time   time.Time
	Metric float64
}