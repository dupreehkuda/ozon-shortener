package httpHandlers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/dupreehkuda/ozon-shortener/internal/mock"
	"github.com/dupreehkuda/ozon-shortener/internal/server"
)

// TestHandlers_ShortenLink tests ShortenLink handler
func TestHandlers_ShortenLink(t *testing.T) {
	tests := map[string]struct {
		input        string
		mock         func(service *mock.MockService)
		expectedCode int
	}{
		"Success": {
			input: `{"url":"https://youtube.com"}`,
			mock: func(repos *mock.MockService) {
				repos.EXPECT().ShortenLink(context.Background(), gomock.Any()).Return("test", nil)
			},
			expectedCode: http.StatusOK,
		},
		"Format Error": {
			input:        `{"long_url":"https://youtube.com"}`,
			mock:         func(repos *mock.MockService) {},
			expectedCode: http.StatusBadRequest,
		},
		"Validation Error": {
			input:        `{"url":"https:/youtube.com"}`,
			mock:         func(repos *mock.MockService) {},
			expectedCode: http.StatusBadRequest,
		},
		"Internal Error": {
			input: `{"url":"https://youtube.com"}`,
			mock: func(repos *mock.MockService) {
				repos.EXPECT().ShortenLink(context.Background(), gomock.Any()).Return("", fmt.Errorf("error occurred"))
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

			req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBufferString(tc.input))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
