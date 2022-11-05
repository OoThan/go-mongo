package presenter

import (
	"github.com/labstack/echo/v4"
	"go-mongo/model/role/model"
	"go-mongo/model/role/usecase"
	"go-mongo/utils"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	roleUseCase usecase.UseCase
}

func NewRoleHandler(e *echo.Echo, useCase usecase.UseCase) {
	injection := &RoleHandler{
		roleUseCase: useCase,
	}

	group := e.Group("/api/v1")
	group.GET("/roles", injection.GetAllRoles)
	group.POST("/role", injection.CreateNewRole)
	group.GET("/role/:id", injection.GetDetailRole)
	group.PUT("/role/:id", injection.UpdateRole)
	group.DELETE("/role/:id", injection.DeleteRole)
}

func (h *RoleHandler) GetAllRoles(ctx echo.Context) error {
	var (
		limit        = ctx.Param("limit")
		page         = ctx.Param("page")
		convLimit, _ = strconv.ParseInt(limit, 10, 64)
		convPage, _  = strconv.ParseInt(page, 10, 64)
	)
	roles, err := h.roleUseCase.FindAllRoles(convLimit, convPage)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}
	total, err := h.roleUseCase.CountAllRoles()
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(roles),
		"total": total,
		"data":  roles,
		"limit": convLimit,
		"page":  convPage,
	})
}

func (h *RoleHandler) CreateNewRole(ctx echo.Context) error {
	payload := new(model.CreateRole)
	err := ctx.Bind(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}
	err = h.roleUseCase.CreateNewRole(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"data":    payload,
		"message": "Role created successfully.",
	})
}

func (h *RoleHandler) GetDetailRole(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": "Parameter id is required.",
		})
	}
	role, err := h.roleUseCase.FindRoleById(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": role,
	})
}

func (h *RoleHandler) UpdateRole(ctx echo.Context) error {
	id := ctx.Param("id")
	payload := new(model.UpdateRole)
	err := ctx.Bind(payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	if id == "" {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": "Parameter id is required.",
		})
	}
	err = h.roleUseCase.UpdateRole(id, payload)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Updated successfully.",
		"data":    payload,
	})
}

func (h *RoleHandler) DeleteRole(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": "Parameter id is required.",
		})
	}

	err := h.roleUseCase.Delete(id)
	if !utils.GlobalErrorDatabaseException(err) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error":   err.Error(),
			"message": "use case error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Deleted successfully.",
	})
}
