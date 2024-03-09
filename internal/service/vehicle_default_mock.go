package service

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// NewServiceVehicleMock is a mock of ServiceVehicle
func NewServiceVehicleMock() *ServiceVehicleMock {
	return &ServiceVehicleMock{}
}

// ServiceVehicleMock is a struct that implements the ServiceVehicle interface
type ServiceVehicleMock struct {
	// embed the mock
	mock.Mock
}

// FindByColorAndYear is a mock method that returns a map of vehicles that match the color and fabrication year
func (s *ServiceVehicleMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := s.Called(color, fabricationYear)
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}

// FindByBrandAndYearRange is a mock method that returns a map of vehicles that match the brand and a range of fabrication years
func (s *ServiceVehicleMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := s.Called(brand, startYear, endYear)
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}

// AverageMaxSpeedByBrand is a mock method that returns the average speed of the vehicles by brand
func (s *ServiceVehicleMock) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	// call the mock
	outputs := s.Called(brand)
	// return the values
	return outputs.Get(0).(float64), outputs.Error(1)
}

// AverageCapacityByBrand is a mock method that returns the average capacity of the vehicles by brand
func (s *ServiceVehicleMock) AverageCapacityByBrand(brand string) (a int, err error) {
	// call the mock
	outputs := s.Called(brand)
	// return the values
	return outputs.Get(0).(int), outputs.Error(1)
}

// SearchByWeightRange is a mock method that returns a map of vehicles that match the weight range
func (s *ServiceVehicleMock) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	// call the mock
	outputs := s.Called(query, ok)
	// return the values
	return outputs.Get(0).(map[int]internal.Vehicle), outputs.Error(1)
}
