package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bsodmike/go_starter_api/api"
	"github.com/bsodmike/go_starter_api/app"
	"github.com/bsodmike/go_starter_api/routes"
)

var config app.Config
var appApi api.API

func TestMain(m *testing.M) {
	config = app.Config{}
	config.Router = routes.NewRoutes(&config, &appApi)

	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	config.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkHeaderExists(t *testing.T, headerField, header string, response *httptest.ResponseRecorder) {
	val := response.HeaderMap[headerField]

	if val == nil {
		t.Errorf("Expected header to exist. Got %v", val)
	}

	if val[0] != header {
		t.Errorf("Expected header value as `%s`. Got %v", header, val)
	}
}

func sanitizeString(s string) string {
	return strings.TrimSuffix(s, "\n")
}

/*
	TESTS
*/

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health-check", nil)
	response := executeRequest(req)

	checkHeaderExists(t, "Content-Type", "application/json", response)
	checkResponseCode(t, http.StatusOK, response.Code)

	body := response.Body.String()
	if sanitizeString(body) != "{\"alive\":true}" {
		t.Errorf("Expected a valid JSON response. Got %s", body)
	}
}

func TestPublicAPIRoot(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusForbidden, response.Code)

	body := response.Body.String()
	if sanitizeString(body) != "{}" {
		t.Errorf("Expected an empty JSON response. Got %s", body)
	}
}

func TestSecuredAPIRoot(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/secured/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusUnauthorized, response.Code)

	body := response.Body.String()
	if sanitizeString(body) != "Required authorization token not found" {
		t.Errorf("Expected JWT error. Got %s", body)
	}
}
