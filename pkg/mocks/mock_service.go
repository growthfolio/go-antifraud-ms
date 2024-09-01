package mocks

//"github.com/stretchr/testify/mock"

type MockService struct {
	//mock.Mock
}

func (m *MockService) SomeMethod(input string) (output string, err error) {
	// args := m.Called(input)
	return "StringValue", err
}
