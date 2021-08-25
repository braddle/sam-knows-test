package report

import (
	"fmt"
	"time"
)

const bytesToMegabitDivider = 125000

type Reportable interface {
	GetStartDate() time.Time
	GetEndDate() time.Time
	GetAverageInBytes() float64
	GetMinimumInBytes() float64
	GetMaximumInBytes() float64
	GetMedianInBytes() float64
}

func Render(r Reportable) string {

	return fmt.Sprintf(
		`SamKnows Metric Analyser v1.0.0
===============================

Period checked:

    From: %s
    To:   %s

Statistics:

    Unit: Megabits per second

    Average: %.1f
    Min: %.2f
    Max: %.2f
    Median: %.2f
`,
	r.GetStartDate().Format("2006-01-02"),
	r.GetEndDate().Format("2006-01-02"),
	r.GetAverageInBytes() / bytesToMegabitDivider,
	r.GetMinimumInBytes() / bytesToMegabitDivider,
	r.GetMaximumInBytes() / bytesToMegabitDivider,
	r.GetMedianInBytes() / bytesToMegabitDivider,
		)
}
