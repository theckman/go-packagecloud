package packagecloud

import (
	"fmt"
	"net/http"
	"strings"

	. "gopkg.in/check.v1"
)

func (t *TestSuite) TestClient_Licenses(c *C) {
	t.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out := `{
  "license": "%s",
  "signature": "testSignature"
}
`
		path := strings.Split(r.URL.Path, "/")
		w.Write([]byte(fmt.Sprintf(out, path[2])))
	})

	cl, err := NewClient(t.defaultConfig)
	c.Assert(err, IsNil)

	var l *License

	l, err = cl.Licenses("testLicense42")
	c.Assert(err, IsNil)
	c.Check(l.License, Equals, "testLicense42")
	c.Check(l.Signature, Equals, "testSignature")
}
