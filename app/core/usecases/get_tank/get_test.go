package get_tank

import (
	"errors"
	"testing"
	"water-tank-api/app/core/entity/access"
	"water-tank-api/app/core/entity/water_tank"
	database_mock "water-tank-api/app/infra/database/mock"
)

var successGetWaterTank = NewGetWaterTank(database_mock.NewWaterTankMockData())
var failGetWaterTank = NewGetWaterTank(database_mock.NewWaterTankFailMockData())

func responsesAreEqual(expected, got *water_tank.WaterTankState) (string, string, bool) {
	if expected.Name != got.Name {
		return "Name", got.Name, false
	}

	if expected.Group != got.Group {
		return "Group", got.Group, false
	}

	if expected.MaximumCapacity != got.MaximumCapacity {
		return "MaximumCapacity", got.MaximumCapacity, false
	}

	if expected.TankState != got.TankState {
		return "TankState", got.TankState, false
	}

	if expected.CurrentWaterLevel != got.CurrentWaterLevel {
		return "CurrentWaterLevel", got.CurrentWaterLevel, false
	}

	if expected.LastFullTime != got.LastFullTime {
		return "LastFullTime", got.LastFullTime.String(), false
	}

	return "", "", true
}

func Test_GetWaterTank_Get(t *testing.T) {
	t.Run("Succesful water_tank water tank", func(t *testing.T) {
		expectedReturn := &water_tank.WaterTankState{
			Name:              "TANK_1",
			Group:             "GROUP_1",
			MaximumCapacity:   "100.00L",
			TankState:         "EMPTY",
			CurrentWaterLevel: "0.00L",
			LastFullTime:      database_mock.MockTimeNow,
		}

		state, err := successGetWaterTank.Get("TANK_1", "GROUP_1")

		if err.HasError() {
			t.Error("Test_GetWaterTank_Get() shouldn't report an error, but it have")
		}

		if field, value, equal := responsesAreEqual(expectedReturn, state); !equal {
			t.Errorf("Test_GetWaterTank_Get() wrong '%s' field, got '%s'", field, value)
		}
	})

	t.Run("Not found water_tank water tank", func(t *testing.T) {
		_, err := successGetWaterTank.Get("TANK_134256", "GROUP_1")

		if !err.HasError() {
			t.Error("Test_GetWaterTank_Get() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorNotFound("TANK_134256").Error() {
			t.Errorf("Test_GetWaterTank_Get() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorNotFound("TANK_134256"), err.LastError())
		}
	})

	t.Run("Internal server error getting water tank", func(t *testing.T) {
		_, err := failGetWaterTank.Get("TANK_1", "GROUP_1")

		if !err.HasError() {
			t.Error("Test_GetWaterTank_Get() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
			t.Errorf("Test_GetWaterTank_Get() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorServerError(errors.New("ERROR").Error()), err.LastError())
		}
	})
}

func Test_GetWaterTank_GetData(t *testing.T) {
	t.Run("Succesful water_tank water tank", func(t *testing.T) {
		expectedCapacity := water_tank.Capacity(100)
		expectedAccess := access.AccessToken("a")

		capacity, access, err := successGetWaterTank.GetData("TANK_1", "GROUP_1")

		if err.HasError() {
			t.Error("Test_GetWaterTank_GetData() shouldn't report an error, but it have")
		}

		if expectedAccess != access {
			t.Errorf("Test_GetWaterTank_GetData() wrong 'access' field. Expected '%s', got '%s'", expectedAccess, access)
		}

		if expectedCapacity != capacity {
			t.Errorf("Test_GetWaterTank_GetData() wrong 'capacity' field. Expected '%f', got '%f'", expectedCapacity, capacity)
		}
	})

	t.Run("Not found water_tank water tank", func(t *testing.T) {
		_, _, err := successGetWaterTank.GetData("TANK_134256", "GROUP_1")

		if !err.HasError() {
			t.Error("Test_GetWaterTank_GetData() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorNotFound("TANK_134256").Error() {
			t.Errorf("Test_GetWaterTank_GetData() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorNotFound("TANK_134256"), err.LastError())
		}
	})

	t.Run("Internal server error getting water tank", func(t *testing.T) {
		_, _, err := failGetWaterTank.GetData("TANK_1", "GROUP_1")

		if !err.HasError() {
			t.Error("Test_GetWaterTank_GetData() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
			t.Errorf("Test_GetWaterTank_Get() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorServerError(errors.New("ERROR").Error()), err.LastError())
		}
	})
}
