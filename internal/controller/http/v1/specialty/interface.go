package specialty

import (
	"context"
	"shifo-backend-website/internal/pkg"
	"shifo-backend-website/internal/repository/postgres/specialty"
)

type Specialty interface {
	AdminGetList(ctx context.Context, filter specialty.Filter) ([]specialty.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (specialty.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request specialty.AdminCreateRequest) (specialty.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request specialty.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}
