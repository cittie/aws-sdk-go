//Package dynamodb provides gucumber integration tests suppport.
package dynamodb

import (
	"github.com/cittie/aws-sdk-go/internal/features/shared"
	"github.com/cittie/aws-sdk-go/service/dynamodb"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@dynamodb", func() {
		World["client"] = dynamodb.New(nil)
	})
}
