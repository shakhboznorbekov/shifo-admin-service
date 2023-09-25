package router

import (
	"github.com/gin-gonic/gin"
)

type Auth interface {
	HasPermission(roles ...string) gin.HandlerFunc
}

type Authorization interface {
	SignIn(*gin.Context)
}

type User interface {
	AdminGetUserList(*gin.Context)
	AdminGetUserDetail(*gin.Context)
	AdminCreateUser(*gin.Context)
	AdminUpdateUser(*gin.Context)
	AdminDeleteUser(*gin.Context)
}

type Doctor interface {
	AdminGetDoctorList(*gin.Context)
	AdminGetDoctorDetail(*gin.Context)
	AdminCreateDoctor(*gin.Context)
	AdminUpdateDoctor(*gin.Context)
	AdminDeleteDoctor(*gin.Context)
}
type Specialty interface {
	AdminGetSpecialtyList(*gin.Context)
	AdminGetSpecialtyDetail(*gin.Context)
	AdminCreateSpecialty(*gin.Context)
	AdminUpdateSpecialty(*gin.Context)
	AdminDeleteSpecialty(*gin.Context)
}
type Workplace interface {
	AdminGetWorkplaceList(*gin.Context)
	AdminGetWorkplaceDetail(*gin.Context)
	AdminCreateWorkplace(*gin.Context)
	AdminUpdateWorkplace(*gin.Context)
	AdminDeleteWorkplace(*gin.Context)
}
type Router struct {
	auth          Auth
	user          User
	authorization Authorization
	doctor        Doctor
	specialty     Specialty
	workplace     Workplace
}

func New(auth Auth, user User, authorization Authorization, doctor Doctor, specialty Specialty, workplace Workplace) *Router {
	return &Router{auth: auth,
		user:          user,
		authorization: authorization,
		doctor:        doctor,
		specialty:     specialty,
		workplace:     workplace,
	}
}

func (r *Router) Init(port string) error {
	router := gin.Default()

	// gin engine
	router.Use(customCORSMiddleware())

	// auth
	router.POST("/api/v1/user/sign-in", r.authorization.SignIn)

	//user
	router.GET("/api/v1/admin/user/list", r.auth.HasPermission("Admin"), r.user.AdminGetUserList)
	router.GET("/api/v1/admin/user/:id", r.auth.HasPermission("Admin"), r.user.AdminGetUserDetail)
	router.POST("/api/v1/admin/user/create", r.auth.HasPermission("Admin"), r.user.AdminCreateUser)
	router.PUT("/api/v1/admin/user/:id", r.auth.HasPermission("Admin"), r.user.AdminUpdateUser)
	router.DELETE("/api/v1/admin/user/:id", r.auth.HasPermission("Admin"), r.user.AdminDeleteUser)

	//doctor
	router.GET("/api/v1/admin/doctor/list", r.auth.HasPermission("Admin"), r.doctor.AdminGetDoctorList)
	router.GET("/api/v1/admin/doctor/:id", r.auth.HasPermission("Admin"), r.doctor.AdminGetDoctorDetail)
	router.POST("/api/v1/admin/doctor/create", r.auth.HasPermission("Admin"), r.doctor.AdminCreateDoctor)
	router.PUT("/api/v1/admin/doctor/:id", r.auth.HasPermission("Admin"), r.doctor.AdminUpdateDoctor)
	router.DELETE("/api/v1/admin/doctor/:id", r.auth.HasPermission("Admin"), r.doctor.AdminDeleteDoctor)

	//specialty
	router.GET("/api/v1/admin/specialty/list", r.auth.HasPermission("Admin"), r.specialty.AdminGetSpecialtyList)
	router.GET("/api/v1/admin/specialty/:id", r.auth.HasPermission("Admin"), r.specialty.AdminGetSpecialtyDetail)
	router.POST("/api/v1/admin/specialty/create", r.auth.HasPermission("Admin"), r.specialty.AdminCreateSpecialty)
	router.PUT("/api/v1/admin/specialty/:id", r.auth.HasPermission("Admin"), r.specialty.AdminUpdateSpecialty)
	router.DELETE("/api/v1/admin/specialty/:id", r.auth.HasPermission("Admin"), r.specialty.AdminDeleteSpecialty)

	//workplace
	router.GET("/api/v1/admin/workplace/list", r.auth.HasPermission("Admin"), r.workplace.AdminGetWorkplaceList)
	router.GET("/api/v1/admin/workplace/:id", r.auth.HasPermission("Admin"), r.workplace.AdminGetWorkplaceDetail)
	router.POST("/api/v1/admin/workplace/create", r.auth.HasPermission("Admin"), r.workplace.AdminCreateWorkplace)
	router.PUT("/api/v1/admin/workplace/:id", r.auth.HasPermission("Admin"), r.workplace.AdminUpdateWorkplace)
	router.DELETE("/api/v1/admin/workplace/:id", r.auth.HasPermission("Admin"), r.workplace.AdminDeleteWorkplace)

	return router.Run(port)
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}
}
