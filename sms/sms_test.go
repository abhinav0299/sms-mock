package sms

import (
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockSMSSender(ctrl)
	h := New(m)
	testcase := []struct {
		to            string
		msg           string
		expectedError error
		mock          *gomock.Call
	}{
		{
			to:            "9898989899",
			msg:           "dtfgihjhg",
			expectedError: nil, mock: m.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil),
		},
		{
			to:            "9898989899898ktufygweruhrwesss",
			msg:           "dtfgihjhgfdsfghj",
			expectedError: errors.New("invalid phone"), mock: nil,
		},
		{
			to:            "98989898",
			msg:           "dtfgihjhgfdsfghjdxdxyfcukfvkuf       vkhvfkhfvkhfvyhfvyfc      yfukgvhfjydj   yxhfjyfyiudyfkuguedyf",
			expectedError: errors.New("invalid sms message"), mock: nil,
		},
	}
	for _, tt := range testcase {
		err := h.SendMessage(tt.to, tt.msg)
		if !reflect.DeepEqual(err, tt.expectedError) {
			t.Errorf("Test is failing")
		}
	}
}
