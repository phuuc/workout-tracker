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
			ApiHost: "127.0.0.1",
			ApiPort: 3000,
		}
	)
	f, err := os.CreateTemp(helpers.RootDir(), ".env")
	if err != nil {
		return
	}
	_, err = f.WriteString(`API_HOST="localhost"\nAPI_PORT=3000`)
	if err != nil {
		return
	}
	defer f.Close()

	s.Run("success", func() {
		cfg, err := config.Load(f.Name())

		s.NoError(err)
		s.Equal(server, cfg.Server)
	})
}
