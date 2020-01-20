//Package efs provides gucumber integration tests suppport.
package efs

import (
	"github.com/cittie/aws-sdk-go/aws"
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(&aws.Config{Region: "us-west-2"})
	})
}
