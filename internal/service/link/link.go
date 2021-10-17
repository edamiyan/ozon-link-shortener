package link

import (
	"context"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
	"github.com/edamiyan/ozon-link-shortener/internal/service"
)

type Service struct {
	repos Repository
}

func NewService(repos Repository) *Service {
	return &Service{repos: repos}
}

func (s Service) CreateShortURL(ctx context.Context, link *model.Link) (string, error) {
	link.Token = service.GenerateToken()
	token, err := s.repos.CreateShortURL(ctx, link)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) GetBaseURL(ctx context.Context, token string) (string, error) {
	panic("implement me")
}
