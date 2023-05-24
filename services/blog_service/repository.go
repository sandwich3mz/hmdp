package blog_service

import (
	"context"
	"hmdp/ent"
)

type Repository interface {
	QueryHotBlog(ctx context.Context, current int) (res []*ent.Blog, err error)
}
