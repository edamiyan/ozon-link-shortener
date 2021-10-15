package v1

import "context"

type LinkService interface {
	CreateShortURL(ctx context.Context, baseUrl string) (string, error)
	GetBaseURL(ctx context.Context, token string) (string, error)
}
