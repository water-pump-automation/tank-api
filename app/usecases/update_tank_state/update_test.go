package update_tank_state

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"
// 	stack "water-tank-api/app/entity/error_stack"
// 	"water-tank-api/app/entity/water_tank"
// 	"water-tank-api/app/usecases/get_tank"
// 	database_mock "water-tank-api/infra/database/mock"
// )

// var successGetWaterTank = get_tank.NewGetWaterTank(database_mock.NewWaterTankMockData())

// type waterTankUpdateMockData struct {
// 	states map[string]map[string]*water_tank.WaterTank
// }

// var updateMockDatabase = NewWaterTankUpdateMockData()
// var MockTimeNow = time.Now()

// func NewWaterTankUpdateMockData() *waterTankUpdateMockData {
// 	return &waterTankUpdateMockData{
// 		states: map[string]map[string]*water_tank.WaterTank{
// 			"GROUP_1": {
// 				"TANK_1": {
// 					Name:              "TANK_1",
// 					Group:             "GROUP_1",
// 					MaximumCapacity:   100,
// 					TankState:         water_tank.Empty,
// 					CurrentWaterLevel: 0,
// 					LastFullTime:      MockTimeNow,
// 				},
// 			},
// 		},
// 	}
// }

// var successUpdateTank = NewWaterTankUpdate(updateMockDatabase, successGetWaterTank)
// var failUpdateTank = NewWaterTankUpdate(database_mock.NewWaterTankFailMockData(), successGetWaterTank)

// func Test_WaterTank_Update(t *testing.T) {
// 	t.Run("Succesful update water tank (filling)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 80)

// 		if err != nil {
// 			t.Error("Test_WaterTank_Update() shouldn't report an error, but it have")
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != water_tank.Filling {
// 			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", water_tank.Filling, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel != 80 {
// 			t.Errorf("Test_WaterTank_Update() wrong CurrentWaterLevel. Expected '%d', set '%f'", 80, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel)
// 		}
// 	})

// 	t.Run("Succesful update water tank (full)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 100)

// 		if err != nil {
// 			t.Error("Test_WaterTank_Update() shouldn't report an error, but it have")
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != water_tank.Full {
// 			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", water_tank.Full, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel != 100 {
// 			t.Errorf("Test_WaterTank_Update() wrong CurrentWaterLevel. Expected '%d', set '%f'", 100, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel)
// 		}
// 	})

// 	t.Run("Succesful update water tank (empty)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 0)

// 		if err != nil {
// 			t.Error("Test_WaterTank_Update() shouldn't report an error, but it have")
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != water_tank.Empty {
// 			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", water_tank.Empty, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel != 0 {
// 			t.Errorf("Test_WaterTank_Update() wrong CurrentWaterLevel. Expected '%d', set '%f'", 0, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel)
// 		}
// 	})

// 	t.Run("Water tank not found", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_135sb3", "GROUP_1", 100)

// 		if !err != nil {
// 			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrWaterTankErrorNotFound("TANK_135sb3").Error() {
// 			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorNotFound("TANK_135sb3").Error(), err.LastUsecaseError().Error())
// 		}
// 	})

// 	t.Run("Tank invalid water level (smaller than 0)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", -1)

// 		if !err != nil {
// 			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrWaterTankCurrentWaterLevelSmallerThanZero {
// 			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankCurrentWaterLevelSmallerThanZero.Error(), err.LastUsecaseError().Error())
// 		}
// 	})

// 	t.Run("Tank invalid water level (bigger thank maximum capacity)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 101)

// 		if !err != nil {
// 			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrWaterTankCurrentWaterLevelBiggerThanMax {
// 			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankCurrentWaterLevelBiggerThanMax.Error(), err.LastUsecaseError().Error())
// 		}
// 	})

// 	t.Run("Internal server error updating water tank", func(t *testing.T) {
// 		err := failUpdateTank.Update("TANK_1", "GROUP_1", 100)

// 		if !err != nil {
// 			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
// 			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorServerError(errors.New("ERROR").Error()), err.LastUsecaseError())
// 		}
// 	})
// }

// func (tank *waterTankUpdateMockData) GetWaterTankState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetWaterTankState) (state *water_tank.WaterTank, err error) {
// 	state = tank.states[group][names[0]]
// 	return
// }

// func (tank *waterTankUpdateMockData) GetTankGroupState(ctx context.Context, connection water_tank.IConn, input *water_tank.GetGroupTanks) (state []*water_tank.WaterTank, err error) {
// 	if group, exists := tank.states[groups[0]]; exists {
// 		for _, tank := range group {
// 			state = append(state, tank)
// 		}
// 	}
// 	return
// }

// func (tank *waterTankUpdateMockData) CreateWaterTank(ctx context.Context, connection water_tank.IConn, input *water_tank.CreateInput) (tankState *water_tank.WaterTank, err error) {
// 	if _, exists := tank.states[group]; !exists {
// 		tank.states[group] = map[string]*water_tank.WaterTank{
// 			name: {
// 				Name:              name,
// 				Group:             group,
// 				MaximumCapacity:   capacity,
// 				TankState:         water_tank.Empty,
// 				CurrentWaterLevel: 0,
// 			},
// 		}
// 		return
// 	}

// 	tank.states[group][name] = &water_tank.WaterTank{
// 		Name:              name,
// 		Group:             group,
// 		MaximumCapacity:   capacity,
// 		TankState:         water_tank.Empty,
// 		CurrentWaterLevel: 0,
// 	}

// 	return
// }

// func (tank *waterTankUpdateMockData) UpdateTankWaterLevel(ctx context.Context, connection water_tank.IConn, input *water_tank.UpdateWaterLevelInput) (state *water_tank.WaterTank, err error) {
// 	if group, exists := tank.states[group]; exists {
// 		group[name].CurrentWaterLevel = waterLevel
// 		group[name].TankState = levelState
// 	}
// 	return
// }
