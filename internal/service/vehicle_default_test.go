package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test for FindByColorAndYear method of ServiceVehicleDefault
func TestFindByColorAndYear_Vehicle_Default(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for FindByBrandAndYearRange method of ServiceVehicleDefault
func TestFindByBrandAndYearRange_Vehicle_Default(t *testing.T) {

	t.Run("success", func(t *testing.T) {

		// arrange
		// - repository
		rp := repository.NewRepositoryReadVehicleMock()
		vehiclesToReturn := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Fiesta",
					Registration:    "ABC-1234",
					Color:           "red",
					FabricationYear: 2012,
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
					Brand:           "Ford",
					Model:           "Onix",
					Registration:    "ACC-1234",
					Color:           "red",
					FabricationYear: 2014,
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
		rp.On("FindByBrandAndYearRange", "Ford", 2010, 2015).Return(vehiclesToReturn, nil)
		// - service
		sv := service.NewServiceVehicleDefault(rp)

		// act
		v, err := sv.FindByBrandAndYearRange("Ford", 2010, 2015)

		// assert
		require.NoError(t, err)
		require.Equal(t, vehiclesToReturn, v)
		rp.AssertExpectations(t)

	})

	t.Run("error", func(t *testing.T) {

		// arrange
		// - repository
		rp := repository.NewRepositoryReadVehicleMock()
		vehiclesToReturn := map[int]internal.Vehicle{}
		rp.On("FindByBrandAndYearRange", "Ford", 2010, 2015).Return(vehiclesToReturn, errors.New("error"))
		// - service
		sv := service.NewServiceVehicleDefault(rp)

		// act
		v, err := sv.FindByBrandAndYearRange("Ford", 2010, 2015)

		// assert
		require.Error(t, err)
		require.Equal(t, vehiclesToReturn, v)
		rp.AssertExpectations(t)

	})

}

// Test for AverageMaxSpeedByBrand method of ServiceVehicleDefault
func TestAverageMaxSpeedByBrand_Vehicle_Default(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for AverageCapacityByBrand method of ServiceVehicleDefault
func TestAverageCapacityByBrand_Vehicle_Default(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for SearchByWeightRange method of ServiceVehicleDefault
func TestSearchByWeightRange_Vehicle_Default(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}
