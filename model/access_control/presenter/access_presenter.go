package presenter

import (
	"github.com/labstack/echo/v4"
	"go-mongo/model/access_control/model"
	"go-mongo/model/access_control/usecase"
	"go-mongo/utils"
	"net/http"
	"strconv"
)

type AccessHandler struct {
	AccessUseCase usecase.UseCase
}

func NewAccessControlHandler(e *echo.Echo, useCase usecase.UseCase) {
	injection := &AccessHandler{
		AccessUseCase: useCase,
	}
	group := e.Group("/api/v1")
	group.GET("/access-controls", injection.GetAllAccessControl)
	group.GET("/access-control/:id", injection.GetDetailAccessControl)
	group.GET("/access-control", injection.CreateNewAccessControl)
}

func (h *AccessHandler) GetAllAccessControl(ctx echo.Context) error {
	var (
		limit        = ctx.Param("limit")
		page         = ctx.Param("page")
		convLimit, _ = strconv.ParseInt(limit, 10, 64)
		convPage, _  = strconv.ParseInt(page, 10, 64)
	)

	count, err := h.AccessUseCase.CountAllUsers()
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	acls, err := h.AccessUseCase.FindAllUsers(convLimit, convPage)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(acls),
		"total": count,
		"data":  acls,
		"page":  convPage,
		"limit": convLimit,
	})
}

func (h *AccessHandler) GetDetailAccessControl(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusOK, echo.Map{
			"message": "Parameter is required.",
		})
	}

	acl, err := h.AccessUseCase.FindAccessControlById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": acl,
	})
}

func (h *AccessHandler) CreateNewAccessControl(ctx echo.Context) error {
	payload := new(model.CreateAccessControl)

	err := ctx.Bind(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	_, err = h.AccessUseCase.CreateNewAccessControl(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Access Control created successfully.",
		"data":    payload,
	})
}
