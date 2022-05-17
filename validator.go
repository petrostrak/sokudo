package sokudo

import (
	"net/http"
	"net/url"
)

type Validation struct {
	Data   url.Values
	Errors map[string]string
}

func (s *Sokudo) Validator(data url.Values) *Validation {
	return &Validation{
		Data:   data,
		Errors: make(map[string]string),
	}
}

func (v *Validation) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validation) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validation) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	return x == ""
}
