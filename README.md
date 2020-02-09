DNS Providers for Caddy
=======================

This repository is for Caddy modules in the `tls.dns` namespace (DNS challenge solvers).

Generally, these are simply config wrappers over [go-acme/lego's DNS providers](https://pkg.go.dev/github.com/go-acme/lego/v3/providers/dns?tab=subdirectories). **We need your help to finish adding all 60+ providers (see below)!**

## Using them in Caddy

First [plug in your desired provider module](https://github.com/caddyserver/caddy/tree/v2#building-with-plugins) and then [configure the DNS challenge](https://caddyserver.com/docs/json/apps/tls/automation/policies/management/acme/challenges/) in your Caddy JSON, for example:

```json
{
	"challenges": {
		"dns": {
			"provider": "cloudflare",
			"api_token": "secret"
		}
	}
}
```

If you prefer, you can set the provider configuration using environment variables. See [the go-acme/lego documentation](https://go-acme.github.io/lego/dns/) to know which variables to set for which providers. (You still have to at least enable the DNS challenge for your provider in the JSON, though. You can use env vars for credentials, etc.) Parameters that you set via env variables can be omitted from the JSON.


## Adding new providers

[go-acme/lego supports dozens of DNS providers](https://github.com/go-acme/lego/tree/master/providers/dns), and Caddy can use all of them too, but a thin wrapper layer (abstraction) is required. The wrapper's job is simply to get a config from the environment from the lego provider package, then set any config fields that are explicitly configured (overriding those from the environment).

- **If lego does not already support your DNS provider:** it is probably best to first submit the implementation to lego via pull request. (It could live anywhere, but might as well contribute it back to the core lib, right?) Once lego supports it, it is easy for Caddy to use it.

- **If lego already supports your DNS provider:** [make sure it has a `NewConfigFromEnv()` function](https://github.com/go-acme/lego/issues/1054), otherwise you will not be able to set credentials in environment variables. We require that all our wrappers support loading config from the environment. Not all providers in lego support this yet, so you may need to submit a pull request to lego first.

- **If your provider in lego already has a `NewConfigFromEnv()` function:** then simply copy [one of the existing provider packages in this repo](https://github.com/caddyserver/tls.dns/tree/master/providers) and change it to work with your provider. Then submit a pull request to this repo. Please be sure the wrapper calls `NewConfigFromEnv()` in order to support credentials from environment variables!

Thanks for contributing!
