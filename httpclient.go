package tlsdns

import (
	"net"
	"net/http"
	"time"

	"github.com/caddyserver/caddy/v2"
)

// HTTPClientConfig enables customizing HTTP clients via JSON.
type HTTPClientConfig struct {
	Transport *TransportConfig `json:"transport,omitempty"`
	Timeout   caddy.Duration   `json:"timeout,omitempty"`
}

// HTTPClient returns a configured HTTP client.
func (h HTTPClientConfig) HTTPClient() *http.Client {
	if h.Transport == nil {
		h.Transport = defaultTransportConfig
	}
	return &http.Client{
		Transport: h.Transport.Transport(),
		Timeout:   time.Duration(h.Timeout),
	}
}

// TransportConfig enables customizing an HTTP transport via JSON.
type TransportConfig struct {
	Dialer                *DialerConfig  `json:"dialer,omitempty"`
	MaxIdleConns          int            `json:"max_idle_conns,omitempty"`
	IdleConnTimeout       caddy.Duration `json:"idle_conn_timeout,omitempty"`
	TLSHandshakeTimeout   caddy.Duration `json:"tls_handshake_timeout,omitempty"`
	ExpectContinueTimeout caddy.Duration `json:"expect_continue_timeout,omitempty"`
}

// Transport returns a configured HTTP round tripper.
func (t TransportConfig) Transport() http.RoundTripper {
	if t.Dialer == nil {
		t.Dialer = defaultDialerConfig
	}
	if t.MaxIdleConns == 0 {
		t.MaxIdleConns = defaultTransportConfig.MaxIdleConns
	}
	if t.IdleConnTimeout == 0 {
		t.IdleConnTimeout = defaultTransportConfig.IdleConnTimeout
	}
	if t.TLSHandshakeTimeout == 0 {
		t.TLSHandshakeTimeout = defaultTransportConfig.TLSHandshakeTimeout
	}
	if t.ExpectContinueTimeout == 0 {
		t.ExpectContinueTimeout = defaultTransportConfig.ExpectContinueTimeout
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           t.Dialer.Dialer().DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          t.MaxIdleConns,
		IdleConnTimeout:       time.Duration(t.IdleConnTimeout),
		TLSHandshakeTimeout:   time.Duration(t.TLSHandshakeTimeout),
		ExpectContinueTimeout: time.Duration(t.ExpectContinueTimeout),
	}
}

// DialerConfig enables custom network dialers via JSON.
type DialerConfig struct {
	Timeout   caddy.Duration `json:"timeout,omitempty"`
	KeepAlive caddy.Duration `json:"keep_alive,omitempty"`
}

// Dialer returns a configured network dialer.
func (d DialerConfig) Dialer() *net.Dialer {
	if d.Timeout == 0 {
		d.Timeout = defaultDialerConfig.Timeout
	}
	if d.KeepAlive == 0 {
		d.KeepAlive = defaultDialerConfig.KeepAlive
	}
	return &net.Dialer{
		Timeout:   time.Duration(d.Timeout),
		KeepAlive: time.Duration(d.KeepAlive),
	}
}

var (
	defaultTransportConfig = &TransportConfig{
		Dialer:                defaultDialerConfig,
		MaxIdleConns:          100,
		IdleConnTimeout:       caddy.Duration(90 * time.Second),
		TLSHandshakeTimeout:   caddy.Duration(10 * time.Second),
		ExpectContinueTimeout: caddy.Duration(1 * time.Second),
	}

	defaultDialerConfig = &DialerConfig{
		Timeout:   caddy.Duration(30 * time.Second),
		KeepAlive: caddy.Duration(30 * time.Second),
	}
)
