package handler

import (
	"net/http"

	"github.com/andy-smoker/wh-server/pkg/structs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateItem(c *gin.Context) {
	var input structs.WHitem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	// добаляем новго пользователя в БД
	id, err := h.services.Warehouse.CreateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// отправляем статус ОК(200) с id созданного пользователя
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
