package validation

import (
	"regexp"

	"github.com/sergios/errors"
)

const formatUUID = "^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$"

var reUUID = regexp.MustCompile(formatUUID)

type UUIDValidator struct{}

func (u UUIDValidator) Validate(key string, value string) *errors.HTTP {
	if isOK := UUIDMath(value); !isOK {
		return errors.HttpParamInvalidError("%s invalid param", key)
	}
	return nil
}

func UUIDMath(value string) bool {
	return reUUID.MatchString(value)
}
