package doctor

import (
	"context"
	"shifo-backend-website/internal/pkg"
	"shifo-backend-website/internal/repository/postgres/doctor"
)

type Doctor interface {
	AdminGetList(ctx context.Context, filter doctor.Filter) ([]doctor.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (doctor.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request doctor.AdminCreateRequest) (doctor.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request doctor.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}
