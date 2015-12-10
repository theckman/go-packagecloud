package packagecloudext

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/theckman/go-packagecloud"

	. "gopkg.in/check.v1"
)

type TestSuite struct {
	token         string
	defaultConfig *packagecloud.Config

	server   *http.Server
	listener net.Listener
	mux      *http.ServeMux
	addr     string
	url      string
}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (t *TestSuite) SetUpSuite(c *C) {
	t.addr = "127.0.0.1:18983"
	t.url = fmt.Sprintf("http://%s/", t.addr)

	t.token = "abc123token"
	t.defaultConfig = &packagecloud.Config{
		Token:   t.token,
		BaseURL: t.url,
	}

	listener, err := net.Listen("tcp", "127.0.0.1:18983")
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

func (t *TestSuite) Test_NextDebRelease(c *C) {
	t.mux.HandleFunc("/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/versions.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(packageVersionsOutput))
	})

	cl, err := packagecloud.NewClient(t.defaultConfig)
	c.Assert(err, IsNil)

	var release int64

	r := &packagecloud.Package{
		User:    "pagerduty",
		Repo:    "test_repo",
		Type:    "deb",
		Distro:  "ubuntu",
		Version: "trusty",
		Package: "testPackage",
		Arch:    "amd64",
	}

	release, err = NextDebRelease(cl, r, "0.0.1")
	c.Assert(err, IsNil)
	c.Check(release, Equals, int64(2))

	release, err = NextDebRelease(cl, r, "0.0.2")
	c.Assert(err, IsNil)
	c.Check(release, Equals, int64(3))

	release, err = NextDebRelease(cl, r, "0.0.3")
	c.Assert(err, IsNil)
	c.Check(release, Equals, int64(1))
}
