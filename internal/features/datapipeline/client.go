//Package datapipeline provides gucumber integration tests suppport.
package datapipeline

import (
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/datapipeline"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@datapipeline", func() {
		World["client"] = datapipeline.New(nil)
	})
}
