package v1

import (
	"context"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
)

//go:generate mockgen -package ${GOPACKAGE}_test -source $GOFILE -destination mocks_test.go
type LinkService interface {
	CreateShortURL(ctx context.Context, link *model.Link) (string, error)
	GetBaseURL(ctx context.Context, link *model.Link) (string, error)
}
