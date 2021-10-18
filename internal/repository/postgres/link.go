package postgres

import (
	"context"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
)

func (r Repository) CreateShortURL(ctx context.Context, link *model.Link) (string, error) {
	panic("implement me")
}

func (r Repository) GetBaseURL(ctx context.Context, link *model.Link) (string, error) {
	panic("implement me")
}
