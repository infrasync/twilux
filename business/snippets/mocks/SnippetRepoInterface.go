// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	snippets "twilux/business/snippets"

	mock "github.com/stretchr/testify/mock"
)

// SnippetRepoInterface is an autogenerated mock type for the SnippetRepoInterface type
type SnippetRepoInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: domain, ctx
func (_m *SnippetRepoInterface) Create(domain snippets.Domain, ctx context.Context) (snippets.Domain, error) {
	ret := _m.Called(domain, ctx)

	var r0 snippets.Domain
	if rf, ok := ret.Get(0).(func(snippets.Domain, context.Context) snippets.Domain); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(snippets.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(snippets.Domain, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: domain, ctx
func (_m *SnippetRepoInterface) Delete(domain snippets.Domain, ctx context.Context) (snippets.Domain, error) {
	ret := _m.Called(domain, ctx)

	var r0 snippets.Domain
	if rf, ok := ret.Get(0).(func(snippets.Domain, context.Context) snippets.Domain); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(snippets.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(snippets.Domain, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *SnippetRepoInterface) GetAll(ctx context.Context) ([]snippets.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []snippets.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []snippets.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]snippets.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: snippetId, ctx
func (_m *SnippetRepoInterface) GetById(snippetId string, ctx context.Context) (snippets.Domain, error) {
	ret := _m.Called(snippetId, ctx)

	var r0 snippets.Domain
	if rf, ok := ret.Get(0).(func(string, context.Context) snippets.Domain); ok {
		r0 = rf(snippetId, ctx)
	} else {
		r0 = ret.Get(0).(snippets.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(snippetId, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: domain, ctx
func (_m *SnippetRepoInterface) Update(domain snippets.Domain, ctx context.Context) (snippets.Domain, error) {
	ret := _m.Called(domain, ctx)

	var r0 snippets.Domain
	if rf, ok := ret.Get(0).(func(snippets.Domain, context.Context) snippets.Domain); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(snippets.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(snippets.Domain, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
