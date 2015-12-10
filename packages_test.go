package packagecloud

import (
	"net/http"

	. "gopkg.in/check.v1"
)

const packageVersionsOutput = `[
  {
    "name": "testPackage",
    "distro_version": "ubuntu/trusty",
    "created_at": "2015-11-16T03:39:15.000Z",
    "version": "0.0.1",
    "release": "1",
    "epoch": 0,
    "private": false,
    "type": "deb",
    "filename": "testPackage_0.0.1-1_amd64.deb",
    "uploader_name": "pagerduty",
    "repository_html_url": "/pagerduty/test_repo",
    "package_url": "/api/v1/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/0.0.1-1.json",
    "package_html_url": "/pagerduty/test_repo/packages/ubuntu/trusty/testPackage_0.0.1-1_amd64.deb"
  },
  {
    "name": "testPackage",
    "distro_version": "ubuntu/trusty",
    "created_at": "2015-12-04T16:18:40.000Z",
    "version": "0.0.2",
    "release": "1",
    "epoch": 0,
    "private": true,
    "type": "deb",
    "filename": "testPackage_0.0.2-1_amd64.deb",
    "uploader_name": "pagerduty",
    "repository_html_url": "/pagerduty/test_repo",
    "package_url": "/api/v1/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/0.0.2-1.json",
    "package_html_url": "/pagerduty/test_repo/packages/ubuntu/trusty/testPackage_0.0.2-1_amd64.deb"
  },
  {
    "name": "testPackage",
    "distro_version": "ubuntu/trusty",
    "created_at": "2015-12-10T00:59:42.000Z",
    "version": "0.0.2",
    "release": "2",
    "epoch": 0,
    "private": true,
    "type": "deb",
    "filename": "testPackage_0.0.2-2_amd64.deb",
    "uploader_name": "pagerduty",
    "repository_html_url": "/pagerduty/test_repo",
    "package_url": "/api/v1/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/0.0.2-2.json",
    "package_html_url": "/pagerduty/test_repo/packages/ubuntu/trusty/testPackage_0.0.2-2_amd64.deb"
  }
]`

func (t *TestSuite) TestClient_Versions(c *C) {
	t.mux.HandleFunc("/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/versions.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(packageVersionsOutput))
	})

	cl, err := NewClient(t.defaultConfig)
	c.Assert(err, IsNil)

	var pvs PackageVersions

	r := &Package{
		User:    "pagerduty",
		Repo:    "test_repo",
		Type:    "deb",
		Distro:  "ubuntu",
		Version: "trusty",
		Package: "testPackage",
		Arch:    "amd64",
	}

	pvs, err = cl.Versions(r)
	c.Assert(err, IsNil)
	c.Assert(len(pvs), Equals, 3)

	var p *PackageVersion

	c.Assert(pvs[0], Not(IsNil))
	p = pvs[0]

	c.Check(p.CreatedAt.String(), Equals, "2015-11-16 03:39:15 +0000 UTC")
	c.Check(p.DistroVersion, Equals, "ubuntu/trusty")
	c.Check(p.Epoch, Equals, 0.0)
	c.Check(p.Filename, Equals, "testPackage_0.0.1-1_amd64.deb")
	c.Check(p.Name, Equals, "testPackage")
	c.Check(p.PackageHTMLURL, Equals, "/pagerduty/test_repo/packages/ubuntu/trusty/testPackage_0.0.1-1_amd64.deb")
	c.Check(p.PackageURL, Equals, "/api/v1/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/0.0.1-1.json")
	c.Check(p.Private, Equals, false)
	c.Check(p.Release, Equals, "1")
	c.Check(p.RepositoryHTMLURL, Equals, "/pagerduty/test_repo")
	c.Check(p.Type, Equals, "deb")
	c.Check(p.UploaderName, Equals, "pagerduty")
	c.Check(p.Version, Equals, "0.0.1")

	c.Assert(pvs[1], Not(IsNil))
	p = pvs[1]

	c.Check(p.CreatedAt.String(), Equals, "2015-12-04 16:18:40 +0000 UTC")
	c.Check(p.DistroVersion, Equals, "ubuntu/trusty")
	c.Check(p.Epoch, Equals, 0.0)
	c.Check(p.Filename, Equals, "testPackage_0.0.2-1_amd64.deb")
	c.Check(p.Name, Equals, "testPackage")
	c.Check(p.PackageHTMLURL, Equals, "/pagerduty/test_repo/packages/ubuntu/trusty/testPackage_0.0.2-1_amd64.deb")
	c.Check(p.PackageURL, Equals, "/api/v1/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/0.0.2-1.json")
	c.Check(p.Private, Equals, true)
	c.Check(p.Release, Equals, "1")
	c.Check(p.RepositoryHTMLURL, Equals, "/pagerduty/test_repo")
	c.Check(p.Type, Equals, "deb")
	c.Check(p.UploaderName, Equals, "pagerduty")
	c.Check(p.Version, Equals, "0.0.2")

	c.Assert(pvs[2], Not(IsNil))
	p = pvs[2]

	c.Check(p.CreatedAt.String(), Equals, "2015-12-10 00:59:42 +0000 UTC")
	c.Check(p.DistroVersion, Equals, "ubuntu/trusty")
	c.Check(p.Epoch, Equals, 0.0)
	c.Check(p.Filename, Equals, "testPackage_0.0.2-2_amd64.deb")
	c.Check(p.Name, Equals, "testPackage")
	c.Check(p.PackageHTMLURL, Equals, "/pagerduty/test_repo/packages/ubuntu/trusty/testPackage_0.0.2-2_amd64.deb")
	c.Check(p.PackageURL, Equals, "/api/v1/repos/pagerduty/test_repo/package/deb/ubuntu/trusty/testPackage/amd64/0.0.2-2.json")
	c.Check(p.Private, Equals, true)
	c.Check(p.Release, Equals, "2")
	c.Check(p.RepositoryHTMLURL, Equals, "/pagerduty/test_repo")
	c.Check(p.Type, Equals, "deb")
	c.Check(p.UploaderName, Equals, "pagerduty")
	c.Check(p.Version, Equals, "0.0.2")
}
