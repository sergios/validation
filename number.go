package validation

import (
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/sergios/errors"
)

type NumberValidator struct {
	MinValue int
}

func (n NumberValidator) Validate(key string, value string) *errors.HTTP {
	number, err := strconv.Atoi(value)
	if err != nil || (number < n.MinValue) {
		log.WithFields(log.Fields{
			"key":   key,
			"value": value,
		}).Warn("Invalid param number")
		return errors.HttpParamInvalidError("%s invalid param", key)
	}
	return nil
}
