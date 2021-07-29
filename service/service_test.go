package service

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type smsServiceMock struct {
	mock.Mock
}

func (m *smsServiceMock) SendChargeNotification(value int) bool  {
	fmt.Println("Mocked charge notification")
	args := m.Called(value)
	return args.Bool(0)
}

func TestChargeCustomer(t *testing.T) {
		smsService := new(smsServiceMock)
		smsService.On("SendChargeNotification", 100).Return(true)
		myService := MyService{smsService}

		myService.ChargeCustomer(100)
		smsService.AssertExpectations(t)
}

