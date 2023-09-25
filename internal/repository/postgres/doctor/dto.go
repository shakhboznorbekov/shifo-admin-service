package doctor

import (
	"github.com/uptrace/bun"
	"time"
)

type Filter struct {
	Limit     *int
	Offset    *int
	FirstName *string
}

type AdminGetListResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AdminGetDetail struct {
	bun.BaseModel `bun:"table:doctors"`

	Id             string `json:"id" bun:"id"`
	FirstName      string `json:"first_name" bun:"first_name"`
	LastName       string `json:"last_name" bun:"last_name"`
	SpecialtyId    string `json:"specialty_id" bun:"specialty_id"`
	FileLink       string `json:"file_link" bun:"file_link"`
	WorkExperience string `json:"work_experience" bun:"work_experience"`
	WorkplaceId    string `json:"workplace_id" bun:"workplace_id"`
	WorkPrice      string `json:"work_price" bun:"work_price"`
	StartWork      string `json:"start_work" bun:"start_work"`
	EndWork        string `json:"end_work" bun:"end_work"`
}

type AdminCreateRequest struct {
	FirstName      string `json:"first_name" form:"first_name"`
	LastName       string `json:"last_name" form:"last_name"`
	SpecialtyId    string `json:"specialty_id" form:"specialty_id"`
	FileLink       string `json:"file_link" form:"file_link"`
	WorkExperience string `json:"work_experience" form:"work_experience"`
	WorkplaceId    string `json:"workplace_id" form:"workplace_id"`
	WorkPrice      string `json:"work_price" form:"work_price"`
	StartWork      string `json:"start_work" form:"start_work"`
	EndWork        string `json:"end_work" form:"end_work"`
}

type AdminCreateResponse struct {
	bun.BaseModel  `bun:"table:doctors"`
	Id             string    `json:"id" bun:"id"`
	FirstName      string    `json:"first_name" bun:"first_name"`
	LastName       string    `json:"last_name" bun:"last_name"`
	SpecialtyId    string    `json:"specialty_id" bun:"specialty_id"`
	FileLink       string    `json:"file_link" bun:"file_link"`
	WorkExperience string    `json:"work_experience" bun:"work_experience"`
	WorkplaceId    string    `json:"workplace_id" bun:"workplace_id"`
	WorkPrice      string    `json:"work_price" bun:"work_price"`
	StartWork      string    `json:"start_work" bun:"start_work"`
	EndWork        string    `json:"end_work" bun:"end_work"`
	CreatedAt      time.Time `json:"-" bun:"created_at"`
	CreatedBy      *string   `json:"-" bun:"created_by"`
}

type AdminUpdateRequest struct {
	Id             string  `json:"id" form:"id"`
	FirstName      *string `json:"first_name" form:"first_name"`
	LastName       *string `json:"last_name" form:"last_name"`
	SpecialtyId    *string `json:"specialty_id" form:"specialty_id"`
	FileLink       *string `json:"file_link" form:"file_link"`
	WorkExperience *string `json:"work_experience" form:"work_experience"`
	WorkplaceId    *string `json:"workplace_id" form:"workplace_id"`
	WorkPrice      *string `json:"work_price" form:"work_price"`
	StartWork      *string `json:"start_work" form:"start_work"`
	EndWork        *string `json:"end_work" form:"end_work"`
}
