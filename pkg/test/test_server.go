package test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/conf"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/server"
)

func Run(t *testing.T, name string, callback func(s *fiber.App)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		config := conf.Config{
			Port:           0,
			EndpointsPefix: "",
		}

		s := server.NewServer(config)

		callback(s)

		if err := s.Shutdown(); err != nil {
			t.Fatalf("failed to shutdown server: %v", err)
		}

		s = nil
	})
}
