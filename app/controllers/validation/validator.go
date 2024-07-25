package validation

import (
	"context"
	"regexp"
	"strings"
	"water-tank-api/app/controllers/response"

	"github.com/xeipuuv/gojsonschema"
)

func init() {
	gojsonschema.FormatCheckers.Add("name_match", nameFormat{})
}

type nameFormat struct{}

func (nameFormat) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}

	r, _ := regexp.Compile("^[A-Z_]*$")

	return r.MatchString(asString)
}

func Validate(ctx context.Context, loader interface{}, schemaLoader string) *response.ControllerResponse {

	jsonLoader := gojsonschema.NewGoLoader(loader)
	schema := gojsonschema.NewStringLoader(schemaLoader)

	result, err := gojsonschema.Validate(schema, jsonLoader)

	if err != nil {
		return response.NewControllerError(response.WaterTankInternalServerError, "Validation error")
	}

	if !result.Valid() {
		validationError := response.NewValidationError()

		for _, err := range result.Errors() {
			field := err.Field()

			if field == "(root)" {
				property, found := err.Details()["property"]
				if found {
					field = property.(string)
				}
			}

			validationError.AddDetails(strings.Split(field, ".")[0], defaultErrors[err.Type()])
		}

		return validationError
	}

	return nil
}
