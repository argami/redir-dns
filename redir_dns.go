package redir_dns

import (
	"net"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(RedirDns{})
	httpcaddyfile.RegisterHandlerDirective("redir_dns", parseCaddyfile)
}

// RedirDns is a RedirDns for manipulating redirecting based on DNS TXT record.
type RedirDns struct{
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (RedirDns) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.redir_dns",
		New: func() caddy.Module { return new(RedirDns) },
	}
}

func (rd *RedirDns) Provision(ctx caddy.Context) error {
	rd.logger = ctx.Logger() // g.logger is a *zap.Logger

	return nil
}

func (rd RedirDns) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	response, err := net.LookupTXT("_redirdns." + r.Host)

	if err != nil || len(response) == 0 {
		rd.logger.Info("error", zap.String("host", "_redirdns." + r.Host), zap.Error(err))
		return next.ServeHTTP(w, r)
	}

	rd.logger.Info("redir_dns", zap.String("host", r.Host), zap.String("response", response[0]))

	w.Header().Set("Location", response[0])
	w.WriteHeader(301)

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (rd *RedirDns) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new RedirDns.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var rd RedirDns
	err := rd.UnmarshalCaddyfile(h.Dispenser)
	return rd, err
}

// Interface guard
var (
	_ caddy.Provisioner           = (*RedirDns)(nil)
	// _ caddy.Validator             = (*RedirDns)(nil)
	_ caddyhttp.MiddlewareHandler = (*RedirDns)(nil)
	_ caddyfile.Unmarshaler       = (*RedirDns)(nil)
)
