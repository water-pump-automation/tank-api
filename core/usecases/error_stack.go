package usecases

type UsecaseErrorStack struct {
	errors []error
}

func (stack *UsecaseErrorStack) HasError() bool {
	if len(stack.errors) > 0 {
		return true
	}

	return false
}

func (stack *UsecaseErrorStack) Append(err error) {
	stack.errors = append(stack.errors, err)
}

func (stack *UsecaseErrorStack) UsecaseError() (err error) {
	if stack.HasError() {
		return stack.errors[len(stack.errors)-1]
	}
	return
}

func (stack *UsecaseErrorStack) EntityError() (err error) {
	if stack.HasError() {
		return stack.errors[len(stack.errors)-2]
	}
	return
}
