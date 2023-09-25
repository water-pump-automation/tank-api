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

func (stack *ErrorStack) AppendEntityError(err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}

	layer := "entity"

	stack.errors[layer] = _err{
		layer: layer,
		err:   err,
	}
}

func (stack *ErrorStack) EntityError() (err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}
	if stackErr, exists := stack.errors["entity"]; exists {
		return stackErr.err
	}

	return nil
}

func (stack *ErrorStack) Append(err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}

	stack.usecaseErrors++
	layer := fmt.Sprintf("usecase_%d", stack.usecaseErrors)

	stack.errors[layer] = _err{
		layer: layer,
		err:   err,
	}
}

func (stack *ErrorStack) LastError() (err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}
	usecase := fmt.Sprintf("usecase_%d", stack.usecaseErrors)
	if stackErr, exists := stack.errors[usecase]; exists {
		return stackErr.err
	}

	return nil
}

func (stack *ErrorStack) PopError() (err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}
	usecase := fmt.Sprintf("usecase_%d", stack.usecaseErrors)

	if stackErr, exists := stack.errors[usecase]; exists {
		delete(stack.errors, usecase)
		stack.usecaseErrors--
		return stackErr.err
	}

	return nil
}
