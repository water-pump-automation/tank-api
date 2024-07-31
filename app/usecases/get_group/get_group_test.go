package get_group

// import (
// 	"errors"
// 	"testing"
// 	"tank-api/app/usecases/ports"
// 	database_mock "tank-api/infra/database/mock"
// )

// var successGetTank = NewGetGroupTank(database_mock.NewTankMockData())
// var failGetTank = NewGetGroupTank(database_mock.NewTankFailMockData())

// func existsInResponse(expected, got *ports.TankGroupState) (string, bool) {
// 	states := map[string]*ports.TankState{}

// 	for _, gotState := range got.Tanks {
// 		states[gotState.Name] = gotState
// 	}

// 	for _, state := range expected.Tanks {
// 		if _, exists := states[state.Name]; !exists {
// 			return state.Name, false
// 		}
// 	}

// 	return "", true
// }

// func responsesAreEqual(expected, got *ports.TankGroupState) (string, string, bool) {
// 	states := map[string]*ports.TankState{}

// 	for _, gotState := range got.Tanks {
// 		states[gotState.Name] = gotState
// 	}

// 	for _, state := range expected.Tanks {
// 		if gotState, exists := states[state.Name]; exists {
// 			if state.Name != gotState.Name {
// 				return "Name", gotState.Name, false
// 			}

// 			if state.Group != gotState.Group {
// 				return "Group", gotState.Group, false
// 			}

// 			if state.MaximumCapacity != gotState.MaximumCapacity {
// 				return "MaximumCapacity", gotState.MaximumCapacity, false
// 			}

// 			if state.TankState != gotState.TankState {
// 				return "TankState", gotState.TankState, false
// 			}

// 			if state.CurrentLevel != gotState.CurrentLevel {
// 				return "CurrentLevel", gotState.CurrentLevel, false
// 			}

// 			if state.LastFullTime != gotState.LastFullTime {
// 				return "LastFullTime", gotState.LastFullTime.String(), false
// 			}
// 		}
// 	}

// 	return "", "", true
// }

// func Test_GetGroupTank_Get(t *testing.T) {
// 	t.Run("Succesful data  tank group", func(t *testing.T) {
// 		expectedReturn := &ports.TankGroupState{
// 			Tanks: []*ports.TankState{
// 				{
// 					Name:              "TANK_1",
// 					Group:             "GROUP_1",
// 					MaximumCapacity:   "100.00L",
// 					TankState:         "EMPTY",
// 					CurrentLevel: "0.00L",
// 					LastFullTime:      database_mock.MockTimeNow,
// 				},
// 				{
// 					Name:              "TANK_2",
// 					Group:             "GROUP_1",
// 					MaximumCapacity:   "80.00L",
// 					TankState:         "FILLING",
// 					CurrentLevel: "50.00L",
// 					LastFullTime:      database_mock.MockTimeNow,
// 				},
// 				{
// 					Name:              "TANK_3",
// 					Group:             "GROUP_1",
// 					MaximumCapacity:   "120.00L",
// 					TankState:         "FULL",
// 					CurrentLevel: "120.00L",
// 					LastFullTime:      database_mock.MockTimeNow,
// 				},
// 			},
// 		}

// 		state, err := successGetTank.Get("GROUP_1")

// 		if err != nil {
// 			t.Error("Test_GetGroupTank_Get() shouldn't report an error, but it have")
// 		}

// 		if tank, equal := existsInResponse(expectedReturn, state); !equal {
// 			t.Errorf("Test_GetGroupTank_Get() didn't foud '%s' tank", tank)
// 		}

// 		if field, value, equal := responsesAreEqual(expectedReturn, state); !equal {
// 			t.Errorf("Test_GetGroupTank_Get() wrong '%s' field, got '%s'", field, value)
// 		}
// 	})

// 	t.Run("Not found data  tank group", func(t *testing.T) {
// 		_, err := successGetTank.Get("GROUP_123532")

// 		if !err != nil {
// 			t.Error("Test_GetGroupTank_Get() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorGroupNotFound("GROUP_123532").Error() {
// 			t.Errorf("Test_GetGroupTank_Get() wrong error. Should return '%s', got '%s'", ErrTankErrorGroupNotFound("TANK_134256"), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Invalid name in data  tank group", func(t *testing.T) {
// 		_, err := successGetTank.Get("")

// 		if !err != nil {
// 			t.Error("Test_GetGroupTank_Get() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankMissingGroup {
// 			t.Errorf("Test_GetGroupTank_Get() wrong error. Should return '%s', got '%s'", ErrTankMissingGroup.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Internal server error getting  tank group", func(t *testing.T) {
// 		_, err := failGetTank.Get("GROUP_1")

// 		if !err != nil {
// 			t.Error("Test_GetGroupTank_Get() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorServerError(errors.New("ERROR").Error()).Error() {
// 			t.Errorf("Test_GetGroupTank_Get() wrong error. Should return '%s', got '%s'", ErrTankErrorServerError(errors.New("ERROR").Error()), err.LastUsecaseError())
// 		}
// 	})
// }
