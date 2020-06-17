package validation_test

import (
	"github.com/sergios/validation"
	"gopkg.in/check.v1"
)

type uuidSuite struct{}

var _ = check.Suite(&uuidSuite{})

func (s *uuidSuite) TestUUIDNotMatch(c *check.C) {
	uuidNotMatch := "noneoenoe-neoneone-neoneoe"
	result := validation.UUIDMath(uuidNotMatch)
	c.Assert(result, check.Equals, false)
}

func (s *uuidSuite) TestUUIDMatch(c *check.C) {
	uuidMatch := "7553e9ae-a65e-4628-8f0e-48c7b4bbaf62"
	result := validation.UUIDMath(uuidMatch)
	c.Assert(result, check.Equals, true)
}
