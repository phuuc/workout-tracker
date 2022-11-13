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

func (l *logSuite) SetupSuite() {
	l.Log = log.NewLog()
}

func TestLogSuiteInit(t *testing.T) {
	suite.Run(t, new(logSuite))
}

func (s *logSuite) TestInfo() {
	var (
		message = "test"
	)
	s.Run("success", func() {
		_, err := s.Log.Info(message)

		s.NoError(err)
	})
}

func (s *logSuite) TestError() {
	var (
		message = "test"
	)
	s.Run("success", func() {
		_, err := s.Log.Error(message)

		s.NoError(err)
	})
}
