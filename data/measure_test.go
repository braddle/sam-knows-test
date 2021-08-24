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

func (s *MeasureSuite) TestGetStartDate() {
	exp := time.Date(2020, 06, 01, 12, 0, 0, 0, time.UTC)
	m := data.Measurements{
 		M: []data.Measure{
 			{
 				Time:   exp,
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
				Time: time.Date(2020, 06, 04, 12, 0,0,0, time.UTC),
				Metric: 11.99,
			},
		},
 	}

 	s.Equal(exp, m.GetStartDate())
}

func (s *MeasureSuite) TestGetEndDate() {
	exp := time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC)
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
				Time: time.Date(2020, 06, 03, 12, 0,0,0, time.UTC),
				Metric: 8.2,
			},
			{
				Time:   exp,
				Metric: 11.99,
			},
		},
	}

	s.Equal(exp, m.GetEndDate())
}

func (s *MeasureSuite) TestGetMinimumInMBPS() {
	exp := 8.2
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
				Metric: exp,
			},
			{
				Time:   time.Date(2020, 06, 04, 12, 0, 0, 0, time.UTC),
				Metric: 11.99,
			},
		},
	}

	s.Equal(exp, m.GetMinimumInMBPS())
}

func (s *MeasureSuite) TestGetMaximumInMBPS() {
	exp := 11.99
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
				Metric: exp,
			},
		},
	}

	s.Equal(exp, m.GetMaximumInMBPS())
}

func (s *MeasureSuite) TestGetAverageInMBPS() {
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
		},
	}

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
		},
	}

	s.Equal(float64(10), m.GetMedianInMBPS())
}

