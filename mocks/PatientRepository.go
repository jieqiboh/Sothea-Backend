// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/jieqiboh/sothea_backend/entities"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PatientRepository is an autogenerated mock type for the PatientRepository type
type PatientRepository struct {
	mock.Mock
}

// CreatePatient provides a mock function with given fields: ctx, admin
func (_m *PatientRepository) CreatePatient(ctx context.Context, admin *entities.Admin) (int32, error) {
	ret := _m.Called(ctx, admin)

	if len(ret) == 0 {
		panic("no return value specified for CreatePatient")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Admin) (int32, error)); ok {
		return rf(ctx, admin)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Admin) int32); ok {
		r0 = rf(ctx, admin)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.Admin) error); ok {
		r1 = rf(ctx, admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePatientVisit provides a mock function with given fields: ctx, id, admin
func (_m *PatientRepository) CreatePatientVisit(ctx context.Context, id int32, admin *entities.Admin) (int32, error) {
	ret := _m.Called(ctx, id, admin)

	if len(ret) == 0 {
		panic("no return value specified for CreatePatientVisit")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, *entities.Admin) (int32, error)); ok {
		return rf(ctx, id, admin)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, *entities.Admin) int32); ok {
		r0 = rf(ctx, id, admin)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, *entities.Admin) error); ok {
		r1 = rf(ctx, id, admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePatientVisit provides a mock function with given fields: ctx, id, vid
func (_m *PatientRepository) DeletePatientVisit(ctx context.Context, id int32, vid int32) error {
	ret := _m.Called(ctx, id, vid)

	if len(ret) == 0 {
		panic("no return value specified for DeletePatientVisit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) error); ok {
		r0 = rf(ctx, id, vid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExportDatabaseToCSV provides a mock function with given fields: ctx, includePhoto
func (_m *PatientRepository) ExportDatabaseToCSV(ctx context.Context, includePhoto bool) error {
	ret := _m.Called(ctx, includePhoto)

	if len(ret) == 0 {
		panic("no return value specified for ExportDatabaseToCSV")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) error); ok {
		r0 = rf(ctx, includePhoto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPatientVisitMeta provides a mock function with given fields: ctx, date
func (_m *PatientRepository) GetAllPatientVisitMeta(ctx context.Context, date time.Time) ([]entities.PatientVisitMeta, error) {
	ret := _m.Called(ctx, date)

	if len(ret) == 0 {
		panic("no return value specified for GetAllPatientVisitMeta")
	}

	var r0 []entities.PatientVisitMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) ([]entities.PatientVisitMeta, error)); ok {
		return rf(ctx, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) []entities.PatientVisitMeta); ok {
		r0 = rf(ctx, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.PatientVisitMeta)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDBUser provides a mock function with given fields: ctx, username
func (_m *PatientRepository) GetDBUser(ctx context.Context, username string) (*entities.DBUser, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetDBUser")
	}

	var r0 *entities.DBUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.DBUser, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.DBUser); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.DBUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPatientMeta provides a mock function with given fields: ctx, id
func (_m *PatientRepository) GetPatientMeta(ctx context.Context, id int32) (*entities.PatientMeta, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPatientMeta")
	}

	var r0 *entities.PatientMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (*entities.PatientMeta, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) *entities.PatientMeta); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.PatientMeta)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPatientVisit provides a mock function with given fields: ctx, id, vid
func (_m *PatientRepository) GetPatientVisit(ctx context.Context, id int32, vid int32) (*entities.Patient, error) {
	ret := _m.Called(ctx, id, vid)

	if len(ret) == 0 {
		panic("no return value specified for GetPatientVisit")
	}

	var r0 *entities.Patient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) (*entities.Patient, error)); ok {
		return rf(ctx, id, vid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) *entities.Patient); ok {
		r0 = rf(ctx, id, vid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Patient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, int32) error); ok {
		r1 = rf(ctx, id, vid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePatientVisit provides a mock function with given fields: ctx, id, vid, patient
func (_m *PatientRepository) UpdatePatientVisit(ctx context.Context, id int32, vid int32, patient *entities.Patient) error {
	ret := _m.Called(ctx, id, vid, patient)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePatientVisit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32, *entities.Patient) error); ok {
		r0 = rf(ctx, id, vid, patient)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPatientRepository creates a new instance of PatientRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPatientRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PatientRepository {
	mock := &PatientRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
