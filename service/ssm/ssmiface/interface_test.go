// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ssmiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/ssm"
	"github.com/upstartmobile/aws-sdk-go/service/ssm/ssmiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*ssmiface.SSMAPI)(nil), ssm.New(nil))
}
