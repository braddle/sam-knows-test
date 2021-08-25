package report_test

import (
	"github.com/braddle/sam-knows-test/report"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type ReportRenderSuite struct {
	suite.Suite
}

func TestReportRenderSuite(t *testing.T) {
	suite.Run(t, new(ReportRenderSuite))
}

func (s *ReportRenderSuite) TestRenderWithoutUnderPerformingPeriods() {
	r := new(MockReportable)
	r.On("GetStartDate").Return(time.Date(2018, 01, 29, 12, 0,0,0, time.UTC))
	r.On("GetEndDate").Return(time.Date(2018, 02, 27, 12, 0,0,0, time.UTC))
	r.On("GetAverageInBytes").Return(float64(12837500))
	r.On("GetMinimumInBytes").Return(float64(12656250))
	r.On("GetMaximumInBytes").Return(float64(13010000))
	r.On("GetMedianInBytes").Return(float64(12866250))
	r.On("HasUnderPerformance").Return(false)

	act := report.Render(r)
	f, _ := os.Open("../outputs/1.output")
	exp, _ := ioutil.ReadAll(f)

	s.Equal(string(exp), act)
}

func (s *ReportRenderSuite) TestRenderWithUnderPerformingPeriods() {
	r := new(MockReportable)
	r.On("GetStartDate").Return(time.Date(2018, 01, 29, 12, 0,0,0, time.UTC))
	r.On("GetEndDate").Return(time.Date(2018, 02, 27, 12, 0,0,0, time.UTC))
	r.On("GetAverageInBytes").Return(float64(11937500))
	r.On("GetMinimumInBytes").Return(float64(3453750))
	r.On("GetMaximumInBytes").Return(float64(13010000))
	r.On("GetMedianInBytes").Return(float64(12863750))
	r.On("HasUnderPerformance").Return(true)
	r.On("GetUnderPerformanceStartDate").Return(time.Date(2018, 02, 05, 12, 0,0,0, time.UTC))
	r.On("GetUnderPerformanceEndDate").Return(time.Date(2018, 02, 07, 12, 0,0,0, time.UTC))


	act := report.Render(r)
	f, _ := os.Open("../outputs/2.output")
	exp, _ := ioutil.ReadAll(f)

	s.Equal(string(exp), act)
}

type MockReportable struct {
	mock.Mock
}

func (m *MockReportable) GetUnderPerformanceEndDate() time.Time {
	args := m.Called()

	return args.Get(0).(time.Time)
}

func (m *MockReportable) GetUnderPerformanceStartDate() time.Time {
	args := m.Called()

	return args.Get(0).(time.Time)
}

func (m *MockReportable) HasUnderPerformance() bool {
	args := m.Called()

	return args.Bool(0)
}

func (m *MockReportable) GetStartDate() time.Time {
	args := m.Called()

	return args.Get(0).(time.Time)
}

func (m *MockReportable) GetEndDate() time.Time {
	args := m.Called()

	return args.Get(0).(time.Time)
}

func (m *MockReportable) GetAverageInBytes() float64 {
	args := m.Called()

	return args.Get(0).(float64)
}

func (m *MockReportable) GetMinimumInBytes() float64 {
	args := m.Called()

	return args.Get(0).(float64)
}

func (m *MockReportable) GetMaximumInBytes() float64 {
	args := m.Called()

	return args.Get(0).(float64)
}

func (m *MockReportable) GetMedianInBytes() float64 {
	args := m.Called()

	return args.Get(0).(float64)
}




