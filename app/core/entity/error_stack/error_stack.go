package error_stack

import "fmt"

type _err struct {
	layer string
	err   error
}

type Error struct {
	errors        map[string]_err
	usecaseErrors int
}

func (stack *Error) HasError() bool {
	return len(stack.errors) > 0
}

func (stack *Error) AddEntityError(err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}

	layer := "entity"
	if _, exists := stack.errors[layer]; !exists {
		stack.errors[layer] = _err{
			layer: layer,
			err:   err,
		}
	}
}

func (stack *Error) EntityError() (err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}
	if stackErr, exists := stack.errors["entity"]; exists {
		return stackErr.err
	}

	return nil
}

func (stack *Error) AppendUsecaseError(err error) {
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

func (stack *Error) LastUsecaseError() (err error) {
	if stack.errors == nil {
		stack.errors = make(map[string]_err)
	}
	usecase := fmt.Sprintf("usecase_%d", stack.usecaseErrors)
	if stackErr, exists := stack.errors[usecase]; exists {
		return stackErr.err
	}

	if stackErr, exists := stack.errors["entity"]; exists {
		return stackErr.err
	}

	return nil
}

func (stack *Error) PopUsecaseError() (err error) {
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
