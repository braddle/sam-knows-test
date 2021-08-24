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
var startDate = time.Date(2020, 06, 01, 12, 0, 0, 0, time.UTC)
var endDate = time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC)

func (s *MeasureSuite) SetupTest() {
	m = data.Measurements{
		M: []data.Measure{
			{
				Time:   startDate,
				Metric: 10.5,
			},
			{
				Time: time.Date(2020, 06, 02, 12, 0,0,0, time.UTC),
				Metric: 9.5,
			},
			{
				Time: time.Date(2020, 06, 03, 12, 0,0,0, time.UTC),
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
 	s.Equal(startDate, m.GetStartDate())
}

func (s *MeasureSuite) TestGetEndDate() {
	s.Equal(endDate, m.GetEndDate())
}

func (s *MeasureSuite) TestGetMinimumInMBPS() {
	s.Equal(8.2, m.GetMinimumInMBPS())
}

func (s *MeasureSuite) TestGetMaximumInMBPS() {
	s.Equal(11.99, m.GetMaximumInMBPS())
}

func (s *MeasureSuite) TestGetAverageInMBPS() {
	s.Equal(float64(10.0475), m.GetAverageInMBPS())
}

func (s *MeasureSuite) TestGetMedianInMBPSOddNumberOfMetrics() {
	m := data.Measurements{
		M: []data.Measure{
			{
				Time:   time.Date(2020, 06, 01, 12, 0, 0, 0, time.UTC),
				Metric: 10.5,
			},
			{
				Time: time.Date(2020, 06, 02, 12, 0,0,0, time.UTC),
				Metric: 9.5,
			},
			{
				Time:   time.Date(2020, 06, 03, 12, 0,0,0, time.UTC),
				Metric: 8.2,
			},
			{
				Time:   time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC),
				Metric: 11.99,
			},
			{
				Time:   time.Date(2020, 06, 05, 12, 0, 0, 0, time.UTC),
				Metric: 9.2,
			},
		},
	}

	s.Equal(float64(9.5), m.GetMedianInMBPS())
}

func (s *MeasureSuite) TestGetMedianInMBPSEvenNumberOfMetrics() {

	s.Equal(float64(10), m.GetMedianInMBPS())
}

