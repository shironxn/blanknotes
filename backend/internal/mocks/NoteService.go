// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/shironxn/blanknotes/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// NoteService is an autogenerated mock type for the NoteService type
type NoteService struct {
	mock.Mock
}

type NoteService_Expecter struct {
	mock *mock.Mock
}

func (_m *NoteService) EXPECT() *NoteService_Expecter {
	return &NoteService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: req
func (_m *NoteService) Create(req domain.NoteRequest) (*domain.Note, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.NoteRequest) (*domain.Note, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(domain.NoteRequest) *domain.Note); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.NoteRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type NoteService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - req domain.NoteRequest
func (_e *NoteService_Expecter) Create(req interface{}) *NoteService_Create_Call {
	return &NoteService_Create_Call{Call: _e.mock.On("Create", req)}
}

func (_c *NoteService_Create_Call) Run(run func(req domain.NoteRequest)) *NoteService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.NoteRequest))
	})
	return _c
}

func (_c *NoteService_Create_Call) Return(_a0 *domain.Note, _a1 error) *NoteService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteService_Create_Call) RunAndReturn(run func(domain.NoteRequest) (*domain.Note, error)) *NoteService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id, claims
func (_m *NoteService) Delete(id uint, claims domain.Claims) error {
	ret := _m.Called(id, claims)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, domain.Claims) error); ok {
		r0 = rf(id, claims)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NoteService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type NoteService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
//   - claims domain.Claims
func (_e *NoteService_Expecter) Delete(id interface{}, claims interface{}) *NoteService_Delete_Call {
	return &NoteService_Delete_Call{Call: _e.mock.On("Delete", id, claims)}
}

func (_c *NoteService_Delete_Call) Run(run func(id uint, claims domain.Claims)) *NoteService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(domain.Claims))
	})
	return _c
}

func (_c *NoteService_Delete_Call) Return(_a0 error) *NoteService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NoteService_Delete_Call) RunAndReturn(run func(uint, domain.Claims) error) *NoteService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: req, metadata
func (_m *NoteService) GetAll(req domain.NoteQuery, metadata *domain.Metadata) ([]domain.Note, error) {
	ret := _m.Called(req, metadata)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []domain.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.NoteQuery, *domain.Metadata) ([]domain.Note, error)); ok {
		return rf(req, metadata)
	}
	if rf, ok := ret.Get(0).(func(domain.NoteQuery, *domain.Metadata) []domain.Note); ok {
		r0 = rf(req, metadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.NoteQuery, *domain.Metadata) error); ok {
		r1 = rf(req, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteService_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type NoteService_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - req domain.NoteQuery
//   - metadata *domain.Metadata
func (_e *NoteService_Expecter) GetAll(req interface{}, metadata interface{}) *NoteService_GetAll_Call {
	return &NoteService_GetAll_Call{Call: _e.mock.On("GetAll", req, metadata)}
}

func (_c *NoteService_GetAll_Call) Run(run func(req domain.NoteQuery, metadata *domain.Metadata)) *NoteService_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.NoteQuery), args[1].(*domain.Metadata))
	})
	return _c
}

func (_c *NoteService_GetAll_Call) Return(_a0 []domain.Note, _a1 error) *NoteService_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteService_GetAll_Call) RunAndReturn(run func(domain.NoteQuery, *domain.Metadata) ([]domain.Note, error)) *NoteService_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: id, claims
func (_m *NoteService) GetByID(id uint, claims *domain.Claims) (*domain.Note, error) {
	ret := _m.Called(id, claims)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, *domain.Claims) (*domain.Note, error)); ok {
		return rf(id, claims)
	}
	if rf, ok := ret.Get(0).(func(uint, *domain.Claims) *domain.Note); ok {
		r0 = rf(id, claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, *domain.Claims) error); ok {
		r1 = rf(id, claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteService_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type NoteService_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - id uint
//   - claims *domain.Claims
func (_e *NoteService_Expecter) GetByID(id interface{}, claims interface{}) *NoteService_GetByID_Call {
	return &NoteService_GetByID_Call{Call: _e.mock.On("GetByID", id, claims)}
}

func (_c *NoteService_GetByID_Call) Run(run func(id uint, claims *domain.Claims)) *NoteService_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint), args[1].(*domain.Claims))
	})
	return _c
}

func (_c *NoteService_GetByID_Call) Return(_a0 *domain.Note, _a1 error) *NoteService_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteService_GetByID_Call) RunAndReturn(run func(uint, *domain.Claims) (*domain.Note, error)) *NoteService_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: req, claims
func (_m *NoteService) Update(req domain.NoteUpdateRequest, claims domain.Claims) (*domain.Note, error) {
	ret := _m.Called(req, claims)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.NoteUpdateRequest, domain.Claims) (*domain.Note, error)); ok {
		return rf(req, claims)
	}
	if rf, ok := ret.Get(0).(func(domain.NoteUpdateRequest, domain.Claims) *domain.Note); ok {
		r0 = rf(req, claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.NoteUpdateRequest, domain.Claims) error); ok {
		r1 = rf(req, claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type NoteService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - req domain.NoteUpdateRequest
//   - claims domain.Claims
func (_e *NoteService_Expecter) Update(req interface{}, claims interface{}) *NoteService_Update_Call {
	return &NoteService_Update_Call{Call: _e.mock.On("Update", req, claims)}
}

func (_c *NoteService_Update_Call) Run(run func(req domain.NoteUpdateRequest, claims domain.Claims)) *NoteService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.NoteUpdateRequest), args[1].(domain.Claims))
	})
	return _c
}

func (_c *NoteService_Update_Call) Return(_a0 *domain.Note, _a1 error) *NoteService_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteService_Update_Call) RunAndReturn(run func(domain.NoteUpdateRequest, domain.Claims) (*domain.Note, error)) *NoteService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewNoteService creates a new instance of NoteService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNoteService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NoteService {
	mock := &NoteService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
