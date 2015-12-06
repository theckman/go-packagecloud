package packagecloud

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

type TestSuite struct {
	token string
}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (t *TestSuite) SetUpSuite(c *C) {
	t.token = "abc123token"
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
	c.Check(cl.token, Equals, "abc123token")

	//
	// test that setting a PACKAGECLOUD_TOKEN environment variable works
	//
	os.Setenv("PACKAGECLOUD_TOKEN", t.token)
	cl, err = NewClient(nil)
	c.Assert(cl, Not(IsNil))
	c.Check(err, IsNil)
	c.Check(cl.token, Equals, "abc123token")
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
