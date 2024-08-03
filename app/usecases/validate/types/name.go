package types

import "regexp"

type NameFormat struct{}

func (NameFormat) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}

	r, _ := regexp.Compile("^[A-Z_1-9]*$")

	return r.MatchString(asString)
}
