package update_tank_state

import (
	"errors"
	"testing"
	"time"
	"water-tank-api/core/entity/access"
	stack "water-tank-api/core/entity/error_stack"
	data "water-tank-api/core/entity/water_tank"
	database_mock "water-tank-api/infra/database/mock"
)

type waterTankUpdateMockData struct {
	states map[string]map[string]*data.WaterTank
}

var updateMockDatabase = NewWaterTankUpdateMockData()
var MockTimeNow = time.Now()

func NewWaterTankUpdateMockData() *waterTankUpdateMockData {
	return &waterTankUpdateMockData{
		states: map[string]map[string]*data.WaterTank{
			"GROUP_1": {
				"TANK_1": {
					Name:              "TANK_1",
					Group:             "GROUP_1",
					MaximumCapacity:   100,
					TankState:         data.Empty,
					CurrentWaterLevel: 0,
					Access:            "a",
					LastFullTime:      MockTimeNow,
				},
			},
		},
	}
}

var successUpdateTank = NewWaterTankUpdate(updateMockDatabase)
var failUpdateTank = NewWaterTankUpdate(database_mock.NewWaterTankFailMockData())

func Test_WaterTank_Update(t *testing.T) {
	t.Run("Succesful update water tank (filling)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", 80)

		if err.HasError() {
			t.Error("Test_WaterTank_Update() shouldn't report an error, but it have")
		}

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != data.Filling {
			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", data.Filling, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
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

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != data.Full {
			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", data.Full, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
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

		if updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState != data.Empty {
			t.Errorf("Test_WaterTank_Update() wrong TankState. Expected '%d', set '%d'", data.Empty, updateMockDatabase.states["GROUP_1"]["TANK_1"].TankState)
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

		if err.LastError().Error() != WaterTankErrorNotFound("TANK_135sb3").Error() {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", WaterTankErrorNotFound("TANK_135sb3").Error(), err.LastError().Error())
		}
	})

	t.Run("Tank invalid water level (smaller than 0)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", -1)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankCurrentWaterLevelSmallerThanZero {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", WaterTankCurrentWaterLevelSmallerThanZero.Error(), err.LastError().Error())
		}
	})

	t.Run("Tank invalid water level (bigger thank maximum capacity)", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "a", 101)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankCurrentWaterLevelBiggerThanMax {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", WaterTankCurrentWaterLevelBiggerThanMax.Error(), err.LastError().Error())
		}
	})

	t.Run("Tank invalid access token", func(t *testing.T) {
		err := successUpdateTank.Update("TANK_1", "GROUP_1", "b", 45)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError() != WaterTankInvalidToken {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", WaterTankInvalidToken.Error(), err.LastError().Error())
		}
	})

	t.Run("Internal server error updating water tank", func(t *testing.T) {
		err := failUpdateTank.Update("TANK_1", "GROUP_1", "a", 100)

		if !err.HasError() {
			t.Error("Test_WaterTank_Update() should report an error, but it haven't")
		}

		if err.LastError().Error() != WaterTankErrorServerError(errors.New("ERROR").Error()).Error() {
			t.Errorf("Test_WaterTank_Update() wrong error. Should return '%s', got '%s'", WaterTankErrorServerError(errors.New("ERROR").Error()), err.LastError())
		}
	})
}

func (tank *waterTankUpdateMockData) GetWaterTankState(group string, names ...string) (state *data.WaterTank, err stack.ErrorStack) {
	state = tank.states[group][names[0]]
	return
}

func (tank *waterTankUpdateMockData) GetTankGroupState(groups ...string) (state []*data.WaterTank, err stack.ErrorStack) {
	if group, exists := tank.states[groups[0]]; exists {
		for _, tank := range group {
			state = append(state, tank)
		}
	}
	return
}

func (tank *waterTankUpdateMockData) NotifyFullTank(name string, group string, currentTime time.Time) (state *data.WaterTank, err stack.ErrorStack) {
	if group, exists := tank.states[group]; exists {
		group[name].LastFullTime = currentTime
	}
	return
}

func (tank *waterTankUpdateMockData) CreateWaterTank(name string, group string, accessToken access.AccessToken, capacity data.Capacity) (err stack.ErrorStack) {
	if _, exists := tank.states[group]; !exists {
		tank.states[group] = map[string]*data.WaterTank{
			name: {
				Name:              name,
				Group:             group,
				MaximumCapacity:   capacity,
				TankState:         data.Empty,
				CurrentWaterLevel: 0,
				Access:            accessToken,
			},
		}
		return
	}

	tank.states[group][name] = &data.WaterTank{
		Name:              name,
		Group:             group,
		MaximumCapacity:   capacity,
		TankState:         data.Empty,
		CurrentWaterLevel: 0,
		Access:            accessToken,
	}

	return
}

func (tank *waterTankUpdateMockData) UpdateWaterTankState(name string, group string, waterLevel data.Capacity, levelState data.State) (state *data.WaterTank, err stack.ErrorStack) {
	if group, exists := tank.states[group]; exists {
		group[name].CurrentWaterLevel = waterLevel
		group[name].TankState = levelState
	}
	return
}
