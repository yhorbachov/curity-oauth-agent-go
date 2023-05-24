package test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/conf"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/server"
)

func WithTestServer(t *testing.T, callback func(server *fiber.App)) {
	t.Helper()

	config := conf.Config{
		Port:           0,
		EndpointsPefix: "",
	}

	server := server.NewServer(config)

	callback(server)

	if err := server.Shutdown(); err != nil {
		t.Fatalf("failed to shutdown server: %v", err)
	}

	server = nil
}
