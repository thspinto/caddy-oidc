package oidc

import (
	log "github.com/sirupsen/logrus"

	"github.com/mholt/caddy"
)

// Config stores the oidc middleware configurations
type Config struct {
	Issuer       string
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

func init() {
	caddy.RegisterPlugin("oidc", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

// setup configures the oidc middleware
func setup(c *caddy.Controller) error {
	for c.Next() {
	}
	log.Info("OIDC setup succesful")
	return nil
}
