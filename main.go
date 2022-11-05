package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-mongo/config"
	"go-mongo/logger"

	UserHandlerPackage "go-mongo/model/user/presenter"
	UserRepoPackage "go-mongo/model/user/repository"
	UserUseCasePackage "go-mongo/model/user/usecase"

	RoleHandlerPackage "go-mongo/model/role/presenter"
	RoleRepoPackage "go-mongo/model/role/repository"
	RoleUseCasePackage "go-mongo/model/role/usecase"

	AccessControlHandlerPackage "go-mongo/model/access_control/presenter"
	AccessControlRepoPackage "go-mongo/model/access_control/repository"
	AccessControlUseCasePackage "go-mongo/model/access_control/usecase"
	"net/http"
)

func main() {
	mongoConn, err := config.MongoConnection()
	if err != nil {
		logger.Sugar.Error(err)
	}

	echoRouter := echo.New()
	echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// User modules
	userRepo := UserRepoPackage.NewUserRepository(mongoConn)
	userUseCase := UserUseCasePackage.NewUserUseCase(userRepo)
	UserHandlerPackage.NewUserHandler(echoRouter, userUseCase)

	// Role modules
	roleRepo := RoleRepoPackage.NewRoleRepository(mongoConn)
	roleUseCase := RoleUseCasePackage.NewRoleUseCase(roleRepo)
	RoleHandlerPackage.NewRoleHandler(echoRouter, roleUseCase)

	accessControlRepo := AccessControlRepoPackage.NewAccessControlRepository(mongoConn)
	accessControlUsecase := AccessControlUseCasePackage.NewAccessControlUseCase(accessControlRepo)
	AccessControlHandlerPackage.NewAccessControlHandler(echoRouter, accessControlUsecase)

	//Configuration of logger
	echoRouter.Use(middleware.Logger())
	echoRouter.Logger.Fatal(echoRouter.Start(":8081"))
}
