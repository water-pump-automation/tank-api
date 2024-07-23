package controllers

import (
	"context"
	"errors"
	"testing"
	"time"
	"water-tank-api/app/core/entity/logs"
	"water-tank-api/app/core/usecases"
	register_tank "water-tank-api/app/core/usecases/create_tank"
	"water-tank-api/app/core/usecases/get_group"
	"water-tank-api/app/core/usecases/get_tank"
	update_tank_state "water-tank-api/app/core/usecases/update_tank_state"
	database_mock "water-tank-api/app/infra/database/mock"
)

func Test_Controller_Create(t *testing.T) {
	var successController = NewController(database_mock.NewWaterTankMockData())
	var failController = NewController(database_mock.NewWaterTankFailMockData())
	logs.SetLogger(&_empty{})

	Test_Controller_Successful_Create := func(t *testing.T) {
		t.Run("Successful create", func(t *testing.T) {
			_, err := successController.Create("TANK_235f", "GROUP_1", 100)

			if err != nil {
				t.Errorf("Test_Controller_Successful_Create() shouldn't return error. Got '%s'", err.Error())
			}
		})
	}

	Test_Controller_InvalidRequest_Create := func(t *testing.T) {
		t.Run("Invalid request create", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankInvalidRequest, register_tank.ErrWaterTankMaximumCapacityZero.Error())

			response, err := successController.Create("TANK_1sdvW4", "GROUP_1", -10)

			if err == nil {
				t.Errorf("Test_Controller_InvalidRequest_Create() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_InvalidRequest_Create() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_InternalError_Create := func(t *testing.T) {
		t.Run("Internal error create", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankInternalServerError, get_tank.ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error())

			response, err := failController.Create("TANK_1", "GROUP_1", 100)

			if err == nil {
				t.Errorf("Test_Controller_InternalError_Create() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_InternalError_Create() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_Successful_Create(t)
	Test_Controller_InvalidRequest_Create(t)
	Test_Controller_InternalError_Create(t)
}

func Test_Controller_Update(t *testing.T) {
	var successController = NewController(database_mock.NewWaterTankMockData())
	var failController = NewController(database_mock.NewWaterTankFailMockData())
	logs.SetLogger(&_empty{})

	Test_Controller_Successful_Update := func(t *testing.T) {
		t.Run("Successful update", func(t *testing.T) {
			_, err := successController.Update("TANK_1", "GROUP_1", "a", 10)

			if err != nil {
				t.Errorf("Test_Controller_Successful_Update() shouldn't return error. Got '%s'", err.Error())
			}
		})
	}

	Test_Controller_NotFound_Update := func(t *testing.T) {
		t.Run("Invalid request update", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankNotFound, update_tank_state.ErrWaterTankErrorNotFound("TANK_1vSND3").Error())

			response, err := successController.Update("TANK_1vSND3", "GROUP_1", "a", 10)

			if err == nil {
				t.Errorf("Test_Controller_InvalidRequest_Update() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_InvalidRequest_Update() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_InvalidRequest_Update := func(t *testing.T) {
		t.Run("Invalid request update", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankInvalidRequest, update_tank_state.ErrWaterTankCurrentWaterLevelSmallerThanZero.Error())

			response, err := successController.Update("TANK_1", "GROUP_1", "a", -10)

			if err == nil {
				t.Errorf("Test_Controller_InvalidRequest_Update() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_InvalidRequest_Update() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_InternalError_Update := func(t *testing.T) {
		t.Run("Internal error update", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankInternalServerError, get_tank.ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error())

			response, err := failController.Update("TANK_1", "GROUP_1", "a", 10)

			if err == nil {
				t.Errorf("Test_Controller_InternalError_Update() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_InternalError_Update() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_Successful_Update(t)
	Test_Controller_NotFound_Update(t)
	Test_Controller_InvalidRequest_Update(t)
	Test_Controller_InternalError_Update(t)
}

func Test_Controller_Get(t *testing.T) {
	var successController = NewController(database_mock.NewWaterTankMockData())
	var failController = NewController(database_mock.NewWaterTankFailMockData())
	logs.SetLogger(&_empty{})

	Test_Controller_Successful_Get := func(t *testing.T) {
		t.Run("Successful get", func(t *testing.T) {
			expectedReturn := NewControllerResponse(WaterTankOK, &usecases.WaterTankState{
				Name:              "TANK_1",
				Group:             "GROUP_1",
				MaximumCapacity:   "100.00L",
				TankState:         "EMPTY",
				CurrentWaterLevel: "0.00L",
				LastFullTime:      database_mock.MockTimeNow,
			})

			response, err := successController.Get("TANK_1", "GROUP_1")

			if err != nil {
				t.Errorf("Test_Controller_Successful_Get() shouldn't return error. Got '%s'", err.Error())
			}

			if field, equal := responsesAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_Successful_Get() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_NotFound_Get := func(t *testing.T) {
		t.Run("Not found get", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankNotFound, get_tank.ErrWaterTankErrorNotFound("TANK_1sdvW4").Error())

			response, err := successController.Get("TANK_1sdvW4", "GROUP_1")

			if err == nil {
				t.Errorf("Test_Controller_NotFound_Get() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_NotFound_Get() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_InternalError_Get := func(t *testing.T) {
		t.Run("Internal error get", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankInternalServerError, get_tank.ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error())

			response, err := failController.Get("TANK_1", "GROUP_1")

			if err == nil {
				t.Errorf("Test_Controller_InternalError_Get() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_InternalError_Get() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_Successful_Get(t)
	Test_Controller_NotFound_Get(t)
	Test_Controller_InternalError_Get(t)
}

func Test_Controller_GetGroup(t *testing.T) {
	var successController = NewController(database_mock.NewWaterTankMockData())
	var failController = NewController(database_mock.NewWaterTankFailMockData())
	logs.SetLogger(&_empty{})

	Test_Controller_Successful_GetGroup := func(t *testing.T) {
		t.Run("Successful get group", func(t *testing.T) {
			_, err := successController.GetGroup("GROUP_1")

			if err != nil {
				t.Errorf("Test_Controller_Successful_GetGroup() shouldn't return error. Got '%s'", err.Error())
			}
		})
	}

	Test_Controller_NotFound_GetGroup := func(t *testing.T) {
		t.Run("Not found get group", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankNotFound, get_group.ErrWaterTankErrorGroupNotFound("GROUP_1sdvW4").Error())

			response, err := successController.GetGroup("GROUP_1sdvW4")

			if err == nil {
				t.Errorf("Test_Controller_NotFound_GetGroup() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_NotFound_GetGroup() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_InternalError_GetGroup := func(t *testing.T) {
		t.Run("Internal error get group", func(t *testing.T) {
			expectedReturn := NewControllerError(WaterTankInternalServerError, get_group.ErrWaterTankErrorServerError(errors.New("ERROR").Error()).Error())

			response, err := failController.GetGroup("GROUP_1")

			if err == nil {
				t.Errorf("Test_Controller_InternalError_GetGroup() should return error")
			}

			if field, equal := responsesErrorAreEqual(expectedReturn, response); !equal {
				t.Errorf("Test_Controller_InternalError_GetGroup() wrong '%s' field in response", field)
			}
		})
	}

	Test_Controller_Successful_GetGroup(t)
	Test_Controller_NotFound_GetGroup(t)
	Test_Controller_InternalError_GetGroup(t)
}

// //////////////////////////////////////////////////////////////////////////////////////////////

func responsesErrorAreEqual(expected, got *ControllerResponse) (string, bool) {
	if expected.Code != got.Code {
		return "Code", false
	}

	if expected.Content["error"] != got.Content["error"] {
		return "error", false
	}

	return "", true
}

func responsesAreEqual(expected, got *ControllerResponse) (string, bool) {
	if expected.Code != got.Code {
		return "Code", false
	}

	if expected.Content["name"] != got.Content["name"] {
		return "name", false
	}

	if expected.Content["group"] != got.Content["group"] {
		return "group", false
	}

	if expected.Content["maximum_capacity"] != got.Content["maximum_capacity"] {
		return "maximum_capacity", false
	}

	if expected.Content["tank_state"] != got.Content["tank_state"] {
		return "tank_state", false
	}

	if expected.Content["current_water_level"] != got.Content["current_water_level"] {
		return "current_water_level", false
	}

	if expected.Content["last_full_time"] != got.Content["last_full_time"] {
		return "last_full_time", false
	}

	return "", true
}

type _empty struct {
}

func (logger *_empty) Context(ctx context.Context) logs.Logger {
	return &_empty{}
}

func (logger *_empty) Error(message string) time.Time {
	return time.Now()
}

func (logger *_empty) Fatal(message string) time.Time {
	return time.Now()
}

func (logger *_empty) Info(message string) time.Time {
	return time.Now()
}
