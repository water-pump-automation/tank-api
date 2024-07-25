package error_stack

import (
	"errors"
	"testing"
)

var mockEntityErr = errors.New("entity error")
var mockUsecaseError = errors.New("usecase error")
var mockUsecaseError2 = errors.New("usecase2 error")

func Test_ErrorStack_HasError(t *testing.T) {
	Test_ErrorStack_HasEntityError := func(t *testing.T) {
		stack := Error{}
		stack.AddEntityError(mockEntityErr)

		t.Run("Check if has entity error", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_HasEntityError() should report an error, but it haven't")
			}
		})
	}

	Test_ErrorStack_HasUsecaseError := func(t *testing.T) {
		stack := Error{}
		stack.AppendUsecaseError(mockUsecaseError)

		t.Run("Check if has usecase error", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_HasUsecaseError() should report an error, but it haven't")
			}
		})
	}

	Test_ErrorStack_DoesntHaveErrorDueToPop := func(t *testing.T) {
		stack := Error{}
		stack.AppendUsecaseError(mockUsecaseError)

		t.Run("Check if has usecase error", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_DoesntHaveErrorDueToPop() should report an error, but it haven't")
			}
		})

		stack.PopUsecaseError()
		t.Run("Check if has has cleared usecase errors", func(t *testing.T) {
			if stack.HasError() {
				t.Error("Test_ErrorStack_DoesntHaveErrorDueToPop() shouldn't report an error, but it have")
			}
		})
	}

	Test_ErrorStack_DoesHaveErrorAfterPop := func(t *testing.T) {
		stack := Error{}
		stack.AddEntityError(mockUsecaseError)

		t.Run("Check if has usecase error", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_DoesHaveErrorAfterPop() should report an error, but it haven't")
			}
		})

		stack.PopUsecaseError()
		t.Run("Check if has has cleared usecase errors", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_DoesHaveErrorAfterPop() shouldn't report an error, but it have")
			}
		})
	}

	Test_ErrorStack_HasEntityError(t)
	Test_ErrorStack_HasUsecaseError(t)
	Test_ErrorStack_DoesntHaveErrorDueToPop(t)
	Test_ErrorStack_DoesHaveErrorAfterPop(t)
}

func Test_ErrorStack_PopError(t *testing.T) {
	Test_ErrorStack_SuccessfulPop := func(t *testing.T) {
		stack := Error{}
		stack.AppendUsecaseError(mockUsecaseError)

		t.Run("Check if has usecase error", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_SuccessfulPop() should report an error, but it haven't")
			}
		})

		t.Run("Check if has error returned is correct", func(t *testing.T) {
			err := stack.PopUsecaseError()
			if err != mockUsecaseError {
				t.Errorf("Test_ErrorStack_SuccessfulPop() returned '%s', when it should return '%s'", err.Error(), mockUsecaseError.Error())
			}
		})
	}

	Test_ErrorStack_MultipleSuccessfulPop := func(t *testing.T) {
		stack := Error{}
		stack.AppendUsecaseError(mockUsecaseError)
		stack.AppendUsecaseError(mockUsecaseError2)

		t.Run("Check if has usecase error", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_MultipleSuccessfulPop() should report an error, but it haven't")
			}
		})

		err1 := stack.PopUsecaseError()
		err2 := stack.PopUsecaseError()
		t.Run("Check if err1 is correct", func(t *testing.T) {
			if err1 != mockUsecaseError2 {
				t.Errorf("Test_ErrorStack_MultipleSuccessfulPop() returned '%s', when it should return '%s'", err1.Error(), mockUsecaseError2.Error())
			}
		})

		t.Run("Check if err2 is correct", func(t *testing.T) {
			if err2 != mockUsecaseError {
				t.Errorf("Test_ErrorStack_MultipleSuccessfulPop() returned '%s', when it should return '%s'", err2.Error(), mockUsecaseError.Error())
			}
		})

	}

	Test_ErrorStack_ErrorRemovalFromStack := func(t *testing.T) {
		stack := Error{}
		stack.AppendUsecaseError(mockUsecaseError)
		stack.AppendUsecaseError(mockUsecaseError2)

		t.Run("Check if has usecase error", func(t *testing.T) {
			if !stack.HasError() {
				t.Error("Test_ErrorStack_ErrorRemovalFromStack() should report an error, but it haven't")
			}
		})

		err1 := stack.PopUsecaseError()
		t.Run("Check if err1 is correct", func(t *testing.T) {
			if err1 != mockUsecaseError2 {
				t.Errorf("Test_ErrorStack_ErrorRemovalFromStack() returned '%s', when it should return '%s'", err1.Error(), mockUsecaseError2.Error())
			}
		})

		t.Run("Check if has decreased error stack (err1)", func(t *testing.T) {
			if len(stack.errors) > 1 {
				t.Errorf("Test_ErrorStack_ErrorRemovalFromStack() didn't decreased error stack")
			}
		})

		t.Run("Check if has removed error from stack (err1)", func(t *testing.T) {
			if _, exists := stack.errors["usecase_2"]; exists {
				t.Errorf("Test_ErrorStack_ErrorRemovalFromStack() didn't removed error from stack")
			}
		})

		err2 := stack.PopUsecaseError()
		t.Run("Check if err2 is correct", func(t *testing.T) {
			if err2 != mockUsecaseError {
				t.Errorf("Test_ErrorStack_ErrorRemovalFromStack() returned '%s', when it should return '%s'", err2.Error(), mockUsecaseError.Error())
			}
		})

		t.Run("Check if has decreased error stack (err2)", func(t *testing.T) {
			if len(stack.errors) > 1 {
				t.Errorf("Test_ErrorStack_ErrorRemovalFromStack() didn't decreased error stack")
			}
		})

		t.Run("Check if has removed error from stack (err2)", func(t *testing.T) {
			if _, exists := stack.errors["usecase_1"]; exists {
				t.Errorf("Test_ErrorStack_ErrorRemovalFromStack() didn't removed error from stack")
			}
		})
	}

	Test_ErrorStack_SuccessfulPop(t)
	Test_ErrorStack_MultipleSuccessfulPop(t)
	Test_ErrorStack_ErrorRemovalFromStack(t)
}

