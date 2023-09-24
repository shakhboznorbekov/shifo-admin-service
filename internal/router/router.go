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

type Router struct {
	auth          Auth
	user          User
	authorization Authorization
}

func New(auth Auth, user User, authorization Authorization) *Router {
	return &Router{auth: auth,
		user:          user,
		authorization: authorization,
	}
}

// Init ...
// @title API
// @version 1
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name Shaxboz
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
