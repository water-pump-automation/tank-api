package create_tank

// import (
// 	"errors"
// 	"testing"
// 	"water-tank-api/app/core/usecases/get_tank"
// 	database_mock "water-tank-api/infra/database/mock"
// )

// var successGetWaterTank = get_tank.NewGetWaterTank(database_mock.NewWaterTankMockData())

// var successCreateTank = NewWaterTank(database_mock.NewWaterTankMockData(), successGetWaterTank)
// var failCreateTank = NewWaterTank(database_mock.NewWaterTankFailMockData(), successGetWaterTank)

// func Test_WaterTank_Create(t *testing.T) {
// 	t.Run("Succesful create water tank", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_235", "GROUP_1", 100)

// 		if err.HasError() {
// 			t.Error("Test_WaterTank_Create() shouldn't report an error, but it have")
// 		}
// 	})

// 	t.Run("Water tank already exists", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_1", "GROUP_1", 100)

// 		if !err.HasError() {
// 			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrWaterTankAlreadyExists {
// 			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", ErrWaterTankAlreadyExists.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Water tank invalid maximum capacity (equal 0)", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_236", "GROUP_1", 0)

// 		if !err.HasError() {
// 			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrWaterTankMaximumCapacityZero {
// 			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", ErrWaterTankMaximumCapacityZero.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Water tank invalid maximum capacity (smaller than 0)", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_237", "GROUP_1", -1)

// 		if !err.HasError() {
// 			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrWaterTankMaximumCapacityZero {
// 			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", ErrWaterTankMaximumCapacityZero.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Water tank invalid name", func(t *testing.T) {
// 		_, err := successCreateTank.Create("", "GROUP_1", 10)

// 		if !err.HasError() {
// 			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrWaterTankInvalidName {
// 			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", ErrWaterTankInvalidName.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Water tank invalid group", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_364", "", 10)

// 		if !err.HasError() {
// 			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrWaterTankInvalidGroup {
// 			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", ErrWaterTankInvalidGroup.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Internal server error updating water tank", func(t *testing.T) {
// 		_, err := failCreateTank.Create("TANK_235", "GROUP_1", 100)

// 		if !err.HasError() {
// 			t.Error("Test_WaterTank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
// 			t.Errorf("Test_WaterTank_Create() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorServerError(errors.New("ERROR").Error()), err.LastUsecaseError())
// 		}
// 	})
// }
