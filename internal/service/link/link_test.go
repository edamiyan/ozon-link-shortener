package link_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
	"github.com/edamiyan/ozon-link-shortener/internal/service/link"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBaseURL(t *testing.T) {
	tests := map[string]struct {
		inputToken *model.Link
		want       string
		mock       func(repos *MockRepository, link *model.Link)
		err        error
	}{
		"ok": {
			inputToken: &model.Link{Token: "abc_012_yz"},
			want:       "https://ozon.ru",
			mock: func(repos *MockRepository, inputToken *model.Link) {
				repos.EXPECT().GetBaseURL(gomock.Any(), inputToken).
					Return("https://ozon.ru", nil)
			},
		},
		"error": {
			inputToken: &model.Link{Token: "abc_012_yz"},
			mock: func(repos *MockRepository, inputToken *model.Link) {
				repos.EXPECT().GetBaseURL(gomock.Any(), inputToken).
					Return("", fmt.Errorf("some error"))
			},
			err: fmt.Errorf("some error"),
		},
		"not_found": {
			inputToken: &model.Link{Token: "abc_012_yz"},
			mock: func(repos *MockRepository, inputToken *model.Link) {
				repos.EXPECT().GetBaseURL(gomock.Any(), inputToken).
					Return("", sql.ErrNoRows)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repos := NewMockRepository(ctrl)
	service := link.NewService(repos)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(repos, tc.inputToken)
			got, err := service.GetBaseURL(context.Background(), tc.inputToken)
			if tc.err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestCreateShortURL(t *testing.T) {
	tests := map[string]struct {
		baseURL *model.Link
		want    string
		mock    func(repos *MockRepository, link *model.Link)
		err     error
	}{
		"ok": {
			baseURL: &model.Link{BaseURL: "https://ozon.ru"},
			want:    "TOKEN_1234",
			mock: func(repos *MockRepository, baseURL *model.Link) {
				repos.EXPECT().CreateShortURL(gomock.Any(), baseURL).
					Return("TOKEN_1234", nil)
			},
		},
		"error": {
			baseURL: &model.Link{BaseURL: "https://ozon.ru"},
			mock: func(repos *MockRepository, baseURL *model.Link) {
				repos.EXPECT().CreateShortURL(gomock.Any(), baseURL).
					Return("", fmt.Errorf("some error"))
			},
			err: fmt.Errorf("some error"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repos := NewMockRepository(ctrl)
	service := link.NewService(repos)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.mock(repos, tc.baseURL)
			got, err := service.CreateShortURL(context.Background(), tc.baseURL)
			if tc.err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
