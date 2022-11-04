package presenter

import (
	"github.com/labstack/echo/v4"
	"go-mongo/model/user/model"
	"go-mongo/model/user/usecase"
	"go-mongo/utils"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUseCase usecase.UseCase
}

func NewUserHandler(e *echo.Echo, userUseCase usecase.UseCase) {
	injectionHandler := &UserHandler{
		userUseCase: userUseCase,
	}

	group := e.Group("/api/v1")
	group.GET("/users", injectionHandler.GetAllUsers)
	group.GET("/users/:id", injectionHandler.GetDetailUser)
	group.POST("/user", injectionHandler.CreateUser)
	group.PUT("/user/:id", injectionHandler.UpdateUser)
	group.DELETE("/user/:id", injectionHandler.DeleteUser)
}

func (uh *UserHandler) GetAllUsers(ctx echo.Context) error {
	var (
		limit           = ctx.QueryParam("limit")
		pages           = ctx.QueryParam("page")
		name            = ctx.QueryParam("name")
		convertLimit, _ = strconv.ParseInt(limit, 10, 64)
		convertPage, _  = strconv.ParseInt(pages, 10, 64)
	)

	users, err := uh.userUseCase.FindAllUsers(name, convertLimit, convertPage)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	total, err := uh.userUseCase.CountAllUsers()
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(users),
		"data":  users,
		"total": total,
		"limit": convertLimit,
		"page":  convertPage,
	})
}

func (uh *UserHandler) GetDetailUser(ctx echo.Context) error {
	id := ctx.Param("id")
	user, err := uh.userUseCase.FindUserById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

func (uh *UserHandler) CreateUser(ctx echo.Context) error {
	payload := new(model.CreateUser)
	errBind := ctx.Bind(payload)
	if !utils.GlobalErrorDatabaseException(errBind) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errBind.Error(),
			"message": "use case error",
		})
	}

	savedUser, err := uh.userUseCase.CreateNewUser(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "User created successfully",
		"data":    savedUser,
	})
}

func (uh *UserHandler) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	userUpdate := new(model.UpdateUser)
	errBind := ctx.Bind(userUpdate)
	if !utils.GlobalErrorDatabaseException(errBind) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errBind.Error(),
			"message": "use case error",
		})
	}

	errUpdate := uh.userUseCase.UpdateUser(id, userUpdate)
	if !utils.GlobalErrorDatabaseException(errUpdate) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errUpdate.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "user updated successfully",
	})
}

func (uh *UserHandler) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")
	errDel := uh.userUseCase.DeleteUser(id)
	if !utils.GlobalErrorDatabaseException(errDel) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errDel.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "user deleted successfully",
	})
}
