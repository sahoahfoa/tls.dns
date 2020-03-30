package rfc2136

import (
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddytls"
	tlsdns "github.com/caddyserver/tls.dns"
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns/rfc2136"
)

func init() {
	caddy.RegisterModule(RFC2136{})
}

// CaddyModule returns the Caddy module information.
func (RFC2136) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "tls.dns.rfc2136",
		New: func() caddy.Module { return new(RFC2136) },
	}
}

// rfc2136 configures a solver for the ACME DNS challenge.
type RFC2136 struct {
	
	// Network address in the form "host" or "host:port".
	Nameserver string `json:"nameserver,omitempty"`
	
	// Defaults to hmac-md5.sig-alg.reg.int. (HMAC-MD5).
	// See https://github.com/miekg/dns/blob/master/tsig.go for supported values.
	// Additional information: https://tools.ietf.org/html/rfc4635#section-2
	TSIGAlgorithm string `json:"tsig_algorithm,omitempty"`
	
	// Name of the secret key as defined in DNS server configuration.
	TSIGKey string `json:"tsig_key,omitempty"`
	
	// Secret key as defined in DNS server configuration.
	TSIGSecret string `json:"tsig_secret,omitempty"`
	
	// DNS Client specific values
	DNSClient DNSClientConfig `json:"dns_client,omitempty"`

	tlsdns.CommonConfig
}

// DNSClientConfig enables customizing DNS clients via JSON.
type DNSClientConfig struct {
	// Time between DNS challenges that will be resolved sequentially
    //
    // Duration is a JSON-string-unmarshable duration type.
	SequenceInterval time.Duration `json:"sequence_interval,omitempty"`
	
	// Maximum wait time to connect to nameserver
    //
    // Duration is a JSON-string-unmarshable duration type.
	DNSTimeout time.Duration `json:"dns_timeout,omitempty"`
}

// NewDNSProvider returns a DNS challenge solver.
func (wrapper RFC2136) NewDNSProvider() (challenge.Provider, error) {
	cfg := rfc2136.NewDefaultConfig()
	if wrapper.Nameserver != "" {
		cfg.Nameserver = wrapper.Nameserver
	}
	if wrapper.TSIGAlgorithm != "" {
		cfg.TSIGAlgorithm = wrapper.TSIGAlgorithm
	}
	if wrapper.TSIGKey != "" {
		cfg.TSIGKey = wrapper.TSIGKey
	}
	if wrapper.TSIGSecret != "" {
		cfg.TSIGSecret = wrapper.TSIGSecret
	}
	if wrapper.DNSClient.SequenceInterval != 0 {
		cfg.SequenceInterval = wrapper.DNSClient.SequenceInterval
	}
	if wrapper.DNSClient.DNSTimeout != 0 {
		cfg.DNSTimeout = wrapper.DNSClient.DNSTimeout
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

	return rfc2136.NewDNSProviderConfig(cfg)
}

// Interface guard
var _ caddytls.DNSProviderMaker = (*RFC2136)(nil)
