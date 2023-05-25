package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/test"
)

func TestHealthcheck(t *testing.T) {
	test.Run(t, "returns 200 OK", func(t *testing.T, s *fiber.App) {
		req := httptest.NewRequest("GET", "/healthcheck", nil)
		resp, _ := s.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
