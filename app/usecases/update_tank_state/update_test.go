package update_tank_state

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"
// 	stack "tank-api/app/entity/error_stack"
// 	"tank-api/app/entity/tank"
// 	"tank-api/app/usecases/get_tank"
// 	database_mock "tank-api/infra/database/mock"
// )

// var successGetTank = get_tank.NewGetTank(database_mock.NewTankMockData())

// type tankUpdateMockData struct {
// 	states map[string]map[string]*tank.Tank
// }

// var updateMockDatabase = NewTankUpdateMockData()
// var MockTimeNow = time.Now()

// func NewTankUpdateMockData() *tankUpdateMockData {
// 	return &tankUpdateMockData{
// 		states: map[string]map[string]*tank.Tank{
// 			"GROUP_1": {
// 				"TANK_1": {
// 					Name:              "TANK_1",
// 					Group:             "GROUP_1",
// 					MaximumCapacity:   100,
// 					TankState:         tank.Empty,
// 					CurrentLevel: 0,
// 					LastFullTime:      MockTimeNow,
// 				},
// 			},
// 		},
// 	}
// }

// var successUpdateTank = NewTankUpdate(updateMockDatabase, successGetTank)
// var failUpdateTank = NewTankUpdate(database_mock.NewTankFailMockData(), successGetTank)

// func Test_Tank_Update(t *testing.T) {
// 	t.Run("Succesful update  tank (filling)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 80)

// 		if err != nil {
// 			t.Error("Test_Tank_Update() shouldn't report an error, but it have")
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != tank.Filling {
// 			t.Errorf("Test_Tank_Update() wrong TankState. Expected '%d', set '%d'", tank.Filling, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentLevel != 80 {
// 			t.Errorf("Test_Tank_Update() wrong CurrentLevel. Expected '%d', set '%f'", 80, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentLevel)
// 		}
// 	})

// 	t.Run("Succesful update  tank (full)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 100)

// 		if err != nil {
// 			t.Error("Test_Tank_Update() shouldn't report an error, but it have")
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != tank.Full {
// 			t.Errorf("Test_Tank_Update() wrong TankState. Expected '%d', set '%d'", tank.Full, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentLevel != 100 {
// 			t.Errorf("Test_Tank_Update() wrong CurrentLevel. Expected '%d', set '%f'", 100, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentLevel)
// 		}
// 	})

// 	t.Run("Succesful update  tank (empty)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 0)

// 		if err != nil {
// 			t.Error("Test_Tank_Update() shouldn't report an error, but it have")
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != tank.Empty {
// 			t.Errorf("Test_Tank_Update() wrong TankState. Expected '%d', set '%d'", tank.Empty, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
// 		}

// 		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentLevel != 0 {
// 			t.Errorf("Test_Tank_Update() wrong CurrentLevel. Expected '%d', set '%f'", 0, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentLevel)
// 		}
// 	})

// 	t.Run(" tank not found", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_135sb3", "GROUP_1", 100)

// 		if !err != nil {
// 			t.Error("Test_Tank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorNotFound("TANK_135sb3").Error() {
// 			t.Errorf("Test_Tank_Update() wrong error. Should return '%s', got '%s'", ErrTankErrorNotFound("TANK_135sb3").Error(), err.LastUsecaseError().Error())
// 		}
// 	})

// 	t.Run("Tank invalid  level (smaller than 0)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", -1)

// 		if !err != nil {
// 			t.Error("Test_Tank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankCurrentLevelSmallerThanZero {
// 			t.Errorf("Test_Tank_Update() wrong error. Should return '%s', got '%s'", ErrTankCurrentLevelSmallerThanZero.Error(), err.LastUsecaseError().Error())
// 		}
// 	})

// 	t.Run("Tank invalid  level (bigger thank maximum capacity)", func(t *testing.T) {
// 		err := successUpdateTank.Update("TANK_1", "GROUP_1", 101)

// 		if !err != nil {
// 			t.Error("Test_Tank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError() != ErrTankCurrentLevelBiggerThanMax {
// 			t.Errorf("Test_Tank_Update() wrong error. Should return '%s', got '%s'", ErrTankCurrentLevelBiggerThanMax.Error(), err.LastUsecaseError().Error())
// 		}
// 	})

// 	t.Run("Internal server error updating  tank", func(t *testing.T) {
// 		err := failUpdateTank.Update("TANK_1", "GROUP_1", 100)

// 		if !err != nil {
// 			t.Error("Test_Tank_Update() should report an error, but it haven't")
// 		}

// 		if err.LastUsecaseError().Error() != ErrTankErrorServerError(errors.New("ERROR").Error()).Error() {
// 			t.Errorf("Test_Tank_Update() wrong error. Should return '%s', got '%s'", ErrTankErrorServerError(errors.New("ERROR").Error()), err.LastUsecaseError())
// 		}
// 	})
// }

// func (tank *tankUpdateMockData) GetTankState(ctx context.Context,   input *tank.GetTankState) (state *tank.Tank, err error) {
// 	state = tank.states[group][names[0]]
// 	return
// }

// func (tank *tankUpdateMockData) GetTankGroupState(ctx context.Context,   input *tank.GetGroupTanksInput) (state []*tank.Tank, err error) {
// 	if group, exists := tank.states[groups[0]]; exists {
// 		for _, tank := range group {
// 			state = append(state, tank)
// 		}
// 	}
// 	return
// }

// func (tank *tankUpdateMockData) CreateTank(ctx context.Context,   input *tank.CreateInput) (tankState *tank.Tank, err error) {
// 	if _, exists := tank.states[group]; !exists {
// 		tank.states[group] = map[string]*tank.Tank{
// 			name: {
// 				Name:              name,
// 				Group:             group,
// 				MaximumCapacity:   capacity,
// 				TankState:         tank.Empty,
// 				CurrentLevel: 0,
// 			},
// 		}
// 		return
// 	}

// 	tank.states[group][name] = &tank.Tank{
// 		Name:              name,
// 		Group:             group,
// 		MaximumCapacity:   capacity,
// 		TankState:         tank.Empty,
// 		CurrentLevel: 0,
// 	}

// 	return
// }

// func (tank *tankUpdateMockData) UpdateTankLevel(ctx context.Context,   input *tank.UpdateLevelInput) (state *tank.Tank, err error) {
// 	if group, exists := tank.states[group]; exists {
// 		group[name].CurrentLevel = level
// 		group[name].TankState = levelState
// 	}
// 	return
// }
