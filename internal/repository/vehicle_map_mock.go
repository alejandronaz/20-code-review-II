package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// NewRepositoryReadVehicleMock is a function that returns a new instance of RepositoryReadVehicleMock
func NewRepositoryReadVehicleMock() *RepositoryReadVehicleMock {
	return &RepositoryReadVehicleMock{}
}

// RepositoryReadVehicleMock is a mock of RepositoryReadVehicle
type RepositoryReadVehicleMock struct {
	// embed the mock
	mock.Mock
}

// FindAll is a mock method that returns a map of all vehicles
func (r *RepositoryReadVehicleMock) FindAll() (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := r.Called()
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}

// FindByColorAndYear is a mock method that returns a map of vehicles that match the color and fabrication year
func (r *RepositoryReadVehicleMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := r.Called(color, fabricationYear)
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}

// FindByBrandAndYearRange is a mock method that returns a map of vehicles that match the brand and a range of fabrication years
func (r *RepositoryReadVehicleMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := r.Called(brand, startYear, endYear)
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}

// FindByBrand is a mock method that returns a map of vehicles that match the brand
func (r *RepositoryReadVehicleMock) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := r.Called(brand)
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}

// FindByWeightRange is a mock method that returns a map of vehicles that match the weight range
func (r *RepositoryReadVehicleMock) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := r.Called(fromWeight, toWeight)
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}
