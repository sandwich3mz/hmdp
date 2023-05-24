package impl

import (
	"context"
	"hmdp/ent"
	"hmdp/services/blog_service"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) blog_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) QueryHotBlog(ctx context.Context, current int) (res []*ent.Blog, err error) {
	res, err = p.dbClient.Blog.
		Query().
		Order().
		Limit(10).
		Offset(current).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