func Test_ErrorStack_AppendAndRetrieve(t *testing.T) {
	Test_ErrorStack_ReturnedCorrectLastError_UsecaseOnly := func(t *testing.T) {
		stack := Error{}
		stack.AppendUsecaseError(mockUsecaseError2)
		stack.AppendUsecaseError(mockUsecaseError)

		err := stack.LastUsecaseError()
		t.Run("Check if last error is correct", func(t *testing.T) {
			if err != mockUsecaseError {
				t.Errorf("Test_ErrorStack_ReturnedCorrectLastError_UsecaseOnly() returned '%s', when it should return '%s'", err.Error(), mockUsecaseError.Error())
			}
		})
	}

	Test_ErrorStack_ReturnedCorrectLastError_EntityOnly := func(t *testing.T) {
		stack := Error{}
		stack.AddEntityError(mockEntityErr)

		err := stack.LastUsecaseError()
		t.Run("Check if last error is correct", func(t *testing.T) {
			if err != mockEntityErr {
				t.Errorf("Test_ErrorStack_ReturnedCorrectLastError_EntityOnly() returned '%s', when it should return '%s'", err.Error(), mockEntityErr.Error())
			}
		})
	}

	Test_ErrorStack_ReturnedCorrectLastError_EntityAndUsecase := func(t *testing.T) {
		stack := Error{}
		stack.AddEntityError(mockEntityErr)
		stack.AppendUsecaseError(mockUsecaseError2)

		err := stack.LastUsecaseError()
		t.Run("Check if last error is correct", func(t *testing.T) {
			if err != mockUsecaseError2 {
				t.Errorf("Test_ErrorStack_ReturnedCorrectLastError_EntityAndUsecase() returned '%s', when it should return '%s'", err.Error(), mockUsecaseError2.Error())
			}
		})
	}

	Test_ErrorStack_ReturnedCorrectEntityError := func(t *testing.T) {
		stack := Error{}
		stack.AddEntityError(mockUsecaseError2)
		stack.AppendUsecaseError(mockEntityErr)

		err := stack.EntityError()
		t.Run("Check if entity error is correct", func(t *testing.T) {
			if err != mockUsecaseError2 {
				t.Errorf("Test_ErrorStack_ReturnedCorrectEntityError() returned '%s', when it should return '%s'", err.Error(), mockUsecaseError2.Error())
			}
		})
	}

	Test_ErrorStack_ReturnedCorrectNilEntityError := func(t *testing.T) {
		stack := Error{}
		stack.AppendUsecaseError(mockUsecaseError)

		err := stack.EntityError()
		t.Run("Check if entity error is correct (nil)", func(t *testing.T) {
			if err != nil {
				t.Errorf("Test_ErrorStack_ReturnedCorrectNilEntityError() returned '%s', when it should return 'nil'", err.Error())
			}
		})
	}

	Test_ErrorStack_ShouldntAppendMoreThanOneEntityError := func(t *testing.T) {
		stack := Error{}
		stack.AddEntityError(mockEntityErr)
		stack.AddEntityError(mockUsecaseError)

		err := stack.EntityError()
		t.Run("Check if entity error is correct (EntityError())", func(t *testing.T) {
			if err != mockEntityErr {
				t.Errorf("Test_ErrorStack_ShouldntAppendMoreThanOneEntityError() returned '%s', when it should return '%s'", err.Error(), mockEntityErr.Error())
			}
		})

		err = stack.LastUsecaseError()
		t.Run("Check if entity error is correct (LastUsecaseError())", func(t *testing.T) {
			if err != mockEntityErr {
				t.Errorf("Test_ErrorStack_ShouldntAppendMoreThanOneEntityError() returned '%s', when it should return '%s'", err.Error(), mockEntityErr.Error())
			}
		})
	}

	Test_ErrorStack_ReturnedCorrectLastError_UsecaseOnly(t)
	Test_ErrorStack_ReturnedCorrectLastError_EntityOnly(t)
	Test_ErrorStack_ReturnedCorrectLastError_EntityAndUsecase(t)
	Test_ErrorStack_ReturnedCorrectEntityError(t)
	Test_ErrorStack_ReturnedCorrectNilEntityError(t)
	Test_ErrorStack_ShouldntAppendMoreThanOneEntityError(t)
}
