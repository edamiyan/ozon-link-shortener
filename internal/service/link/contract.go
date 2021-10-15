package link

import "context"

type Repository interface {
	CreateShortURL(ctx context.Context, baseUrl string) (string, error)
	GetBaseURL(ctx context.Context, token string) (string, error)
}
