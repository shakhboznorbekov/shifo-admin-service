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

type Post interface {
	AdminGetPostList(*gin.Context)
	AdminGetPostDetail(*gin.Context)
	AdminCreatePost(*gin.Context)
	AdminUpdatePost(*gin.Context)
	AdminDeletePost(*gin.Context)
	AdminPostFileCreate(ctx *gin.Context)
	AdminUpdatePostFile(ctx *gin.Context)
	AdminDeletePostFile(ctx *gin.Context)
	SiteGetPostList(*gin.Context)
	SiteGetPostDetail(*gin.Context)
}

type Faq interface {
	AdminGetFaqList(ctx *gin.Context)
	AdminCreateFaq(ctx *gin.Context)
	AdminUpdateFaq(ctx *gin.Context)
	AdminUpdatePatchFaq(ctx *gin.Context)
	AdminDeleteFaq(ctx *gin.Context)
	SiteGetFaqList(ctx *gin.Context)
}

type Opportunity interface {
	AdminGetOpportunityList(*gin.Context)
	AdminGetOpportunityDetail(*gin.Context)
	AdminCreateOpportunity(*gin.Context)
	AdminUpdateOpportunity(*gin.Context)
	AdminDeleteOpportunity(*gin.Context)
	AdminOpportunityFileCreate(ctx *gin.Context)
	AdminUpdateOpportunityFile(ctx *gin.Context)
	AdminDeleteOpportunityFile(ctx *gin.Context)
	SiteGetOpportunityList(*gin.Context)
	SiteGetOpportunityDetail(*gin.Context)
}

type Menu interface {
	AdminGetMenuList(ctx *gin.Context)
	AdminGetMenuDetail(ctx *gin.Context)
	AdminCreateMenu(ctx *gin.Context)
	AdminUpdateMenu(ctx *gin.Context)
	AdminDeleteMenu(ctx *gin.Context)
}
type Request interface {
	AdminGetRequestList(ctx *gin.Context)
	AdminGetRequestDetail(ctx *gin.Context)
	SiteCreateRequest(ctx *gin.Context)
	SiteRequestFileCreate(ctx *gin.Context)
}

type Contact interface {
	AdminGetContactList(ctx *gin.Context)
	AdminGetContactDetail(ctx *gin.Context)
	AdminCreateContact(ctx *gin.Context)
	AdminUpdateContact(ctx *gin.Context)
	AdminUpdatePatchContact(ctx *gin.Context)
	AdminDeleteContact(ctx *gin.Context)
	SiteGetContactList(ctx *gin.Context)
	SiteGetContactDetail(ctx *gin.Context)
}
type Router struct {
	auth          Auth
	user          User
	authorization Authorization
	post          Post
	faq           Faq
	opportunity   Opportunity
	menu          Menu
	request       Request
	contact       Contact
}

