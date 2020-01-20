//Package kinesis provides gucumber integration tests suppport.
package kinesis

import (
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/kinesis"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@kinesis", func() {
		World["client"] = kinesis.New(nil)
	})
}
