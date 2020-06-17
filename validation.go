package validation

import (
	goerrors "errors"
	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/sergios/errors"
)

type Validator interface {
	Validate(key, value string) *errors.HTTP
}

var validators = map[string]Validator{}

func AddParamValidator(key string, f Validator) error {
	if _, found := validators[key]; !found {
		validators[key] = f
		return nil
	}
	return goerrors.New(fmt.Sprintf("Validator %s já adicionado", key))
}

func ValidatorParams(ps httprouter.Params) error {
	for i := range ps {
		value := ps[i].Value
		key := ps[i].Key

		if len(value) <= 0 {
			return errors.HttpParamInvalidError("%s invalid param", key)
		}

		if f, found := validators[key]; found {
			if err := f.Validate(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}
