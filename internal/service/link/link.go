package link

import "context"

type Service struct {
	repos Repository
}

func NewService(repos Repository) *Service {
	return &Service{repos: repos}
}

func (s Service) CreateShortURL(ctx context.Context, baseUrl string) (string, error) {
	panic("implement me")
}

func (s Service) GetBaseURL(ctx context.Context, token string) (string, error) {
	panic("implement me")
}
