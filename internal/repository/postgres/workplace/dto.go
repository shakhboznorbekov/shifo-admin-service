package workplace

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
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Long    string `json:"long"`
}

type AdminGetDetail struct {
	bun.BaseModel `bun:"table:workplaces"`

	Id      string `json:"id" bun:"id"`
	Name    string `json:"name" bun:"name"`
	Address string `json:"address" bun:"address"`
	Lat     string `json:"lat" bun:"lat"`
	Long    string `json:"long" bun:"long"`
}

type AdminCreateRequest struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Lat     string `json:"lat" form:"lat"`
	Long    string `json:"long" form:"long"`
}

type AdminCreateResponse struct {
	bun.BaseModel `bun:"table:workplaces"`
	Id            string    `json:"id" bun:"id"`
	Name          string    `json:"name" bun:"name"`
	Address       string    `json:"address" bun:"address"`
	Lat           string    `json:"lat" bun:"lat"`
	Long          string    `json:"long" bun:"long"`
	CreatedAt     time.Time `json:"-" bun:"created_at"`
	CreatedBy     *string   `json:"-" bun:"created_by"`
}

type AdminUpdateRequest struct {
	Id      string  `json:"id" form:"id"`
	Name    *string `json:"name" form:"name"`
	Address *string `json:"address" form:"address"`
	Lat     *string `json:"lat" form:"lat"`
	Long    *string `json:"long" form:"long"`
}
