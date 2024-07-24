package update_tank_state

import (
	"errors"
	"testing"
	"time"
	"water-tank-api/app/core/entity/access"
	stack "water-tank-api/app/core/entity/error_stack"
	"water-tank-api/app/core/entity/water_tank"
	"water-tank-api/app/core/usecases/get_tank"
	database_mock "water-tank-api/infra/database/mock"
)

var successGetWaterTank = get_tank.NewGetWaterTank(database_mock.NewWaterTankMockData())

type waterTankUpdateMockData struct {
	states map[string]map[string]*water_tank.WaterTank
}

var updateMockDatabase = NewWaterTankUpdateMockData()
var MockTimeNow = time.Now()

func NewWaterTankUpdateMockData() *waterTankUpdateMockData {
	return &waterTankUpdateMockData{
		states: map[string]map[string]*water_tank.WaterTank{
			"GROUP_1": {
				"TANK_1": {
					Name:              "TANK_1",
					Group:             "GROUP_1",
					MaximumCapacity:   100,
					TankState:         water_tank.Empty,
					CurrentWaterLevel: 0,
					Access:            "a",
					LastFullTime:      MockTimeNow,
				},
			},
		},
	}
}

var successUpdateTank = NewWaterTankUpdate(updateMockDatabase, successGetWaterTank)
var failUpdateTank = NewWaterTankUpdate(database_mock.NewWaterTankFailMockData(), successGetWaterTank)

func Test_WaterTank_Update(t *testing.T) {
	t.Run("Succesful update water tank (filling)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", 80)

		if err.HasError() {
			t.Error("Test_WaterTank_Update() shouldn't report an error, but it have")
		}

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != water_tank.Filling {
			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", water_tank.Filling, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
		}

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel != 80 {
			t.Errorf("Test_WaterTank_Update() wrong CurrentWaterLevel. Expected '%d', set '%f'", 80, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel)
		}
	})

	t.Run("Succesful update water tank (full)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", 100)

		if err.HasError() {
			t.Error("Test_WaterTank_Update() shouldn't report an error, but it have")
		}

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != water_tank.Full {
			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", water_tank.Full, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
		}

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel != 100 {
			t.Errorf("Test_WaterTank_Update() wrong CurrentWaterLevel. Expected '%d', set '%f'", 100, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel)
		}
	})

	t.Run("Succesful update water tank (empty)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", 0)

		if err.HasError() {
			t.Error("Test_WaterTank_Update() shouldn't report an error, but it have")
		}

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != water_tank.Empty {
			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", water_tank.Empty, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
		}

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel != 0 {
			t.Errorf("Test_WaterTank_Update() wrong CurrentWaterLevel. Expected '%d', set '%f'", 0, updateMockDatabase.states["GROUP_1"]["TANK_1"].CurrentWaterLevel)
		}
	})

	t.Run("Water tank not found", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_135sb3", "GROUP_1", "a", 100)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorNotFound("TANK_135sb3").Error() {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorNotFound("TANK_135sb3").Error(), err.LastError().Error())
		}
	})

	t.Run("Tank invalid water level (smaller than 0)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", -1)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError() != ErrWaterTankCurrentWaterLevelSmallerThanZero {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankCurrentWaterLevelSmallerThanZero.Error(), err.LastError().Error())
		}
	})

	t.Run("Tank invalid water level (bigger thank maximum capacity)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", 101)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError() != ErrWaterTankCurrentWaterLevelBiggerThanMax {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankCurrentWaterLevelBiggerThanMax.Error(), err.LastError().Error())
		}
	})

	t.Run("Tank invalid access token", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "b", 45)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError() != ErrWaterTankInvalidToken {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankInvalidToken.Error(), err.LastError().Error())
		}
	})

	t.Run("Internal server error updating water tank", func(t *testing.T) {
		err := failUpdateTank.Update("TANK_1", "GROUP_1", "a", 100)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError().Error() != ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", ErrWaterTankErrorServerError(errors.New("ERROR").Error()), err.LastError())
		}
	})
}

func (tank *waterTankUpdateMockData) GetWaterTankState(group string, names ...string) (state *water_tank.WaterTank, err stack.ErrorStack) {
	state = tank.states[group][names[0]]
	return
}

func (tank *waterTankUpdateMockData) GetTankGroupState(groups ...string) (state []*water_tank.WaterTank, err stack.ErrorStack) {
	if group, exists := tank.states[groups[0]]; exists {
		for _, tank := range group {
			state = append(state, tank)
		}
	}
	return
}

func (tank *waterTankUpdateMockData) NotifyFullTank(name string, group string, currentTime time.Time) (state *water_tank.WaterTank, err stack.ErrorStack) {
	if group, exists := tank.states[group]; exists {
		group[name].LastFullTime = currentTime
	}
	return
}

func (tank *waterTankUpdateMockData) CreateWaterTank(name string, group string, accessToken access.AccessToken, capacity water_tank.Capacity) (err stack.ErrorStack) {
	if _, exists := tank.states[group]; !exists {
		tank.states[group] = map[string]*water_tank.WaterTank{
			name: {
				Name:              name,
				Group:             group,
				MaximumCapacity:   capacity,
				TankState:         water_tank.Empty,
				CurrentWaterLevel: 0,
				Access:            accessToken,
			},
		}
		return
	}

	tank.states[group][name] = &water_tank.WaterTank{
		Name:              name,
		Group:             group,
		MaximumCapacity:   capacity,
		TankState:         water_tank.Empty,
		CurrentWaterLevel: 0,
		Access:            accessToken,
	}

	return
}

func (tank *waterTankUpdateMockData) UpdateWaterTankState(name string, group string, waterLevel water_tank.Capacity, levelState water_tank.State) (state *water_tank.WaterTank, err stack.ErrorStack) {
	if group, exists := tank.states[group]; exists {
		group[name].CurrentWaterLevel = waterLevel
		group[name].TankState = levelState
	}
	return
}
