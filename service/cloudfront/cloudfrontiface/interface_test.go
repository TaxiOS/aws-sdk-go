// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudfrontiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/cloudfront"
	"github.com/upstartmobile/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cloudfrontiface.CloudFrontAPI)(nil), cloudfront.New(nil))
}
