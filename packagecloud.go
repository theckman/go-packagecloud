package packagecloud

import (
	"errors"
	"fmt"
	"os"
)

// Version is WISOTT.
const Version = "0.0.1"

// Client is the type for the PackageCloud API client. It has functio methods
// for interacting with the PackageCloud API.
type Client struct {
	token string
}

// Config is the type for the PackageCloud client configuration.
type Config struct {
	// API Token for use with the PackageCloud API
	Token string
}

func (c *Client) String() string {
	var token string

	if len(c.token) > 0 {
		token = fmt.Sprintf("%s...%s", c.token[0:3], c.token[len(c.token)-4:])
	}

	return fmt.Sprintf(
		"[*packagecloud.Client:<token:%s>]",
		token,
	)
}

// NewClient is a function for building a PacakgeCloud *Client instance. If the
// PACKAGECLOUD_TOKEN environment value is set the *Config struct can be nil.
func NewClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = &Config{}
	}

	if cfg.Token == "" {
		cfg.Token = os.Getenv("PACKAGECLOUD_TOKEN")

		if cfg.Token == "" {
			return nil, errors.New("a *Config must be provided with a token or set the PACKAGECLOUD_TOKEN environment variable")
		}
	}

	c := &Client{
		token: cfg.Token,
	}

	return c, nil
}
