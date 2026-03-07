package rest_handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/zlietapki/microboiler_rest_server/internal/rest_models"
)

func (h *UserHandler) CreateUser(c *echo.Context) error {
	var req rest_models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := h.uc.Create(req.Name, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, domainUserToRestAPI(user))
}
