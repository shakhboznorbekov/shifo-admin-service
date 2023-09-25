package specialty

import (
	"github.com/uptrace/bun"
	"time"
)

type Filter struct {
	Limit  *int
	Offset *int
	Name   *string
}

type AdminGetListResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AdminGetDetail struct {
	bun.BaseModel `bun:"table:specialties"`

	Id   string `json:"id" bun:"id"`
	Name string `json:"name" bun:"name"`
}

type AdminCreateRequest struct {
	Name string `json:"name" form:"name"`
}

type AdminCreateResponse struct {
	bun.BaseModel `bun:"table:doctors"`
	Id            string    `json:"id" bun:"id"`
	Name          string    `json:"name" bun:"name"`
	CreatedAt     time.Time `json:"-" bun:"created_at"`
	CreatedBy     *string   `json:"-" bun:"created_by"`
}

type AdminUpdateRequest struct {
	Id   string  `json:"id" form:"id"`
	Name *string `json:"name" form:"name"`
}
