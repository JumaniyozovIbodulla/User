package handler

import (
	"net/http"
	_ "user/api/docs"
	"user/api/models"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Security ApiKeyAuth
// @Router		/api/user [POST]
// @Summary		create a user
// @Description	This api create a user and returns its id
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		user body models.AddUser true "user"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateUser(c *gin.Context) {
	user := models.AddUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.User().Create(c.Request.Context(), user)
	if err != nil {
		handleResponse(c, h.Log, "error while creating user", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}