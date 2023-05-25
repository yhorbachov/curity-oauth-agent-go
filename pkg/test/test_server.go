package test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/conf"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/server"
)

func Run(t *testing.T, name string, callback func(t *testing.T, s *fiber.App)) {
	// t.Helper()
	t.Run(name, func(t *testing.T) {
		config := conf.Config{
			Port:              0,
			EndpointsPefix:    "",
			TrustedWebOrigins: []string{"http://valid-origin.com"},
		}

		s := server.NewServer(config)

		callback(t, s)

		if err := s.Shutdown(); err != nil {
			t.Fatalf("failed to shutdown server: %v", err)
		}

		s = nil
	})
}
