package inMemory

import (
	"context"
	"fmt"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
)

func (r Repository) CreateShortURL(ctx context.Context, link *model.Link) (string, error) {
	if _, ok := r.briefToFull[link.Token]; ok {
		return "", fmt.Errorf("token already exist")
	}

	if token, ok := r.fullToBrief[link.BaseURL]; ok {
		return token, nil
	}

	r.briefToFull[link.Token] = link.BaseURL
	r.fullToBrief[link.BaseURL] = link.Token

	return link.Token, nil
}

func (r Repository) GetBaseURL(ctx context.Context, link *model.Link) (string, error) {
	if baseURL, ok := r.briefToFull[link.Token]; ok {
		return baseURL, nil
	}

	return "", fmt.Errorf("URL with this token not exist")
}
