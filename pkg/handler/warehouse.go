package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}
	item, err := h.services.Warehouse.GetItem(id)
	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "no rows")
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetItemsList(c *gin.Context) {
	var input struct {
		ItemsType string `json:"items_type" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	items, err := h.services.Warehouse.GetItemsList(input.ItemsType)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "no rows")
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) UpdateItem(c *gin.Context) {
	var input structs.WHitem
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}
	input.ID = id
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	item, err := h.services.Warehouse.UpdateItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}
	var input structs.WHitem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	err = h.services.Warehouse.DeleteItem(id, input.ItemsType)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "no rows")
		return
	}
	c.JSON(http.StatusOK, true)
}
