package rest_handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *UserHandler) GetUsers(c *echo.Context) error {
	users, err := h.uc.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, domainUsersToRestAPI(users))
}
