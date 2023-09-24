package main

import (
	"fmt"
	"log"
	auth2 "shifo-backend-website/internal/auth"
	"shifo-backend-website/internal/controller/http/v1/auth"
	contact_controller "shifo-backend-website/internal/controller/http/v1/contact"
	faq_controller "shifo-backend-website/internal/controller/http/v1/faq"
	menu_controller "shifo-backend-website/internal/controller/http/v1/menu"
	opportunity_controller "shifo-backend-website/internal/controller/http/v1/opportunity"
	post_controller "shifo-backend-website/internal/controller/http/v1/post"
	request_controller "shifo-backend-website/internal/controller/http/v1/request"
	user_controller "shifo-backend-website/internal/controller/http/v1/user"
	"shifo-backend-website/internal/pkg/config"
	"shifo-backend-website/internal/pkg/repository/postgres"
	"shifo-backend-website/internal/pkg/script"
	contact_repo "shifo-backend-website/internal/repository/postgres/contact"
	faq_repo "shifo-backend-website/internal/repository/postgres/faq"
	menu_repo "shifo-backend-website/internal/repository/postgres/menu"
	opportunity_repo "shifo-backend-website/internal/repository/postgres/opportunity"
	opportunity_file_repo "shifo-backend-website/internal/repository/postgres/opportunity_file"
	post_repo "shifo-backend-website/internal/repository/postgres/post"
	post_file_repo "shifo-backend-website/internal/repository/postgres/post_file"
	request_repo "shifo-backend-website/internal/repository/postgres/request"
	request_file_repo "shifo-backend-website/internal/repository/postgres/request_file"
	user_repo "shifo-backend-website/internal/repository/postgres/user"

	"shifo-backend-website/internal/router"
)

func main() {
	// config
	cfg := config.GetConf()

	// databases
	postgresDB := postgres.New(cfg.DBUsername, cfg.DBPassword, cfg.DBPort, cfg.DBName, config.GetConf().DefaultLang, config.GetConf().BaseUrl)

	//migration
	script.MigrateUP(postgresDB)

	// authenticator
	authenticator := auth2.New(postgresDB)

	//repository
	userRepo := user_repo.NewRepository(postgresDB)
	postRepo := post_repo.NewRepository(postgresDB)
	postFileRepo := post_file_repo.NewRepository(postgresDB)
	faqRepo := faq_repo.NewRepository(postgresDB)
	opportunityRepo := opportunity_repo.NewRepository(postgresDB)
	opportunityFileRepo := opportunity_file_repo.NewRepository(postgresDB)
	menuRepo := menu_repo.NewRepository(postgresDB)
	requestRepo := request_repo.NewRepository(postgresDB)
	requestFileRepo := request_file_repo.NewRepository(postgresDB)
	contactRepo := contact_repo.NewRepository(postgresDB)

	//controller
	userController := user_controller.NewController(userRepo, authenticator)
	postController := post_controller.NewController(postRepo, postFileRepo)
	authController := auth.NewController(userRepo, authenticator)
	faqController := faq_controller.NewController(faqRepo)
	opportunityController := opportunity_controller.NewController(opportunityRepo, opportunityFileRepo)
	menuController := menu_controller.NewController(menuRepo)
	requestController := request_controller.NewController(requestRepo, requestFileRepo)
	contactController := contact_controller.NewController(contactRepo)

	// router
	r := router.New(authenticator, userController, authController, postController, faqController, opportunityController, menuController, requestController, contactController)
	log.Fatalln(r.Init(fmt.Sprintf(":%s", cfg.Port)))

}
