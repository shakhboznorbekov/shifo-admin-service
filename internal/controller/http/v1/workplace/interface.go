package workplace

import (
	"context"
	"shifo-backend-website/internal/pkg"
	"shifo-backend-website/internal/repository/postgres/workplace"
)

type Workplace interface {
	AdminGetList(ctx context.Context, filter workplace.Filter) ([]workplace.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (workplace.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request workplace.AdminCreateRequest) (workplace.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request workplace.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}
