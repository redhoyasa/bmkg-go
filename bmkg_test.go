package bmkg

import (
	"fmt"
	"io/ioutil"

	"github.com/stretchr/testify/mock"
)

const (
	mocksLocation = "./mocks"
)

var (
	testClient *mockClient
)

type mockClient struct {
	mock.Mock
}

func (m *mockClient) GetXMLBytes(url string) ([]byte, error) {
	args := m.Called(url)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func setup() {
	testClient = new(mockClient)
}

func getMockResponse(file string) ([]byte, error) {
	fileName := fmt.Sprintf("%v/%v", mocksLocation, file)

	res, err := ioutil.ReadFile(fileName)
	if err != nil {
		return res, err
	}

	return res, nil
}
