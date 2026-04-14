//go:build functional

package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/CiscoDevNet/go-ciscosecureaccess/client"
)

type functionalEnv struct {
	ctx     context.Context
	factory client.SSEClientFactory
}

func newFunctionalEnv(t *testing.T) functionalEnv {
	t.Helper()

	if !envEnabled("GO_CISCOSECUREACCESS_RUN_FUNCTIONAL") {
		t.Skip("set GO_CISCOSECUREACCESS_RUN_FUNCTIONAL=1 to run functional tests")
	}

	keyID := strings.TrimSpace(os.Getenv("API_KEY_ID"))
	keySecret := strings.TrimSpace(os.Getenv("API_KEY_SECRET"))
	if keyID == "" || keySecret == "" {
		t.Skip("set API_KEY_ID and API_KEY_SECRET to run functional tests")
	}

	return functionalEnv{
		ctx: context.Background(),
		factory: client.SSEClientFactory{
			KeyId:       keyID,
			KeySecret:   keySecret,
			ApiEndpoint: strings.TrimSpace(os.Getenv("GO_CISCOSECUREACCESS_API_ENDPOINT")),
		},
	}
}

func envEnabled(name string) bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(name))) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func envInt64(name string) (int64, bool, error) {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return 0, false, nil
	}

	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, true, fmt.Errorf("parse %s: %w", name, err)
	}

	return value, true, nil
}

func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

func isHTTP2xx(resp *http.Response) bool {
	return resp != nil && resp.StatusCode >= 200 && resp.StatusCode < 300
}
