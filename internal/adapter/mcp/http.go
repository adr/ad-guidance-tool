package mcp

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"strings"
	"time"

	decisiondomain "github.com/adr/ad-guidance-tool/internal/domain/decision"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// EndpointPath is the path of the MCP endpoint when serving over HTTP.
const EndpointPath = "/mcp"

// ServeStreamableHTTP serves the ADG MCP server over Streamable HTTP on ln.
func ServeStreamableHTTP(ln net.Listener, modelPath string, decisionSvc decisiondomain.DecisionService) error {
	srv := &http.Server{
		Handler:           newHTTPHandler(modelPath, decisionSvc),
		ReadHeaderTimeout: 10 * time.Second,
	}
	return srv.Serve(ln)
}

func newHTTPHandler(modelPath string, decisionSvc decisiondomain.DecisionService) http.Handler {
	// No tool keeps per-client state or sends server-initiated messages, so
	// sessions and a standalone GET stream would only hold memory and connections.
	streamable := mcpserver.NewStreamableHTTPServer(
		buildServer(modelPath, decisionSvc),
		mcpserver.WithStateLess(true),
		mcpserver.WithDisableStreaming(true),
	)

	mux := http.NewServeMux()
	mux.Handle(EndpointPath, rejectDNSRebinding(streamable))
	return mux
}

// A DNS-rebound request reaches a local server under the attacker's hostname,
// so a connection accepted on loopback must also name a loopback host. This is
// the official Go SDK's default check; it leaves connections that arrive over
// the network, such as from a reverse proxy, to that proxy's Host handling.
func rejectDNSRebinding(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if localAddr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok &&
			isLoopback(localAddr.String()) && !isLoopback(r.Host) {
			http.Error(w, fmt.Sprintf("Forbidden: invalid Host header %q", r.Host), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isLoopback(hostport string) bool {
	host, _, err := net.SplitHostPort(hostport)
	if err != nil {
		host = strings.Trim(hostport, "[]")
	}
	if strings.EqualFold(host, "localhost") {
		return true
	}
	ip, err := netip.ParseAddr(host)
	return err == nil && ip.IsLoopback()
}
