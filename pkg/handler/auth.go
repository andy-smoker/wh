package handler

import (
	"net/http"

	"github.com/andy-smoker/wh-server/pkg/structs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input structs.User
	// проверяем валидность введённых данных
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	// добаляем новго пользователя в БД
	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// отправляем статус ОК(200) с id созданного пользователя
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"login" binding:"required"`
	Password string `json:"pass" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	// проверяем валидность запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// формируем токен пользователя
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// отправляем статус OK(200) с токеном в теле
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
