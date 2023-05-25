package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/test"
)

func TestLoginStart(t *testing.T) {
	test.Run(t, "Returns Unauthorized on invalid origin", func(t *testing.T, s *fiber.App) {
		req := httptest.NewRequest("POST", "/login/start", nil)
		req.Header = newHeaders(t)
		req.Header.Set("origin", "http://invalid-origin.com")

		resp, _ := s.Test(req)
		assert.Equal(t, http.StatusForbidden, resp.StatusCode, "Status code should be 403")
	})

	test.Run(t, "Returns BadRequest on missing redirectUri", func(t *testing.T, s *fiber.App) {
		req := httptest.NewRequest("POST", "/login/start", bytes.NewReader([]byte("{}")))
		req.Header = newHeaders(t)

		resp, _ := s.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Status code should be 400")
	})

	test.Run(t, "Returns Ok on valid request", func(t *testing.T, s *fiber.App) {
		req := httptest.NewRequest("POST", "/login/start", bytes.NewReader([]byte(`{"redirectUri":"http://valid-redirect-uri.com"}`)))
		req.Header = newHeaders(t)

		resp, _ := s.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200")
	})
}

func newHeaders(t *testing.T) map[string][]string {
	t.Helper()

	return map[string][]string{
		"origin":       {"http://valid-origin.com"},
		"accept":       {"application/json"},
		"content-type": {"application/json"},
	}
}
