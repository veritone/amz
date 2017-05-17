//
// goamz - Go packages to interact with the Amazon Web Services.
//
//   https://wiki.ubuntu.com/goamz
//
package mturk

import (
	"github.com/veritone/amz/aws"
)

func Sign(auth aws.Auth, service, method, timestamp string, params map[string]string) {
	sign(auth, service, method, timestamp, params)
}
