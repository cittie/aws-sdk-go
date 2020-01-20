//Package directconnect provides gucumber integration tests suppport.
package directconnect

import (
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/directconnect"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@directconnect", func() {
		World["client"] = directconnect.New(nil)
	})
}
