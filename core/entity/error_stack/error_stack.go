package error_stack

import "fmt"

type _err struct {
	layer string
	err   error
}

type ErrorStack struct {
	errors        map[string]_err
	usecaseErrors int
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
	} else if len(stack.errors) > 0 {
		layer = fmt.Sprintf("usecase_%d", stack.usecaseErrors)
		stack.usecaseErrors++
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
	usecase := fmt.Sprintf("usecase_%d", stack.usecaseErrors)
	if stackErr, exists := stack.errors[usecase]; exists {
		return stackErr.err
	}

	return nil
}

func (stack *ErrorStack) PopUsecaseError() (err error) {
	usecase := fmt.Sprintf("usecase_%d", stack.usecaseErrors)

	if stackErr, exists := stack.errors[usecase]; exists {
		delete(stack.errors, usecase)
		stack.usecaseErrors--
		return stackErr.err
	}

	return nil
}
