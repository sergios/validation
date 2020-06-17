package validation_test

import (
	"github.com/sergios/errors"
	"github.com/sergios/validation"
	"gopkg.in/check.v1"
)

type numberSuite struct{}

var _ = check.Suite(&numberSuite{})

func (s *numberSuite) TestNumberPositive(c *check.C) {
	numberV := validation.NumberValidator{}
	err := numberV.Validate("pageNumber", "17823654")
	c.Assert(err, check.IsNil)
}

func (s *numberSuite) TestNumberNegative(c *check.C) {
	numberV := validation.NumberValidator{}
	err := numberV.Validate("pageNumber", "-1")
	c.Assert(err, check.NotNil)
	c.Assert(err.Code, check.Equals, errors.StatusUnprocessable)
}

func (s *numberSuite) TestNumberNotInt(c *check.C) {
	numberV := validation.NumberValidator{}
	err := numberV.Validate("pageNumber", "nonene")
	c.Assert(err, check.NotNil)
	c.Assert(err.Code, check.Equals, errors.StatusUnprocessable)
}

func (s *numberSuite) TestNumberMinValue(c *check.C) {
	numberV := validation.NumberValidator{
		MinValue: 1,
	}
	err := numberV.Validate("pageNumber", "1")
	c.Assert(err, check.IsNil)

	err = numberV.Validate("pageNumber", "0")
	c.Assert(err, check.NotNil)
	c.Assert(err.Code, check.Equals, errors.StatusUnprocessable)
}

func (s *numberSuite) TestNumberDefault(c *check.C) {
	numberV := validation.NumberValidator{}
	err := numberV.Validate("pageNumber", "1")
	c.Assert(err, check.IsNil)
	err = numberV.Validate("pageNumber", "0")
	c.Assert(err, check.IsNil)
	err = numberV.Validate("pageNumber", "-1")
	c.Assert(err, check.NotNil)
}

func (s *numberSuite) TestNumberZero(c *check.C) {
	numberV := validation.NumberValidator{
		MinValue: 0,
	}
	err := numberV.Validate("pageNumber", "1")
	c.Assert(err, check.IsNil)
	err = numberV.Validate("pageNumber", "0")
	c.Assert(err, check.IsNil)
	err = numberV.Validate("pageNumber", "-1")
	c.Assert(err, check.NotNil)
}
