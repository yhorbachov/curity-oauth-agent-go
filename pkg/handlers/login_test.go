package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/test"
)

func TestLoginStart(t *testing.T) {
	test.Run(t, "Returns Unauthorized on invalid origin", func(s *fiber.App) {
		req := httptest.NewRequest("POST", "/login/start", nil)
		req.Header.Set("Origin", "http://invalid-origin.com")

		resp, _ := s.Test(req)

		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	})
}
