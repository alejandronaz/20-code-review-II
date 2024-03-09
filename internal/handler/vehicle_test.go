package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// Test for FindByColorAndYear method of HandlerVehicle
// GET /vehicles/color/{color}/year/{year}
func TestFindByColorAndYear_Vehicle(t *testing.T) {

	t.Run("vehicles found - success", func(t *testing.T) {

		// arrange
		// - service
		sv := service.NewServiceVehicleMock()
		vehiclesToReturn := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Fiesta",
					Registration:    "ABC-1234",
					Color:           "red",
					FabricationYear: 2020,
					Capacity:        5,
					MaxSpeed:        180,
					FuelType:        "gasoline",
					Transmission:    "manual",
					Weight:          1000,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4,
						Width:  1.8,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Chevrolet",
					Model:           "Onix",
					Registration:    "ACC-1234",
					Color:           "red",
					FabricationYear: 2020,
					Capacity:        4,
					MaxSpeed:        180,
					FuelType:        "gasoline",
					Transmission:    "manual",
					Weight:          1000,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4,
						Width:  1.8,
					},
				},
			},
		}
		sv.On("FindByColorAndYear", "red", 2020).Return(vehiclesToReturn, nil)
		// - handler
		hd := handler.NewHandlerVehicle(sv)

		// - request
		req := httptest.NewRequest("GET", "/color/red/year/2020", nil)
		// - response
		res := httptest.NewRecorder()

		// - router
		r := chi.NewRouter()
		r.Get("/color/{color}/year/{year}", hd.FindByColorAndYear())

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedStatus := http.StatusOK
		require.Equal(t, expectedStatus, res.Code)
		expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
		require.Equal(t, expectedHeader, res.Header())
		expectedBody := `
		{
			"message": "vehicles found",
			"data": {
				"1": {
					"Id": 1,
					"Brand": "Ford",
					"Model": "Fiesta",
					"Registration": "ABC-1234",
					"Color": "red",
					"FabricationYear": 2020,
					"Capacity": 5,
					"MaxSpeed": 180,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1000,
					"Height": 1.5,
					"Length": 4,
					"Width": 1.8
				},
				"2": {
					"Id": 2,
					"Brand": "Chevrolet",
					"Model": "Onix",
					"Registration": "ACC-1234",
					"Color": "red",
					"FabricationYear": 2020,
					"Capacity": 4,
					"MaxSpeed": 180,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1000,
					"Height": 1.5,
					"Length": 4,
					"Width": 1.8
				}
			}
		}
		`
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("error - invalid year", func(t *testing.T) {

		// arrange
		// - service
		sv := service.NewServiceVehicleMock()
		// - handler
		hd := handler.NewHandlerVehicle(sv)

		// - request
		req := httptest.NewRequest("GET", "/color/red/year/200o", nil)
		// - response
		res := httptest.NewRecorder()

		// - router
		r := chi.NewRouter()
		r.Get("/color/{color}/year/{year}", hd.FindByColorAndYear())

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedStatus := http.StatusBadRequest
		require.Equal(t, expectedStatus, res.Code)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedHeader, res.Header())
		expectedBody := `
		{
			"status": "Bad Request",
			"message": "invalid year"
		}
		`
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("error - internal error", func(t *testing.T) {

		// arrange
		// - service
		sv := service.NewServiceVehicleMock()
		// - handler
		hd := handler.NewHandlerVehicle(sv)
		sv.On("FindByColorAndYear", "red", 2000).Return(map[int]internal.Vehicle{}, errors.New("internal error"))
		// - request
		req := httptest.NewRequest("GET", "/color/red/year/2000", nil)
		// - response
		res := httptest.NewRecorder()

		// - router
		r := chi.NewRouter()
		r.Get("/color/{color}/year/{year}", hd.FindByColorAndYear())

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedStatus := http.StatusInternalServerError
		require.Equal(t, expectedStatus, res.Code)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedHeader, res.Header())
		expectedBody := `
		{
			"status": "Internal Server Error",
			"message": "internal error"
		}
		`
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

}

// Test for FindByBrandAndYearRange method of HandlerVehicle
// GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func TestFindByBrandAndYearRange_Vehicle(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for AverageMaxSpeedByBrand method of HandlerVehicle
// GET /vehicles/average_speed/brand/{brand}
func TestAverageMaxSpeedByBrand_Vehicle(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for AverageCapacityByBrand method of HandlerVehicle
// GET /vehicles/average_capacity/brand/{brand}
func TestAverageCapacityByBrand_Vehicle(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for SearchByWeightRange method of HandlerVehicle
// GET /vehicles/weight
// GET /vehicles/weight?weight_min=1000&weight_max=2000
func TestSearchByWeightRange_Vehicle(t *testing.T) {

	t.Run("success", func(t *testing.T) {

		// arrange
		// - service
		sv := service.NewServiceVehicleMock()
		vehiclesToReturn := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Fiesta",
					Registration:    "ABC-1234",
					Color:           "red",
					FabricationYear: 2020,
					Capacity:        5,
					MaxSpeed:        180,
					FuelType:        "gasoline",
					Transmission:    "manual",
					Weight:          1200,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4,
						Width:  1.8,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Chevrolet",
					Model:           "Onix",
					Registration:    "ACC-1234",
					Color:           "red",
					FabricationYear: 2020,
					Capacity:        4,
					MaxSpeed:        180,
					FuelType:        "gasoline",
					Transmission:    "manual",
					Weight:          1400,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4,
						Width:  1.8,
					},
				},
			},
		}
		sv.On("SearchByWeightRange", internal.SearchQuery{
			FromWeight: 1000,
			ToWeight:   2000,
		}, true).Return(vehiclesToReturn, nil)
		// - handler
		hd := handler.NewHandlerVehicle(sv)

		// - request
		req := httptest.NewRequest("GET", "/vehicles/weight?weight_min=1000&weight_max=2000", nil)
		// - response
		res := httptest.NewRecorder()

		// - router
		r := chi.NewRouter()
		r.Get("/vehicles/weight", hd.SearchByWeightRange())

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedStatus := http.StatusOK
		require.Equal(t, expectedStatus, res.Code)
		expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
		require.Equal(t, expectedHeader, res.Header())
		expectedBody := `
		{
			"message": "vehicles found",
			"data": {
				"1": {
					"Id": 1,
					"Brand": "Ford",
					"Model": "Fiesta",
					"Registration": "ABC-1234",
					"Color": "red",
					"FabricationYear": 2020,
					"Capacity": 5,
					"MaxSpeed": 180,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1200,
					"Height": 1.5,
					"Length": 4,
					"Width": 1.8
				},
				"2": {
					"Id": 2,
					"Brand": "Chevrolet",
					"Model": "Onix",
					"Registration": "ACC-1234",
					"Color": "red",
					"FabricationYear": 2020,
					"Capacity": 4,
					"MaxSpeed": 180,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1400,
					"Height": 1.5,
					"Length": 4,
					"Width": 1.8
				}
			}
		}
		`
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)

	})

	t.Run("error - parse weight min", func(t *testing.T) {

		// arrange
		// - service
		sv := service.NewServiceVehicleMock()
		// - handler
		hd := handler.NewHandlerVehicle(sv)

		// - request
		req := httptest.NewRequest("GET", "/vehicles/weight?weight_min=100o&weight_max=2000", nil)
		// - response
		res := httptest.NewRecorder()

		// - router
		r := chi.NewRouter()
		r.Get("/vehicles/weight", hd.SearchByWeightRange())

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		require.Equal(t, expectedCode, res.Code)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedHeader, res.Header())
		expectedBody := `
		{
			"status": "Bad Request",
			"message": "invalid weight_min"
		}
		`
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)

	})

	t.Run("error - parse weight max", func(t *testing.T) {

		// arrange
		// - service
		sv := service.NewServiceVehicleMock()
		// - handler
		hd := handler.NewHandlerVehicle(sv)

		// - request
		req := httptest.NewRequest("GET", "/vehicles/weight?weight_min=1000&weight_max=200o", nil)
		// - response
		res := httptest.NewRecorder()

		// - router
		r := chi.NewRouter()
		r.Get("/vehicles/weight", hd.SearchByWeightRange())

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		require.Equal(t, expectedCode, res.Code)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedHeader, res.Header())
		expectedBody := `
		{
			"status": "Bad Request",
			"message": "invalid weight_max"
		}
		`
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)

	})

	t.Run("error - service error", func(t *testing.T) {

		// arrange
		// - service
		sv := service.NewServiceVehicleMock()
		vehiclesToReturn := make(map[int]internal.Vehicle)
		sv.On("SearchByWeightRange", internal.SearchQuery{
			FromWeight: 1000,
			ToWeight:   2000,
		}, true).Return(vehiclesToReturn, errors.New("error"))
		// - handler
		hd := handler.NewHandlerVehicle(sv)

		// - request
		req := httptest.NewRequest("GET", "/vehicles/weight?weight_min=1000&weight_max=2000", nil)
		// - response
		res := httptest.NewRecorder()

		// - router
		r := chi.NewRouter()
		r.Get("/vehicles/weight", hd.SearchByWeightRange())

		// act
		r.ServeHTTP(res, req)

		// assert
		expectedCode := http.StatusInternalServerError
		require.Equal(t, expectedCode, res.Code)
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedHeader, res.Header())
		expectedBody := `
		{
			"status": "Internal Server Error",
			"message": "internal error"
		}
		`
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)

	})

}
