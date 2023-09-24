package error_stack

import "fmt"

type _err struct {
	layer string
	err   error
}

type ErrorStack struct {
	errors map[string]_err
}

func (stack *ErrorStack) HasError() bool {
	if len(stack.errors) > 0 {
		return true
	}

	return false
}

func (stack *ErrorStack) Append(err error) {
	var layer string
	if len(stack.errors) == 0 {
		layer = "entity"
	} else if len(stack.errors) == 1 {
		layer = "usecase"
	} else {
		layer = fmt.Sprintf("layer_%d", len(stack.errors)-1)
	}

	stack.errors[layer] = _err{
		layer: layer,
		err:   err,
	}
}

func (stack *ErrorStack) EntityError() (err error) {
	if stackErr, exists := stack.errors["entity"]; exists {
		return stackErr.err
	}

	return nil
}

func (stack *ErrorStack) UsecaseError() (err error) {
	if stackErr, exists := stack.errors["usecase"]; exists {
		return stackErr.err
	}

	return nil
}
