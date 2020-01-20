//Package sts provides gucumber integration tests suppport.
package sts

import (
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/sts"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@sts", func() {
		World["client"] = sts.New(nil)
	})
}
