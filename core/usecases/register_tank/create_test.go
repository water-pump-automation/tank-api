package register_tank

import (
	"errors"
	"testing"
	database_mock "water-tank-api/infra/database/mock"
)

var successCreateTank = NewWaterTank(database_mock.NewWaterTankMockData())
var failCreateTank = NewWaterTank(database_mock.NewWaterTankFailMockData())

func Test_WaterTank_Create(t *testing.T) {
	t.Run("Succesful create water tank", func(t *testing.T) {
		_, err := successCreateTank.Create("TANK_235", "GROUP_1", 100)

		if err.HasError() {
			t.Error("Test_WaterTank_Create() shouldn't report an error, but it have")
		}
	})

	t.Run("Water tank already exists", func(t *testing.T) {
		_, err := successCreateTank.Create("TANK_1", "GROUP_1", 100)

		if !err.HasError() {
			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankAlreadyExists {
			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", WaterTankAlreadyExists.Error(), err.LastError())
		}
	})

	t.Run("Water tank invalid maximum capacity (equal 0)", func(t *testing.T) {
		_, err := successCreateTank.Create("TANK_236", "GROUP_1", 0)

		if !err.HasError() {
			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankMaximumCapacityZero {
			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", WaterTankMaximumCapacityZero.Error(), err.LastError())
		}
	})

	t.Run("Water tank invalid maximum capacity (smaller than 0)", func(t *testing.T) {
		_, err := successCreateTank.Create("TANK_237", "GROUP_1", -1)

		if !err.HasError() {
			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankMaximumCapacityZero {
			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", WaterTankMaximumCapacityZero.Error(), err.LastError())
		}
	})

	t.Run("Water tank invalid name", func(t *testing.T) {
		_, err := successCreateTank.Create("", "GROUP_1", 10)

		if !err.HasError() {
			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankInvalidName {
			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", WaterTankInvalidName.Error(), err.LastError())
		}
	})

	t.Run("Water tank invalid group", func(t *testing.T) {
		_, err := successCreateTank.Create("TANK_364", "", 10)

		if !err.HasError() {
			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankInvalidGroup {
			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", WaterTankInvalidGroup.Error(), err.LastError())
		}
	})

	t.Run("Internal server error updating water tank", func(t *testing.T) {
		_, err := failCreateTank.Create("TANK_235", "GROUP_1", 100)

		if !err.HasError() {
			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
		}

		if err.LastError().Error() != WaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", WaterTankErrorServerError(errors.New("ERROR").Error()), err.LastError())
		}
	})
}
