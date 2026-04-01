package rest_handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *UserHandler) GetCounter(c *echo.Context) error {
	counter, err := h.uc.GetCounter(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, domainCounterToRestAPI(counter))
}
