package validate

import (
	"context"
	"encoding/json"
	"water-tank-api/app/entity/validation"
	"water-tank-api/app/usecases/ports"
)

func ValidateInput(ctx context.Context, input ports.UsecaseInput, databaseInput any, schemaLoader string) error {
	validationErr, err := validation.Validate(ctx, input, schemaLoader)

	if err != nil {
		return err
	}

	if validationErr != nil {
		return validationErr
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = json.Unmarshal(inputBytes, databaseInput)
	if err != nil {
		return err
	}

	return nil
}