func New(auth Auth, user User, authorization Authorization, post Post, faq Faq, opportunity Opportunity, menu Menu, request Request, contact Contact) *Router {
	return &Router{auth: auth,
		user:          user,
		authorization: authorization,
		post:          post,
		faq:           faq,
		opportunity:   opportunity,
		menu:          menu,
		request:       request,
		contact:       contact,
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

	//post
	router.GET("/api/v1/admin/post/list", r.auth.HasPermission("Admin"), r.post.AdminGetPostList)
	router.GET("/api/v1/admin/post/:id", r.auth.HasPermission("Admin"), r.post.AdminGetPostDetail)
	router.POST("/api/v1/admin/post/create", r.auth.HasPermission("Admin"), r.post.AdminCreatePost)
	router.PUT("/api/v1/admin/post/:id", r.auth.HasPermission("Admin"), r.post.AdminUpdatePost)
	router.DELETE("/api/v1/admin/post/:id", r.auth.HasPermission("Admin"), r.post.AdminDeletePost)
	router.POST("/api/v1/admin/post-file/create", r.auth.HasPermission("Admin"), r.post.AdminPostFileCreate)
	router.PUT("/api/v1/admin/post-file/:id", r.auth.HasPermission("Admin"), r.post.AdminUpdatePostFile)
	router.DELETE("/api/v1/admin/post-file/:id", r.auth.HasPermission("Admin"), r.post.AdminUpdatePostFile)

	//faq
	router.GET("/api/v1/admin/faq/list", r.auth.HasPermission("Admin"), r.faq.AdminGetFaqList)
	router.POST("/api/v1/admin/faq/create", r.auth.HasPermission("Admin"), r.faq.AdminCreateFaq)
	router.PUT("/api/v1/admin/faq/:id", r.auth.HasPermission("Admin"), r.faq.AdminUpdateFaq)
	router.PATCH("/api/v1/admin/faq/:id", r.auth.HasPermission("Admin"), r.faq.AdminUpdatePatchFaq)
	router.DELETE("/api/v1/admin/faq/:id", r.auth.HasPermission("Admin"), r.faq.AdminDeleteFaq)

	//opportunity
	router.GET("/api/v1/admin/opportunity/list", r.auth.HasPermission("Admin"), r.opportunity.AdminGetOpportunityList)
	router.GET("/api/v1/admin/opportunity/:id", r.auth.HasPermission("Admin"), r.opportunity.AdminGetOpportunityDetail)
	router.POST("/api/v1/admin/opportunity/create", r.auth.HasPermission("Admin"), r.opportunity.AdminCreateOpportunity)
	router.PUT("/api/v1/admin/opportunity/:id", r.auth.HasPermission("Admin"), r.opportunity.AdminUpdateOpportunity)
	router.DELETE("/api/v1/admin/opportunity/:id", r.auth.HasPermission("Admin"), r.opportunity.AdminDeleteOpportunity)
	router.POST("/api/v1/admin/opportunity-file/create", r.auth.HasPermission("Admin"), r.opportunity.AdminOpportunityFileCreate)
	router.PUT("/api/v1/admin/opportunity-file/:id", r.auth.HasPermission("Admin"), r.opportunity.AdminUpdateOpportunityFile)
	router.DELETE("/api/v1/admin/opportunity-file/:id", r.auth.HasPermission("Admin"), r.opportunity.AdminDeleteOpportunityFile)

	//menu
	router.GET("/api/v1/admin/menu/list", r.auth.HasPermission("Admin"), r.menu.AdminGetMenuList)
	router.GET("/api/v1/admin/menu/:id", r.auth.HasPermission("Admin"), r.menu.AdminGetMenuDetail)
	router.POST("/api/v1/admin/menu/create", r.auth.HasPermission("Admin"), r.menu.AdminCreateMenu)
	router.PUT("/api/v1/admin/menu/:id", r.auth.HasPermission("Admin"), r.menu.AdminUpdateMenu)
	router.DELETE("/api/v1/admin/menu/:id", r.auth.HasPermission("Admin"), r.menu.AdminDeleteMenu)

	//request
	router.GET("/api/v1/admin/request/list", r.auth.HasPermission("Admin"), r.request.AdminGetRequestList)
	router.GET("/api/v1/admin/request/:id", r.auth.HasPermission("Admin"), r.request.AdminGetRequestDetail)

	//contact
	router.GET("/api/v1/admin/contact/list", r.auth.HasPermission("Admin"), r.contact.AdminGetContactList)
	router.GET("/api/v1/admin/contact/:id", r.auth.HasPermission("Admin"), r.contact.AdminGetContactDetail)
	router.POST("/api/v1/admin/contact/create", r.auth.HasPermission("Admin"), r.contact.AdminCreateContact)
	router.PUT("/api/v1/admin/contact/:id", r.auth.HasPermission("Admin"), r.contact.AdminUpdateContact)
	router.PATCH("/api/v1/admin/contact/:id", r.auth.HasPermission("Admin"), r.contact.AdminUpdatePatchContact)
	router.DELETE("/api/v1/admin/contact/:id", r.auth.HasPermission("Admin"), r.contact.AdminDeleteContact)

	//post-site
	router.GET("/api/v1/site/post/list", r.post.SiteGetPostList)
	router.GET("/api/v1/site/post/:id", r.post.SiteGetPostDetail)

	//faq-site
	router.GET("/api/v1/site/faq/list", r.faq.SiteGetFaqList)

	//opportunity-site
	router.GET("/api/v1/site/opportunity/list", r.opportunity.SiteGetOpportunityList)
	router.GET("/api/v1/site/opportunity/:id", r.opportunity.SiteGetOpportunityDetail)

	//request-site
	router.POST("/api/v1/site/request/create", r.request.SiteCreateRequest)
	router.POST("/api/v1/site/request/create-file", r.request.SiteRequestFileCreate)

	//contact-site
	router.GET("/api/v1/site/contact/list", r.contact.SiteGetContactList)
	router.GET("/api/v1/site/contact/:key", r.contact.SiteGetContactDetail)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
