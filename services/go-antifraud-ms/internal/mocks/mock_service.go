package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) SomeMethod(input string) (output string, err error) {
	args := m.Called(input)
	return args.String(0), args.Error(1)
}
