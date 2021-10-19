package v1_test

import (
	"bytes"
	"fmt"
	v1 "github.com/edamiyan/ozon-link-shortener/internal/handler/v1"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_createShortURL(t *testing.T) {
	tests := map[string]struct {
		input              string
		mock               func(repos *MockLinkService)
		expectedStatusCode int
	}{
		"ok": {
			input: `{"base_url":"https://yandex.ru"}`,
			mock: func(repos *MockLinkService) {
				repos.EXPECT().CreateShortURL(gomock.Any(), gomock.Any()).Return("token", nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		"error_validation": {
			input:              `{"base_url":"https://yandex"}`,
			mock:               func(repos *MockLinkService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		"error_internal": {
			input: `{"base_url":"https://yandex.ru"}`,
			mock: func(repos *MockLinkService) {
				repos.EXPECT().CreateShortURL(gomock.Any(), gomock.Any()).Return("", fmt.Errorf("some error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	linkService := NewMockLinkService(ctrl)
	handlers := v1.NewHandler(linkService)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(linkService)

			r := echo.New()
			handlers.Init(r)

			req := httptest.NewRequest(
				http.MethodPost,
				"/api/v1/tokens",
				bytes.NewBufferString(tc.input),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatusCode, rec.Code)
		})
	}
}

func TestHandler_getBaseURL(t *testing.T) {
	tests := map[string]struct {
		inputToken         string
		mock               func(repos *MockLinkService)
		expectedStatusCode int
	}{
		"ok": {
			inputToken: "abc_012_yz",
			mock: func(repos *MockLinkService) {
				repos.EXPECT().GetBaseURL(gomock.Any(), &model.Link{Token: "abc_012_yz"}).
					Return("https://yandex.ru", nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		"error_validation": {
			inputToken:         "!@#$AZ09az_",
			mock:               func(repos *MockLinkService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		"error_internal": {
			inputToken: "abc_012_yz",
			mock: func(repos *MockLinkService) {
				repos.EXPECT().GetBaseURL(gomock.Any(), &model.Link{Token: "abc_012_yz"}).
					Return("", fmt.Errorf("some error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		"error_not_such_baseURL": {
			inputToken: "abc_012_yz",
			mock: func(repos *MockLinkService) {
				repos.EXPECT().GetBaseURL(gomock.Any(), &model.Link{Token: "abc_012_yz"}).
					Return("", nil)
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	personService := NewMockLinkService(ctrl)
	handlers := v1.NewHandler(personService)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(personService)

			r := echo.New()
			handlers.Init(r)

			req := httptest.NewRequest(
				http.MethodGet,
				"/api/v1/tokens/"+tc.inputToken,
				nil,
			)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatusCode, rec.Code)
		})
	}
}
