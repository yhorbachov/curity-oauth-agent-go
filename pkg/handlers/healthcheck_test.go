package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/test"
)

func TestHealthcheck(t *testing.T) {
	test.WithTestServer(t, func(server *fiber.App) {
		t.Run("returns 200 OK", func(t *testing.T) {
			req := httptest.NewRequest("GET", "/healthcheck", nil)
			resp, _ := server.Test(req)

			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
			}
		})
	})
}
