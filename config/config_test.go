package config_test

import (
	"os"
	"testing"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/pkg/helpers"
	"github.com/stretchr/testify/suite"
)

type configSuite struct {
	suite.Suite

	*config.Config
}

func TestConfigSuiteInit(t *testing.T) {
	suite.Run(t, new(configSuite))
}

func (s *configSuite) TestNewConfig() {
	var (
		server = &config.Server{
			Host: "localhost",
			Port: 3000,
		}
	)
	f, err := os.CreateTemp(helpers.RootDir(), ".env")
	if err != nil {
		return
	}
	f.WriteString(`API_HOST="localhost"\nAPI_PORT=3000`)
	defer f.Close()

	s.Run("success", func() {
		cfg := config.NewConfig()

		s.Equal(server, cfg.Server)
	})
}
