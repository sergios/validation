package validation_test

import (
	"github.com/sergios/errors"
	"github.com/sergios/validation"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/check.v1"
)

type validationSuite struct{}

var _ = check.Suite(&validationSuite{})

func (s *validationSuite) TestAddParamValidator(c *check.C) {
	err := validation.AddParamValidator("feedId", nil)
	c.Assert(err, check.IsNil)

	err = validation.AddParamValidator("feedId", nil)
	c.Assert(err, check.NotNil)

	err = validation.AddParamValidator("postId", nil)
	c.Assert(err, check.IsNil)
}

func (s *validationSuite) TestValidatorParamsValid(c *check.C) {
	err := validation.AddParamValidator("testeId", validation.UUIDValidator{})
	c.Assert(err, check.IsNil)

	ps := httprouter.Params{
		httprouter.Param{Key: "testeId", Value: "8f9d69dc-d991-4c79-a796-943bc39d9abc"},
		httprouter.Param{Key: "xpto", Value: "xpto"},
	}
	err = validation.ValidatorParams(ps)
	c.Assert(err, check.IsNil)
}

func (s *validationSuite) TestValidatorParamsUUIDInvalid(c *check.C) {
	err := validation.AddParamValidator("testeIdInvalid", validation.UUIDValidator{})
	c.Assert(err, check.IsNil)

	ps := httprouter.Params{
		httprouter.Param{Key: "testeIdInvalid", Value: "8f9d69dc-d991-4c79-a796-noneoneo"},
	}
	err = validation.ValidatorParams(ps)
	c.Assert(err, check.NotNil)

	e, ok := err.(*errors.HTTP)
	c.Assert(ok, check.Equals, true)
	c.Assert(e.Code, check.Equals, errors.StatusUnprocessable)
}

func (s *validationSuite) TestValidatorParamsEmpty(c *check.C) {
	ps := httprouter.Params{
		httprouter.Param{Key: "t1", Value: "t1t1t1t1t1"},
		httprouter.Param{Key: "t2", Value: ""},
	}
	err := validation.ValidatorParams(ps)
	c.Assert(err, check.NotNil)

	e, ok := err.(*errors.HTTP)
	c.Assert(ok, check.Equals, true)
	c.Assert(e.Code, check.Equals, errors.StatusUnprocessable)
	c.Assert(e.Message, check.Equals, "t2 invalid param")
}
