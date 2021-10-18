package v1

import (
	"context"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
)

type LinkService interface {
	CreateShortURL(ctx context.Context, link *model.Link) (string, error)
	GetBaseURL(ctx context.Context, link *model.Link) (string, error)
}
