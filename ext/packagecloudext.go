package packagecloudext

import (
	"strconv"

	"github.com/theckman/go-packagecloud"
)

// NextDebRelease is a helper function for getting the next release version
// of a Debian package. The idea is to allow your build system to call this
// function to dynamically pick a release number based on what's uploaded to
// packagecloud.
//
// You need to provide a *packagecloud.Client and a *packagecloud.Package
// struct to configure which package you are looking at. You must also specify
// a version string as releases are scoped to versions (i.e., 0.0.1-1, 0.0.1-2, etc.).
func NextDebRelease(client *packagecloud.Client, pkg *packagecloud.Package, packageVersion string) (int64, error) {
	pvs, err := client.Versions(pkg)

	if err != nil {
		return 1, err
	}

	var rel int64

	// loop over the returned package versions
	for _, version := range pvs {
		// if this isn't the right version,
		// move on to the next PackageVersion
		if version.Version != packageVersion {
			continue
		}

		// convert the value to an int64
		// ignore an error here and assume the value is 0
		vRel, _ := strconv.ParseInt(version.Release, 10, 64)

		// if the release version is larger than the biggest release
		// version seen so far let's swap them
		if vRel > rel {
			rel = vRel
		}
	}

	// increment the release number to get the next release number
	rel++

	return rel, nil
}
