package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test for FindAll method of RepositoryReadVehicleMap
func TestFindAll_Vehicle_Map(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for FindByColorAndYear method of RepositoryReadVehicleMap
func TestFindByColorAndYear_Vehicle_Map(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for FindByBrandAndYearRange method of RepositoryReadVehicleMap
func TestFindByBrandAndYearRange_Vehicle_Map(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for FindByBrand method of RepositoryReadVehicleMap
func TestFindByBrand_Vehicle_Map(t *testing.T) {

	t.Run("success", func(t *testing.T) {

	})

	t.Run("other", func(t *testing.T) {

	})

}

// Test for FindByWeightRange method of RepositoryReadVehicleMap
func TestFindByWeightRange_Vehicle_Map(t *testing.T) {

	t.Run("success", func(t *testing.T) {

		// arrange
		// - db (map)
		db := map[int]internal.Vehicle{
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
					Brand:           "Chevrolet",
					Model:           "Onix",
					Registration:    "ACC-1234",
					Color:           "red",
					FabricationYear: 2014,
					Capacity:        4,
					MaxSpeed:        180,
					FuelType:        "gasoline",
					Transmission:    "manual",
					Weight:          1100,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4,
						Width:  1.8,
					},
				},
			},
			3: {
				Id: 3,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Bronco",
					Registration:    "CCC-1234",
					Color:           "red",
					FabricationYear: 2014,
					Capacity:        4,
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
			4: {
				Id: 4,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Volkswagen",
					Model:           "Polo",
					Registration:    "AAA-1234",
					Color:           "red",
					FabricationYear: 2014,
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
		// - repository
		rp := repository.NewRepositoryReadVehicleMap(db)

		// act
		vehicles, err := rp.FindByWeightRange(900, 1100)

		// assert
		require.NoError(t, err)
		expectedVehicles := make(map[int]internal.Vehicle)
		expectedVehicles[1] = db[1]
		expectedVehicles[2] = db[2]
		require.Equal(t, expectedVehicles, vehicles)

	})

	t.Run("success - no vehicles returned", func(t *testing.T) {

		// arrange
		// - db (map)
		db := map[int]internal.Vehicle{
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
					Brand:           "Chevrolet",
					Model:           "Onix",
					Registration:    "ACC-1234",
					Color:           "red",
					FabricationYear: 2014,
					Capacity:        4,
					MaxSpeed:        180,
					FuelType:        "gasoline",
					Transmission:    "manual",
					Weight:          1100,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4,
						Width:  1.8,
					},
				},
			},
			3: {
				Id: 3,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Bronco",
					Registration:    "CCC-1234",
					Color:           "red",
					FabricationYear: 2014,
					Capacity:        4,
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
			4: {
				Id: 4,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Volkswagen",
					Model:           "Polo",
					Registration:    "AAA-1234",
					Color:           "red",
					FabricationYear: 2014,
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
		// - repository
		rp := repository.NewRepositoryReadVehicleMap(db)

		// act
		vehicles, err := rp.FindByWeightRange(900, 999)

		// assert
		require.NoError(t, err)
		expectedVehicles := make(map[int]internal.Vehicle)
		require.Equal(t, expectedVehicles, vehicles)

	})

}
