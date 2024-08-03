package validate

import (
	"encoding/json"
	"tank-api/app/usecases/ports"
)

func ValidateInput(input ports.UsecaseInput, databaseInput any, schemaLoader string) error {
	validationErr, err := validate(input, schemaLoader)

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
