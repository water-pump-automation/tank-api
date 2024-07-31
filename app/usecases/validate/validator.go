package validate

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func init() {
	gojsonschema.FormatCheckers.Add("name_match", nameFormat{})
}

func validate(loader interface{}, schemaLoader string) (ValidationError, error) {
	jsonLoader := gojsonschema.NewGoLoader(loader)
	schema := gojsonschema.NewStringLoader(schemaLoader)

	result, err := gojsonschema.Validate(schema, jsonLoader)

	if err != nil {
		return nil, err
	}

	if !result.Valid() {
		validationError := NewValidationError()

		for _, err := range result.Errors() {
			field := err.Field()

			if field == "(root)" {
				property, found := err.Details()["property"]
				if found {
					field = property.(string)
				}
			}

			validationError.AddDetails(strings.Split(field, ".")[0], err.Type())
		}

		return validationError, nil
	}

	return nil, nil
}

// *************************************************
//			CUSTOM VALIDATION
// *************************************************

type nameFormat struct{}

func (nameFormat) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}

	r, _ := regexp.Compile("^[A-Z_1-9]*$")

	return r.MatchString(asString)
}

// *************************************************
//			VALIDATION ERROR
// *************************************************

type ValidationError map[string]interface{}

func NewValidationError() ValidationError {
	return make(ValidationError)
}

func (err ValidationError) AddDetails(field, value string) {
	err[field] = defaultErrors[value]
}

func (err ValidationError) Error() string {
	bytes, _ := json.Marshal(err)
	return string(bytes)
}

var defaultErrors = map[string]string{
	"required":                        "REQUIRED_ATTRIBUTE_MISSING",
	"invalid_type":                    "INVALID_DATA_TYPE",
	"number_any_of":                   "INVALID_DATA_TYPE",
	"number_one_of":                   "INVALID_DATA_TYPE",
	"number_all_of":                   "INVALID_DATA_TYPE",
	"number_not":                      "INVALID_DATA_TYPE",
	"missing_dependency":              "INVALID_DATA_TYPE",
	"internal":                        "INVALID_DATA_TYPE",
	"const":                           "INVALID_DATA_TYPE",
	"enum":                            "INVALID_VALUE",
	"array_no_additional_items":       "INVALID_DATA_TYPE",
	"array_min_items":                 "INVALID_DATA_TYPE",
	"array_max_items":                 "INVALID_DATA_TYPE",
	"unique":                          "INVALID_DATA_TYPE",
	"contains":                        "INVALID_DATA_TYPE",
	"array_min_properties":            "INVALID_DATA_TYPE",
	"array_max_properties":            "INVALID_DATA_TYPE",
	"additional_property_not_allowed": "INVALID_DATA_TYPE",
	"invalid_property_pattern":        "INVALID_DATA_TYPE",
	"invalid_property_name":           "INVALID_DATA_TYPE",
	"string_gte":                      "INVALID_LENGTH",
	"string_lte":                      "INVALID_LENGTH",
	"pattern":                         "INVALID_DATA_TYPE",
	"multiple_of":                     "INVALID_DATA_TYPE",
	"number_gte":                      "INVALID_VALUE",
	"number_gt":                       "INVALID_VALUE",
	"number_lte":                      "INVALID_VALUE",
	"number_lt":                       "INVALID_VALUE",
	"condition_then":                  "INVALID_DATA_TYPE",
	"condition_else":                  "INVALID_DATA_TYPE",
	"format":                          "INVALID_FORMAT",
}
