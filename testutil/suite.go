//
// goamz - Go packages to interact with the Amazon Web Services.
//
// https://wiki.ubuntu.com/goamz
//
package testutil

import (
	"flag"

	. "gopkg.in/check.v1"

	"github.com/veritone/amz/aws"
)

// Amazon must be used by all tested packages to determine whether to
// run functional tests against the real AWS servers.
var Amazon bool

func init() {
	flag.BoolVar(&Amazon, "amazon", false, "Enable tests against amazon server")
}

type LiveSuite struct {
	auth aws.Auth
}

func (s *LiveSuite) SetUpSuite(c *C) {
	if !Amazon {
		c.Skip("amazon tests not enabled (-amazon flag)")
	}
	auth, err := aws.EnvAuth()
	if err != nil {
		c.Fatal(err.Error())
	}
	s.auth = auth
}
