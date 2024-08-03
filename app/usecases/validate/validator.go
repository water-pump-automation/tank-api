package validate

import (
	"tank-api/app/usecases/validate/types"

	"github.com/xeipuuv/gojsonschema"
)

func init() {
	gojsonschema.FormatCheckers.Add("name_match", types.NameFormat{})
}

func validate(loader interface{}, schemaLoader string) (ValidationError, error) {
	jsonLoader := gojsonschema.NewGoLoader(loader)
	schema := gojsonschema.NewStringLoader(schemaLoader)

	result, err := gojsonschema.Validate(schema, jsonLoader)

	if err != nil {
		return nil, err
	}

	if !result.Valid() {
		validationError := ValidationError{}

		for _, err := range result.Errors() {
			validationError.AddDetails(err)
		}

		return validationError, nil
	}

	return nil, nil
}
