package create_tank

// import (
// 	"errors"
// 	"testing"
// 	"tank-api/app/usecases/get_tank"
// 	database_mock "tank-api/infra/database/mock"
// )

// var successGetTank = get_tank.NewGetTank(database_mock.NewTankMockData())

// var successCreateTank = NewTank(database_mock.NewTankMockData(), successGetTank)
// var failCreateTank = NewTank(database_mock.NewTankFailMockData(), successGetTank)

// func Test_Tank_Create(t *testing.T) {
// 	t.Run("Succesful create  tank", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_235", "GROUP_1", 100)

// 		if err != nil {
// 			t.Error("Test_Tank_Create() shouldn't report an error, but it have")
// 		}
// 	})

// 	t.Run(" tank already exists", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_1", "GROUP_1", 100)

// 		if !err != nil {
// 			t.Error("Test_Tank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankAlreadyExists {
// 			t.Errorf("Test_Tank_Create() wrong error. Should return '%s', got '%s'", ErrTankAlreadyExists.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run(" tank invalid maximum capacity (equal 0)", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_236", "GROUP_1", 0)

// 		if !err != nil {
// 			t.Error("Test_Tank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankMaximumCapacityZero {
// 			t.Errorf("Test_Tank_Create() wrong error. Should return '%s', got '%s'", ErrTankMaximumCapacityZero.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run(" tank invalid maximum capacity (smaller than 0)", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_237", "GROUP_1", -1)

// 		if !err != nil {
// 			t.Error("Test_Tank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankMaximumCapacityZero {
// 			t.Errorf("Test_Tank_Create() wrong error. Should return '%s', got '%s'", ErrTankMaximumCapacityZero.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run(" tank invalid name", func(t *testing.T) {
// 		_, err := successCreateTank.Create("", "GROUP_1", 10)

// 		if !err != nil {
// 			t.Error("Test_Tank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankInvalidName {
// 			t.Errorf("Test_Tank_Create() wrong error. Should return '%s', got '%s'", ErrTankInvalidName.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run(" tank invalid group", func(t *testing.T) {
// 		_, err := successCreateTank.Create("TANK_364", "", 10)

// 		if !err != nil {
// 			t.Error("Test_Tank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankInvalidGroup {
// 			t.Errorf("Test_Tank_Create() wrong error. Should return '%s', got '%s'", ErrTankInvalidGroup.Error(), err.LastUsecaseError())
// 		}
// 	})

// 	t.Run("Internal server error updating  tank", func(t *testing.T) {
// 		_, err := failCreateTank.Create("TANK_235", "GROUP_1", 100)

// 		if !err != nil {
// 			t.Error("Test_Tank_Create() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorServerError(errors.New("ERROR").Error()).Error() {
// 			t.Errorf("Test_Tank_Create() wrong error. Should return '%s', got '%s'", ErrTankErrorServerError(errors.New("ERROR").Error()), err.LastUsecaseError())
// 		}
// 	})
// }
