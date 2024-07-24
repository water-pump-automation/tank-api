package get_group

import (
	"errors"
	"testing"
	"water-tank-api/app/core/usecases/ports"
	database_mock "water-tank-api/infra/database/mock"
)

var successGetWaterTank = NewGetGroupWaterTank(database_mock.NewWaterTankMockData())
var failGetWaterTank = NewGetGroupWaterTank(database_mock.NewWaterTankFailMockData())

func existsInResponse(expected, got *ports.WaterTankGroupState) (string, bool) {
	states := map[string]*ports.WaterTankState{}

	for _, gotState := range got.Tanks {
		states[gotState.Name] = gotState
	}

	for _, state := range expected.Tanks {
		if _, exists := states[state.Name]; !exists {
			return state.Name, false
		}
	}

	return "", true
}

func responsesAreEqual(expected, got *ports.WaterTankGroupState) (string, string, bool) {
	states := map[string]*ports.WaterTankState{}

	for _, gotState := range got.Tanks {
		states[gotState.Name] = gotState
	}

	for _, state := range expected.Tanks {
		if gotState, exists := states[state.Name]; exists {
			if state.Name != gotState.Name {
				return "Name", gotState.Name, false
			}

			if state.Group != gotState.Group {
				return "Group", gotState.Group, false
			}

			if state.MaximumCapacity != gotState.MaximumCapacity {
				return "MaximumCapacity", gotState.MaximumCapacity, false
			}

			if state.TankState != gotState.TankState {
				return "TankState", gotState.TankState, false
			}

			if state.CurrentWaterLevel != gotState.CurrentWaterLevel {
				return "CurrentWaterLevel", gotState.CurrentWaterLevel, false
			}

			if state.LastFullTime != gotState.LastFullTime {
				return "LastFullTime", gotState.LastFullTime.String(), false
			}
		}
	}

	return "", "", true
}

func Test_GetGroupWaterTank_Get(t *testing.T) {
	t.Run("Succesful data water tank group", func(t *testing.T) {
		expectedReturn := &ports.WaterTankGroupState{
			Tanks: []*ports.WaterTankState{
				{
					Name:              "TANK_1",
					Group:             "GROUP_1",
					MaximumCapacity:   "100.00L",
					TankState:         "EMPTY",
					CurrentWaterLevel: "0.00L",
					LastFullTime:      database_mock.MockTimeNow,
				},
				{
					Name:              "TANK_2",
					Group:             "GROUP_1",
					MaximumCapacity:   "80.00L",
					TankState:         "FILLING",
					CurrentWaterLevel: "50.00L",
					LastFullTime:      database_mock.MockTimeNow,
				},
				{
					Name:              "TANK_3",
					Group:             "GROUP_1",
					MaximumCapacity:   "120.00L",
					TankState:         "FULL",
					CurrentWaterLevel: "120.00L",
					LastFullTime:      database_mock.MockTimeNow,
				},
			},
		}

		state, err := successGetWaterTank.Get("GROUP_1")

		if err.HasError() {
			t.Error("Test_GetGroupWaterTank_Get() shouldn't report an error, but it have")
		}

		if tank, equal := existsInResponse(expectedReturn, state); !equal {
			t.Errorf("Test_GetGroupWaterTank_Get() didn't foud '%s' tank", tank)
		}

		if field, value, equal := responsesAreEqual(expectedReturn, state); !equal {
			t.Errorf("Test_GetGroupWaterTank_Get() wrong '%s' field, got '%s'", field, value)
		}
	})

	t.Run("Not found data water tank group", func(t *testing.T) {
		_, err := successGetWaterTank.Get("GROUP_123532")

		if !err.HasError() {
			t.Error("Test_GetGroupWaterTank_Get() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorGroupNotFound("GROUP_123532").Error() {
			t.Errorf("Test_GetGroupWaterTank_Get() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorGroupNotFound("TANK_134256"), err.LastError())
		}
	})

	t.Run("Invalid name in data water tank group", func(t *testing.T) {
		_, err := successGetWaterTank.Get("")

		if !err.HasError() {
			t.Error("Test_GetGroupWaterTank_Get() should report an error, but it haven't")
		}

		if err.LastError() != ErrWaterTankMissingGroup {
			t.Errorf("Test_GetGroupWaterTank_Get() wrong error. Should return '%s', got '%s'", ErrWaterTankMissingGroup.Error(), err.LastError())
		}
	})

	t.Run("Internal server error getting water tank group", func(t *testing.T) {
		_, err := failGetWaterTank.Get("GROUP_1")

		if !err.HasError() {
			t.Error("Test_GetGroupWaterTank_Get() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
			t.Errorf("Test_GetGroupWaterTank_Get() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorServerError(errors.New("ERROR").Error()), err.LastError())
		}
	})
}
