package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/dupreehkuda/ozon-shortener/internal/mock"
)

// TestService_ShortenLink tests ShortenLink logic
func TestService_ShortenLink(t *testing.T) {
	tests := map[string]struct {
		input          string
		mock           func(service *mock.MockStorage, url string)
		expectedOutput string
		expectedError  error
	}{
		"Success": {
			input: "https://youtube.com",
			mock: func(storage *mock.MockStorage, url string) {
				storage.EXPECT().AddNewLink(context.Background(), gomock.Any(), url).Return("qWeRtY_123", nil)
			},
			expectedOutput: "qWeRtY_123",
			expectedError:  nil,
		},
		"Internal Error": {
			input: "https://youtube.com",
			mock: func(storage *mock.MockStorage, url string) {
				storage.EXPECT().AddNewLink(context.Background(), gomock.Any(), url).Return("", fmt.Errorf("error occurred"))
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("error occurred"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockStorage(ctrl)
	logger, _ := zap.NewDevelopment()
	serv := New(storage, logger)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(storage, tc.input)

			ans, err := serv.ShortenLink(context.Background(), tc.input)
			if tc.expectedError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedOutput, ans)
			}
		})
	}
}
