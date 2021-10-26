package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/edamiyan/ozon-link-shortener/internal/model"
)

func (r Repository) CreateShortURL(ctx context.Context, link *model.Link) (string, error) {
	var (
		query string
		args  []interface{}
		err   error
		token string
	)

	query, args, err = sq.Select("token").
		From(linksTable).
		Where(sq.Eq{
			"base_url": link.BaseURL,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	if err := r.db.GetContext(ctx, &token, query, args...); err == nil {
		return token, nil
	}

	query, args, err = sq.Insert(linksTable).
		SetMap(linkData(link)).
		Suffix("RETURNING token").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	if err := r.db.GetContext(ctx, &token, query, args...); err != nil {
		return "", err
	}

	return token, nil
}

func (r Repository) GetBaseURL(ctx context.Context, link *model.Link) (string, error) {
	query, args, err := sq.Select("base_url").
		From(linksTable).
		Where(sq.Eq{
			"token": link.Token,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	var baseURL string
	if err := r.db.GetContext(ctx, &baseURL, query, args...); err != nil {
		return "", err
	}

	return baseURL, nil
}

func linkData(p *model.Link) map[string]interface{} {
	data := map[string]interface{}{
		"base_url": p.BaseURL,
		"token":    p.Token,
	}

	return data
}
