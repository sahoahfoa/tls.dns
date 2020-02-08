package tlsdns

import "github.com/caddyserver/caddy/v2"

// CommonConfig contains some config parameters that are
// common among many DNS providers. Note that some DNS
// providers support most but not all of these parameters.
// It is not an error to configure them in that case, but
// they might not be used.
type CommonConfig struct {
	// The base URL to use for the provider API.
	BaseURL string `json:"base_url,omitempty"`

	// The TTL of the TXT record used for the DNS challenge.
	TTL int `json:"ttl,omitempty"`

	// Maximum waiting time for DNS propagation.
	PropagationTimeout caddy.Duration `json:"propagation_timeout,omitempty"`

	// Time between DNS propagation check.
	PollingInterval caddy.Duration `json:"polling_interval,omitempty"`

	// HTTP client customizations, if necessary.
	HTTPClient *HTTPClientConfig `json:"http_client,omitempty"`
}
