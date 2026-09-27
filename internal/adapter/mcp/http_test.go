package mcp

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	svc_mocks "github.com/adr/ad-guidance-tool/mocks/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const initializeRequest = `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"test","version":"0"}}}`

func TestServeStreamableHTTP_InitializesWithoutSession(t *testing.T) {
	baseURL := startServer(t)

	resp, body := send(t, newJSONRPCRequest(t, baseURL+EndpointPath, initializeRequest))

	require.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	assert.Empty(t, resp.Header.Get("Mcp-Session-Id"))
	var out struct {
		Result struct {
			ServerInfo struct{ Name string } `json:"serverInfo"`
		} `json:"result"`
	}
	require.NoError(t, json.Unmarshal(body, &out))
	assert.Equal(t, "adg", out.Result.ServerInfo.Name)
}

func TestServeStreamableHTTP_ListsTools(t *testing.T) {
	baseURL := startServer(t)

	resp, body := send(t, newJSONRPCRequest(t, baseURL+EndpointPath, `{"jsonrpc":"2.0","id":2,"method":"tools/list"}`))

	require.Equal(t, http.StatusOK, resp.StatusCode, string(body))
	var out struct {
		Result struct {
			Tools []struct{ Name string } `json:"tools"`
		} `json:"result"`
	}
	require.NoError(t, json.Unmarshal(body, &out))
	var names []string
	for _, tool := range out.Result.Tools {
		names = append(names, tool.Name)
	}
	assert.ElementsMatch(t, []string{"list_adrs", "get_adr", "get_dsl_reference", "list_rule_files", "validate_rule"}, names)
}

func TestServeStreamableHTTP_OffersNoStandaloneStream(t *testing.T) {
	baseURL := startServer(t)
	req, err := http.NewRequest(http.MethodGet, baseURL+EndpointPath, nil)
	require.NoError(t, err)
	req.Header.Set("Accept", "text/event-stream")

	// Only the status is read: a stream, if one were offered, would never end.
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}

func TestServeStreamableHTTP_ServesOnlyTheEndpointPath(t *testing.T) {
	baseURL := startServer(t)

	resp, _ := send(t, newJSONRPCRequest(t, baseURL+"/", initializeRequest))

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestServeStreamableHTTP_RejectsLoopbackRequestNamingAnotherHost(t *testing.T) {
	baseURL := startServer(t)
	req := newJSONRPCRequest(t, baseURL+EndpointPath, initializeRequest)
	req.Host = "attacker.example"

	resp, _ := send(t, req)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
}

// A test cannot portably listen on a non-loopback address, so this one sets
// the accepting address that net/http would otherwise record in the context.
func TestNewHTTPHandler_AcceptsNetworkRequestNamingPublicHost(t *testing.T) {
	req := newJSONRPCRequest(t, EndpointPath, initializeRequest)
	req.Host = "adr.example.com"
	req = req.WithContext(context.WithValue(req.Context(), http.LocalAddrContextKey, &net.TCPAddr{IP: net.IPv4(10, 0, 0, 5), Port: 8080}))
	rec := httptest.NewRecorder()

	newHTTPHandler(t.TempDir(), new(svc_mocks.DecisionService)).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, rec.Body.String())
}

func TestIsLoopback_LoopbackAddresses(t *testing.T) {
	for _, addr := range []string{"127.0.0.1:8080", "127.0.0.2", "[::1]:8080", "::1", "localhost:8080", "LOCALHOST"} {
		assert.True(t, isLoopback(addr), addr)
	}
}

func TestIsLoopback_OtherAddresses(t *testing.T) {
	for _, addr := range []string{"10.0.0.5:8080", "adr.example.com", "localhost.attacker.example", ""} {
		assert.False(t, isLoopback(addr), addr)
	}
}

func startServer(t *testing.T) string {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	modelPath := t.TempDir()
	go func() { _ = ServeStreamableHTTP(ln, modelPath, new(svc_mocks.DecisionService)) }()
	t.Cleanup(func() { _ = ln.Close() })
	return "http://" + ln.Addr().String()
}

func newJSONRPCRequest(t *testing.T, url, body string) *http.Request {
	t.Helper()
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/event-stream")
	req.Close = true
	return req
}

func send(t *testing.T, req *http.Request) (*http.Response, []byte) {
	t.Helper()
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	return resp, body
}
