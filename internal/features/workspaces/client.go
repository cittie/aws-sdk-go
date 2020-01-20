//Package workspaces provides gucumber integration tests suppport.
package workspaces

import (
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/workspaces"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@workspaces", func() {
		World["client"] = workspaces.New(nil)
	})
}
