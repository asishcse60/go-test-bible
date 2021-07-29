package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendChargeNotification(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("SendChargeNotificationTest", func(t *testing.T) {
		smsServiceMock := NewMockMessageService(mockCtrl)
		smsServiceMock.EXPECT().SendChargeNotification(100).Return(true)
		dd := smsServiceMock.SendChargeNotification(100)
		assert.Equal(t, true, dd)
	})
}
func TestCustomerCharge(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("TestCustomerCharge", func(t *testing.T) {
		smsServiceMock := NewMockMessageService(mockCtrl)
		smsServiceMock.EXPECT().SendChargeNotification(100).Return(true)

		myService := MyService{smsServiceMock}
		dd := myService.ChargeCustomer(100)

		assert.Equal(t, nil, dd)
	})
}
