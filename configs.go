package main

import (
	"github.com/mustafa-qamaruddin/simple-dns-proxy/common"
	"github.com/pkg/errors"
	"net/mail"
	"os"
)

func NewConfigs() (*common.Configs, error) {
	cloudFlareApiEmail := os.Getenv("CLOUDFLARE_API_EMAIL")
	if cloudFlareApiEmail == "" {
		return nil, errors.New("environment variable `CLOUDFLARE_API_EMAIL` is required")
	}
	_, err := mail.ParseAddress(cloudFlareApiEmail)
	if err != nil {
		return nil, errors.Wrapf(err, "environment variable `CLOUDFLARE_API_EMAIL` is not a valid e-mail address")
	}
	cloudFlareApiKey := os.Getenv("CLOUDFLARE_API_KEY")
	if cloudFlareApiKey == "" {
		return nil, errors.New("environment variable `CLOUDFLARE_API_EMAIL` is required")
	}
	return &common.Configs{
		CloudFlareApiEmail: cloudFlareApiEmail,
		CloudFlareApiKey:   cloudFlareApiKey,
	}, nil
}
