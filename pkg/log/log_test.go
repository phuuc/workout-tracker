package log_test

import (
	"testing"

	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/stretchr/testify/suite"
)

type logSuite struct {
	suite.Suite

	*log.Log
}

func (s *logSuite) SetupSuite() {
	s.Log = log.NewLog()
}

func TestLogSuiteInit(t *testing.T) {
	suite.Run(t, new(logSuite))
}

func (s *logSuite) TestInfo() {
	var (
		message = "test"
	)
	s.Run("success", func() {
		s.Log.Info(message)
	})
}

func (s *logSuite) TestError() {
	var (
		message = "test%v"
	)
	s.Run("success", func() {
		s.Log.Error(message, 01)
	})
}
