package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
	"github.com/dupreehkuda/ozon-shortener/internal/mock"
)

// TestService_GetFullLink tests GetFullLink logic
func TestService_GetFullLink(t *testing.T) {
	tests := map[string]struct {
		input          string
		mock           func(service *mock.MockStorage, token string)
		expectedOutput string
		expectedError  error
	}{
		"Success": {
			input: "qWeRtY_123",
			mock: func(storage *mock.MockStorage, token string) {
				storage.EXPECT().GetFullLink(context.Background(), token).Return("https://youtube.com", nil)
			},
			expectedOutput: "https://youtube.com",
			expectedError:  nil,
		},
		"Not found in storage": {
			input: "qWeRtY_123",
			mock: func(storage *mock.MockStorage, token string) {
				storage.EXPECT().GetFullLink(context.Background(), token).Return("", er.ErrNoSuchURL)
			},
			expectedOutput: "",
			expectedError:  er.ErrNoSuchURL,
		},
		"Internal Error": {
			input: "qWeRtY_123",
			mock: func(storage *mock.MockStorage, token string) {
				storage.EXPECT().GetFullLink(context.Background(), token).Return("", fmt.Errorf("error occurred"))
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

			ans, err := serv.GetFullLink(context.Background(), tc.input)
			if tc.expectedError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedOutput, ans)
			}
		})
	}
}
