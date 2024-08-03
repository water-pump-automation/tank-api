package validate

import (
	"encoding/json"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const (
	required     = "required"
	invalid_type = "invalid_type"
	format       = "format"
)

var defaultErrors = map[string]string{
	required:     "REQUIRED_ATTRIBUTE_MISSING",
	invalid_type: "INVALID_DATA_TYPE",
	format:       "INVALID_FORMAT",
}

type ValidationError map[string]interface{}

func NewValidationError() ValidationError {
	return make(ValidationError)
}

func (err ValidationError) Error() string {
	bytes, _ := json.Marshal(err)
	return string(bytes)
}

func (err ValidationError) AddDetails(result gojsonschema.ResultError) {
	errField := getInvalidField(result)
	errType := result.Type()
	err[errField] = defaultErrors[errType]
}

func getInvalidField(result gojsonschema.ResultError) (field string) {
	field = result.Field()

	if field == "(root)" {
		property, found := result.Details()["property"]
		if found {
			field = property.(string)
		}
	}
	return strings.Split(field, ".")[0]
}
