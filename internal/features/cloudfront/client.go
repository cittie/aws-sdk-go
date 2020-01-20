//Package cloudfront provides gucumber integration tests suppport.
package cloudfront

import (
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/cloudfront"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudfront", func() {
		World["client"] = cloudfront.New(nil)
	})
}
