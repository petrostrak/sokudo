package sokudo

import "net/url"

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
