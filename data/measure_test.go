package data_test

import (
	"github.com/braddle/sam-knows-test/data"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type MeasureSuite struct {
	suite.Suite
}

func TestMeasureSuite(t *testing.T) {
	suite.Run(t, new(MeasureSuite))
}

var m data.Measurements
var startDate = data.JsonTime{time.Date(2020, 06, 01, 12, 0, 0, 0, time.UTC)}
var endDate = data.JsonTime{time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC)}

func (s *MeasureSuite) SetupTest() {
	m = data.Measurements{
		M: []data.Measure{
			{
				Time:   startDate,
				Metric: 10.5,
			},
			{
				Time: data.JsonTime{time.Date(2020, 06, 02, 12, 0,0,0, time.UTC)},
				Metric: 9.5,
			},
			{
				Time: data.JsonTime{time.Date(2020, 06, 03, 12, 0,0,0, time.UTC)},
				Metric: 8.2,
			},
			{
				Time: endDate,
				Metric: 11.99,
			},
		},
	}
}

func (s *MeasureSuite) TestGetStartDate() {
 	s.Equal(startDate.Time, m.GetStartDate())
}

func (s *MeasureSuite) TestGetEndDate() {
	s.Equal(endDate.Time, m.GetEndDate())
}

func (s *MeasureSuite) TestGetMinimumInMBPS() {
	s.Equal(8.2, m.GetMinimumInBytes())
}

func (s *MeasureSuite) TestGetMaximumInMBPS() {
	s.Equal(11.99, m.GetMaximumInBytes())
}

func (s *MeasureSuite) TestGetAverageInMBPS() {
	s.Equal(float64(10.0475), m.GetAverageInBytes())
}

func (s *MeasureSuite) TestGetMedianInMBPSOddNumberOfMetrics() {
	m := data.Measurements{
		M: []data.Measure{
			{
				Time:   data.JsonTime{time.Date(2020, 06, 01, 12, 0, 0, 0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time: data.JsonTime{time.Date(2020, 06, 02, 12, 0,0,0, time.UTC)},
				Metric: 9.5,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 03, 12, 0,0,0, time.UTC)},
				Metric: 8.2,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC)},
				Metric: 11.99,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 05, 12, 0, 0, 0, time.UTC)},
				Metric: 9.2,
			},
		},
	}

	s.Equal(float64(9.5), m.GetMedianInBytes())
}

func (s *MeasureSuite) TestGetMedianInMBPSEvenNumberOfMetrics() {
	s.Equal(float64(10), m.GetMedianInBytes())
}

func (s *MeasureSuite) TestWithoutUnderPerformance() {
	m := data.Measurements{
		M: []data.Measure{
			{
				Time:   data.JsonTime{time.Date(2020, 06, 01, 12, 0, 0, 0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time: data.JsonTime{time.Date(2020, 06, 02, 12, 0,0,0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 03, 12, 0,0,0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 05, 12, 0, 0, 0, time.UTC)},
				Metric: 10.5,
			},
		},
	}

	s.False(m.HasUnderPerformance())
}

func (s *MeasureSuite) TestWithUnderPerformance() {
	start := time.Date(2020, 06, 03, 12, 0, 0, 0, time.UTC)
	end := time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC)
	m := data.Measurements{
		M: []data.Measure{
			{
				Time:   data.JsonTime{time.Date(2020, 06, 01, 12, 0, 0, 0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time: data.JsonTime{time.Date(2020, 06, 02, 12, 0,0,0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time:   data.JsonTime{start},
				Metric: 1.5,
			},
			{
				Time:   data.JsonTime{end},
				Metric: 1.5,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 05, 12, 0, 0, 0, time.UTC)},
				Metric: 10.5,
			},
			{
				Time:   data.JsonTime{time.Date(2020, 06, 06, 12, 0, 0, 0, time.UTC)},
				Metric: 6.2,
			},
		},
	}

	s.True(m.HasUnderPerformance())
	s.Equal(start, m.GetUnderPerformanceStartDate())
	s.Equal(end, m.GetUnderPerformanceEndDate())
}



