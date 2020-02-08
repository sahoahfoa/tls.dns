package cloudflare

import (
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddytls"
	tlsdns "github.com/caddyserver/tls.dns"
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/cloudflare"
)

func init() {
	caddy.RegisterModule(Cloudflare{})
}

// CaddyModule returns the Caddy module information.
func (Cloudflare) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "tls.dns.cloudflare",
		New: func() caddy.Module { return new(Cloudflare) },
	}
}

// Cloudflare configures a solver for the ACME DNS challenge.
//
// Please see the following documentation about which credentials
// to supply: https://go-acme.github.io/lego/dns/cloudflare/#api-tokens.
type Cloudflare struct {
	// An API token with the scoped to all applicable domains with the
	// following permissions:
	//
	// - Zone / Zone / Read
	// - Zone / DNS / Edit
	//
	// Or, if you prefer a more strict set of privileges: give this token
	// only the `Zone / DNS / Edit` permission, scoped only to the domains
	// you want to manage certificates for, then provide a ZoneAPIToken
	// scoped to all your zones with the `Zone / Zone / Read` permission.
	APIToken string `json:"api_token,omitempty"`

	// An optional API token used in conjunction with APIToken, only
	// needed if you prefer a stricter set of privileges. If used, this
	// API token must have the `Zone / Zone / Read` for all zones.
	ZoneAPIToken string `json:"zone_api_token,omitempty"`

	tlsdns.CommonConfig
}

// NewDNSProvider returns a DNS challenge solver.
func (wrapper Cloudflare) NewDNSProvider() (challenge.Provider, error) {
	cfg := cloudflare.NewDefaultConfig()
	if wrapper.APIToken != "" {
		cfg.AuthToken = wrapper.APIToken
	}
	if wrapper.ZoneAPIToken != "" {
		cfg.ZoneToken = wrapper.ZoneAPIToken
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
	if wrapper.CommonConfig.HTTPClient != nil {
		cfg.HTTPClient = wrapper.CommonConfig.HTTPClient.HTTPClient()
	}
	return cloudflare.NewDNSProviderConfig(cfg)
}

// Interface guard
var _ caddytls.DNSProviderMaker = (*Cloudflare)(nil)
