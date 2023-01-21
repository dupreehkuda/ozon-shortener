package httpHandlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
	"github.com/dupreehkuda/ozon-shortener/internal/mock"
	"github.com/dupreehkuda/ozon-shortener/internal/server"
)

// TestHandlers_GetFullLink tests GetFullLink handler
func TestHandlers_GetFullLink(t *testing.T) {
	tests := map[string]struct {
		input        string
		mock         func(service *mock.MockService)
		expectedCode int
	}{
		"Success": {
			input: "qWeRtY_123",
			mock: func(repos *mock.MockService) {
				repos.EXPECT().GetFullLink(context.Background(), gomock.Any()).Return("https://youtube.com", nil)
			},
			expectedCode: http.StatusMovedPermanently,
		},
		"Not found in storage": {
			input: "qWeRtY_123",
			mock: func(repos *mock.MockService) {
				repos.EXPECT().GetFullLink(context.Background(), gomock.Any()).Return("", er.ErrNoSuchURL)
			},
			expectedCode: http.StatusNoContent,
		},
		"Format Error": {
			input:        "qWeRtY",
			mock:         func(repos *mock.MockService) {},
			expectedCode: http.StatusBadRequest,
		},
		"Validation Error": {
			input:        "qW+RtY=123",
			mock:         func(repos *mock.MockService) {},
			expectedCode: http.StatusBadRequest,
		},
		"Internal Error": {
			input: "qWeRtY_123",
			mock: func(repos *mock.MockService) {
				repos.EXPECT().GetFullLink(context.Background(), gomock.Any()).Return("", fmt.Errorf("error occurred"))
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockService(ctrl)
	logger, _ := zap.NewDevelopment()
	handle := New(service, logger)

	srv := server.NewByConfig(handle, nil, logger, nil)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(service)
			r := srv.Router()

			req := httptest.NewRequest(http.MethodGet, "/"+tc.input, nil)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
