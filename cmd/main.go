package main

import (
	"fmt"
	"log"
	auth2 "shifo-backend-website/internal/auth"
	"shifo-backend-website/internal/controller/http/v1/auth"

	doctor_controller "shifo-backend-website/internal/controller/http/v1/doctor"
	specialty_controller "shifo-backend-website/internal/controller/http/v1/specialty"
	workplace_controller "shifo-backend-website/internal/controller/http/v1/workplace"
	// opportunity_controller "shifo-backend-website/internal/controller/http/v1/opportunity"
	// post_controller "shifo-backend-website/internal/controller/http/v1/post"
	// request_controller "shifo-backend-website/internal/controller/http/v1/request"
	user_controller "shifo-backend-website/internal/controller/http/v1/user"
	"shifo-backend-website/internal/pkg/config"
	"shifo-backend-website/internal/pkg/repository/postgres"
	"shifo-backend-website/internal/pkg/script"
	doctor_repo "shifo-backend-website/internal/repository/postgres/doctor"
	specialty_repo "shifo-backend-website/internal/repository/postgres/specialty"
	workplace_repo "shifo-backend-website/internal/repository/postgres/workplace"
	// opportunity_repo "shifo-backend-website/internal/repository/postgres/opportunity"
	// opportunity_file_repo "shifo-backend-website/internal/repository/postgres/opportunity_file"
	// post_repo "shifo-backend-website/internal/repository/postgres/post"
	// post_file_repo "shifo-backend-website/internal/repository/postgres/post_file"
	// request_repo "shifo-backend-website/internal/repository/postgres/request"
	// request_file_repo "shifo-backend-website/internal/repository/postgres/request_file"
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
	doctorRepo := doctor_repo.NewRepository(postgresDB)
	specialtyRepo := specialty_repo.NewRepository(postgresDB)
	workplaceRepo := workplace_repo.NewRepository(postgresDB)
	// opportunityRepo := opportunity_repo.NewRepository(postgresDB)
	// opportunityFileRepo := opportunity_file_repo.NewRepository(postgresDB)
	// menuRepo := menu_repo.NewRepository(postgresDB)
	// requestRepo := request_repo.NewRepository(postgresDB)
	// requestFileRepo := request_file_repo.NewRepository(postgresDB)
	// contactRepo := contact_repo.NewRepository(postgresDB)

	//controller
	userController := user_controller.NewController(userRepo, authenticator)
	doctorController := doctor_controller.NewController(doctorRepo)
	authController := auth.NewController(userRepo, authenticator)
	specialtyController := specialty_controller.NewController(specialtyRepo)
	workplaceController := workplace_controller.NewController(workplaceRepo)
	// menuController := menu_controller.NewController(menuRepo)
	// requestController := request_controller.NewController(requestRepo, requestFileRepo)
	// contactController := contact_controller.NewController(contactRepo)

	// router
	r := router.New(authenticator, userController, authController, doctorController, specialtyController, workplaceController)
	log.Fatalln(r.Init(fmt.Sprintf(":%s", cfg.Port)))

}
