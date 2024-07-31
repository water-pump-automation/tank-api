package get_tank

// import (
// 	"errors"
// 	"testing"
// 	"tank-api/app/entity/tank"
// 	"tank-api/app/usecases/ports"
// 	database_mock "tank-api/infra/database/mock"
// )

// var successGetTank = NewGetTank(database_mock.NewTankMockData())
// var failGetTank = NewGetTank(database_mock.NewTankFailMockData())

// func responsesAreEqual(expected, got *ports.TankState) (string, string, bool) {
// 	if expected.Name != got.Name {
// 		return "Name", got.Name, false
// 	}

// 	if expected.Group != got.Group {
// 		return "Group", got.Group, false
// 	}

// 	if expected.MaximumCapacity != got.MaximumCapacity {
// 		return "MaximumCapacity", got.MaximumCapacity, false
// 	}

// 	if expected.TankState != got.TankState {
// 		return "TankState", got.TankState, false
// 	}

// 	if expected.CurrentLevel != got.CurrentLevel {
// 		return "CurrentLevel", got.CurrentLevel, false
// 	}

// 	if expected.LastFullTime != got.LastFullTime {
// 		return "LastFullTime", got.LastFullTime.String(), false
// 	}

// 	return "", "", true
// }

// func Test_GetTank_Get(t *testing.T) {
// 	t.Run("Succesful tank  tank", func(t *testing.T) {
// 		expectedReturn := &ports.TankState{
// 			Name:              "TANK_1",
// 			Group:             "GROUP_1",
// 			MaximumCapacity:   "100.00L",
// 			TankState:         "EMPTY",
// 			CurrentLevel: "0.00L",
// 			LastFullTime:      database_mock.MockTimeNow,
// 		}

// 		state, err := successGetTank.Get("TANK_1", "GROUP_1")

// 		if err != nil {
// 			t.Error("Test_GetTank_Get() shouldn't report an error, but it have")
// 		}

// 		if field, value, equal := responsesAreEqual(expectedReturn, state); !equal {
// 			t.Errorf("Test_GetTank_Get() wrong '%s' field, got '%s'", field, value)
// 		}
// 	})

// 	t.Run("Not found tank  tank", func(t *testing.T) {
// 		_, err := successGetTank.Get("TANK_134256", "GROUP_1")

// 		if !err != nil {
// 			t.Error("Test_GetTank_Get() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorNotFound("TANK_134256").Error() {
// 			t.Errorf("Test_GetTank_Get() wrong error. Should return '%s', got '%s'", ErrTankErrorNotFound("TANK_134256"), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Internal server error getting  tank", func(t *testing.T) {
// 		_, err := failGetTank.Get("TANK_1", "GROUP_1")

// 		if !err != nil {
// 			t.Error("Test_GetTank_Get() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorServerError(errors.New("ERROR").Error()).Error() {
// 			t.Errorf("Test_GetTank_Get() wrong error. Should return '%s', got '%s'", ErrTankErrorServerError(errors.New("ERROR").Error()), err.LastUsecaseError())
// 		}
// 	})
// }

// func Test_GetTank_GetData(t *testing.T) {
// 	t.Run("Succesful tank  tank", func(t *testing.T) {
// 		expectedCapacity := tank.MaximumCapacity(100)

// 		capacity, err := successGetTank.GetMaximumCapacity("TANK_1", "GROUP_1")

// 		if err != nil {
// 			t.Error("Test_GetTank_GetData() shouldn't report an error, but it have")
// 		}

// 		if expectedCapacity != capacity {
// 			t.Errorf("Test_GetTank_GetData() wrong 'capacity' field. Expected '%f', got '%f'", expectedCapacity, capacity)
// 		}
// 	})

// 	t.Run("Not found tank  tank", func(t *testing.T) {
// 		_, err := successGetTank.GetMaximumCapacity("TANK_134256", "GROUP_1")

// 		if !err != nil {
// 			t.Error("Test_GetTank_GetData() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorNotFound("TANK_134256").Error() {
// 			t.Errorf("Test_GetTank_GetData() wrong error. Should return '%s', got '%s'", ErrTankErrorNotFound("TANK_134256"), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Internal server error getting  tank", func(t *testing.T) {
// 		_, err := failGetTank.GetMaximumCapacity("TANK_1", "GROUP_1")

// 		if !err != nil {
// 			t.Error("Test_GetTank_GetData() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorServerError(errors.New("ERROR").Error()).Error() {
// 			t.Errorf("Test_GetTank_Get() wrong error. Should return '%s', got '%s'", ErrTankErrorServerError(errors.New("ERROR").Error()), err.LastUsecaseError())
// 		}
// 	})
// }
