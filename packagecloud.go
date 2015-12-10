package packagecloud

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hashicorp/go-cleanhttp"
)

// Version is WISOTT.
const Version = "0.0.1"

const baseURL = "https://packagecloud.io/api/v1/"

// Client is the type for the PackageCloud API client. It has functio methods
// for interacting with the PackageCloud API.
type Client struct {
	cfg    *Config
	client *http.Client
}

// Config is the type for the PackageCloud client configuration.
type Config struct {
	// API Token for use with the PackageCloud API
	Token string

	// BaseURL is the base URL for the API if this is left empty it defaults
	// to "https://packagecloud.io/api/v1/"
	BaseURL string
}

func (c *Client) String() string {
	var token string

	if len(c.cfg.Token) > 0 {
		token = fmt.Sprintf("%s...%s", c.cfg.Token[0:3], c.cfg.Token[len(c.cfg.Token)-4:])
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

	if cfg.BaseURL == "" {
		cfg.BaseURL = baseURL
	}

	c := &Client{
		cfg:    cfg,
		client: cleanhttp.DefaultClient(),
	}

	return c, nil
}

func (c *Client) request(method, path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.cfg.BaseURL, path)

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.cfg.Token, "")

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
