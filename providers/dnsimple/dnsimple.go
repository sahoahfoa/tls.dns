package dnsimple

import (
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddytls"
	tlsdns "github.com/caddyserver/tls.dns"
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/dnsimple"
)

func init() {
	caddy.RegisterModule(DNSimple{})
}

// CaddyModule returns the Caddy module information.
func (DNSimple) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "tls.dns.dnsimple",
		New: func() caddy.Module { return new(DNSimple) },
	}
}

// DNSimple configures a solver for the ACME DNS challenge.
type DNSimple struct {
	// An OAuth2 token from your account.
	AccessToken string `json:"access_token,omitempty"`

	tlsdns.CommonConfig
}

// NewDNSProvider returns a DNS challenge solver.
func (wrapper DNSimple) NewDNSProvider() (challenge.Provider, error) {
	cfg := dnsimple.NewDefaultConfig()
	if wrapper.AccessToken != "" {
		cfg.AccessToken = wrapper.AccessToken
	}
	if wrapper.CommonConfig.BaseURL != "" {
		cfg.BaseURL = wrapper.CommonConfig.BaseURL
	}
	if wrapper.CommonConfig.TTL != 0 {
		cfg.TTL = wrapper.CommonConfig.TTL
	}
	if wrapper.CommonConfig.PropagationTimeout != 0 {
		cfg.PropagationTimeout = time.Duration(wrapper.CommonConfig.PropagationTimeout)
	}
	if wrapper.CommonConfig.PollingInterval != 0 {
		cfg.PollingInterval = time.Duration(wrapper.CommonConfig.PollingInterval)
	}
	return dnsimple.NewDNSProviderConfig(cfg)
}

// Interface guard
var _ caddytls.DNSProviderMaker = (*DNSimple)(nil)
