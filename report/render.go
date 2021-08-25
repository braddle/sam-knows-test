package report

import (
	"fmt"
	"time"
)

const bytesToMegabitDivider = 125000
const outputDateFormat = "2006-01-02"

type Reportable interface {
	GetStartDate() time.Time
	GetEndDate() time.Time
	GetAverageInBytes() float64
	GetMinimumInBytes() float64
	GetMaximumInBytes() float64
	GetMedianInBytes() float64
	HasUnderPerformance() bool
	GetUnderPerformanceStartDate() time.Time
	GetUnderPerformanceEndDate() time.Time
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
%s`,
	r.GetStartDate().Format(outputDateFormat),
	r.GetEndDate().Format(outputDateFormat),
	r.GetAverageInBytes() / bytesToMegabitDivider,
	r.GetMinimumInBytes() / bytesToMegabitDivider,
	r.GetMaximumInBytes() / bytesToMegabitDivider,
	r.GetMedianInBytes() / bytesToMegabitDivider,
	renderUnderPerforming(r),
		)
}

func renderUnderPerforming(r Reportable) string {
	if !r.HasUnderPerformance() {
		return ""
	}

	return fmt.Sprintf(`
Under-performing periods:

    * The period between %s and %s
      was under-performing.

`,
	r.GetUnderPerformanceStartDate().Format(outputDateFormat),
	r.GetUnderPerformanceEndDate().Format(outputDateFormat),
	)
}
