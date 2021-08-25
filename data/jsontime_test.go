package data_test

import (
	"github.com/braddle/sam-knows-test/data"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type JsonTimeSuite struct {
	suite.Suite
}

func TestJsonTimeSuite(t *testing.T) {
	suite.Run(t, new(JsonTimeSuite))
}

func (s *JsonTimeSuite) TestJsonTimeUnmarshalValidFormat() {
	exp := time.Date(2020, 06, 05, 0, 0, 0, 0, time.UTC)

	jt := data.JsonTime{}
	err := jt.UnmarshalJSON([]byte(`"2020-06-05"`))

	s.NoError(err)
	s.Equal(exp, jt.Time)
}

func (s *JsonTimeSuite) TestJsonTimeUnmarshalInalidFormat() {
	jt := data.JsonTime{}
	err := jt.UnmarshalJSON([]byte(`"20-13-05"`))

	s.Error(err)
}
