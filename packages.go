package packagecloud

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// PackageVersion is a struct containing version information about packagecloud
// packages.
type PackageVersion struct {
	CreatedAt         *time.Time `json:"created_at"`
	DistroVersion     string     `json:"distro_version"`
	Epoch             float64    `json:"epoch"`
	Filename          string     `json:"filename"`
	Name              string     `json:"name"`
	PackageHTMLURL    string     `json:"package_html_url"`
	PackageURL        string     `json:"package_url"`
	Private           bool       `json:"private"`
	Release           string     `json:"release"`
	RepositoryHTMLURL string     `json:"repository_html_url"`
	Type              string     `json:"type"`
	UploaderName      string     `json:"uploader_name"`
	Version           string     `json:"version"`
}

func (pv *PackageVersion) String() string {
	return fmt.Sprintf(
		"<package: %s [distro: %s|version: %s|release: %s|private: %t]>",
		pv.Name, pv.DistroVersion, pv.Version, pv.Release, pv.Private,
	)
}

// Package is a struct to be passed in to certain *packagecloud.Client functions
// to specify which location to use.
type Package struct {
	User    string // the user this repoistory belongs to
	Repo    string // the name of the repository
	Type    string // the type of package it is (e.g., "deb" or "rpm")
	Distro  string // the name of the distribution the package is in (e.g., "ubuntu")
	Version string // the version of the distro the package is in (e.g., trusty)
	Package string // the name of the package
	Arch    string // the architecture of the package; deb 64-bit: "amd64"; rpm 64-bit: "x86_64"
}

// PackageVersions is a slice of individual Package Version structs.
type PackageVersions []*PackageVersion

func (pvs PackageVersions) String() string {
	var pckgs []string

	for _, p := range pvs {
		pckgs = append(pckgs, p.String())
	}

	return fmt.Sprintf("[ %s ]", strings.Join(pckgs, ", "))
}

// Versions returns the package versions for a specific package. You must
// specify a package to be retrieved.
func (c *Client) Versions(r *Package) (PackageVersions, error) {
	path := fmt.Sprintf(
		"repos/%s/%s/package/%s/%s/%s/%s/%s/versions.json",
		r.User, r.Repo, r.Type, r.Distro,
		r.Version, r.Package, r.Arch,
	)

	b, err := c.request("GET", path)

	if err != nil {
		return nil, err
	}

	pvs := make(PackageVersions, 0)

	err = json.Unmarshal(b, &pvs)

	return pvs, err
}
