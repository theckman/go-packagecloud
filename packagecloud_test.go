package packagecloud

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

type TestSuite struct {
	token         string
	defaultConfig *Config

	server   *http.Server
	listener net.Listener
	mux      *http.ServeMux
	addr     string
	url      string
}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (t *TestSuite) SetUpSuite(c *C) {
	t.addr = "127.0.0.1:18982"
	t.url = fmt.Sprintf("http://%s/", t.addr)

	t.token = "abc123token"
	t.defaultConfig = &Config{
		Token:   t.token,
		BaseURL: t.url,
	}

	listener, err := net.Listen("tcp", "127.0.0.1:18982")
	c.Assert(err, IsNil)

	t.listener = listener

	t.server = &http.Server{
		Addr:         t.addr,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
		ErrorLog:     log.New(ioutil.Discard, "", 0), // disable logging in HTTP server
	}

	go t.server.Serve(t.listener)
}

func (t *TestSuite) SetUpTest(c *C) {
	t.mux = http.NewServeMux()
	t.server.Handler = t.mux
}

func (t *TestSuite) TearDownSuite(c *C) {
	t.listener.Close()
}

func (t *TestSuite) TearDownTest(c *C) {
	os.Unsetenv("PACKAGECLOUD_TOKEN")
}

func (t *TestSuite) TestNewClient(c *C) {
	var cl *Client
	var err error

	//
	// test that no token fails
	//
	cl, err = NewClient(nil)
	c.Assert(err, Not(IsNil))
	c.Check(cl, IsNil)
	c.Check(err.Error(), Equals, "a *Config must be provided with a token or set the PACKAGECLOUD_TOKEN environment variable")

	//
	// test that setting a config token works
	//
	cl, err = NewClient(&Config{Token: t.token})
	c.Assert(cl, Not(IsNil))
	c.Check(err, IsNil)
	c.Check(cl.cfg.Token, Equals, "abc123token")

	//
	// test that setting a PACKAGECLOUD_TOKEN environment variable works
	//
	os.Setenv("PACKAGECLOUD_TOKEN", t.token)
	cl, err = NewClient(nil)
	c.Assert(cl, Not(IsNil))
	c.Check(err, IsNil)
	c.Check(cl.cfg.Token, Equals, "abc123token")
}

func (t *TestSuite) TestString(c *C) {
	var cl *Client
	var err error

	//
	// test that the String() output is correct
	//
	cl, err = NewClient(&Config{Token: t.token})
	c.Assert(err, IsNil)
	c.Check(cl.String(), Equals, "[*packagecloud.Client:<token:abc...oken>]")
}

func (t *TestSuite) TestClient_request(c *C) {
	t.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	cl, err := NewClient(t.defaultConfig)
	c.Assert(err, IsNil)

	b, err := cl.request("GET", t.url)
	c.Assert(err, IsNil)
	c.Check(string(b), Equals, "ok")
}
