package inMemory

import "context"

func (r Repository) CreateShortURL(ctx context.Context, baseUrl string) (string, error) {
	panic("implement me")
}

func (r Repository) GetBaseURL(ctx context.Context, token string) (string, error) {
	panic("implement me")
}
