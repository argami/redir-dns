package redir_dns

import (
	"net"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(RedirDns{})
}

// RedirDns is a middleware for manipulating redirecting based on DNS TXT record.
type RedirDns struct{}

// CaddyModule returns the Caddy module information.
func (RedirDns) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.redir_dns",
		New: func() caddy.Module { return new(RedirDns) },
	}
}

func (rb RedirDns) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	response, err := net.LookupTXT("_redir_dns." + r.Host)

	if err != nil || len(response) == 0 {
		return next.ServeHTTP(w, r)
	}

	w.Header().Set("Location", response[0])

	return next.ServeHTTP(w, r)
}

// Interface guard
var _ caddyhttp.MiddlewareHandler = (*RedirDns)(nil)
