package handler

import (
	"net/http"
	"strconv"
	_ "user/api/docs"
	"user/api/models"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user [POST]
// @Summary		create an user
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

// CreateUsers godoc
// @Security ApiKeyAuth
// @Router		/api/v1/users [POST]
// @Summary		create users
// @Description	This api create users and returns error or null
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		user body models.AddUsers true "user"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateMany(c *gin.Context) {
	users := models.AddUsers{}

	if err := c.ShouldBindJSON(&users); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.User().CreateMany(c.Request.Context(), users)
	if err != nil {
		handleResponse(c, h.Log, "error while creating users", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, "Created successfully")
}

// Update godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user [PUT]
// @Summary		update an user
// @Description	This api update an user
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		user body models.UpdateUser true "user"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) Update(c *gin.Context) {
	user := models.UpdateUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.User().Update(c.Request.Context(), user)

	if err != nil {
		handleResponse(c, h.Log, "error while updating an user", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// UpdateMany godoc
// @Security ApiKeyAuth
// @Router		/api/v1/users [PUT]
// @Summary		update users
// @Description	This api update users
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		user body models.UpdateUsers true "user"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateMany(c *gin.Context) {
	users := models.UpdateUsers{}

	if err := c.ShouldBindJSON(&users); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.User().UpdateMany(c.Request.Context(), users)

	if err != nil {
		handleResponse(c, h.Log, "error while updating users", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, "Updated successfully")
}

// Delete godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user/ [DELETE]
// @Summary		delete an user
// @Description	This api delete a user
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		user body models.DeleteUser true "user"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) Delete(c *gin.Context) {
	user := models.DeleteUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.User().Delete(c.Request.Context(), user)

	if err != nil {
		handleResponse(c, h.Log, "error while deleting an user", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, "Deleted successfully")
}

// DeleteMany godoc
// @Security ApiKeyAuth
// @Router		/api/v1/users [DELETE]
// @Summary		delete users
// @Description	This api delete users and returns error or null
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		user body models.DeleteUsers true "user"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteMany(c *gin.Context) {
	users := models.DeleteUsers{}

	if err := c.ShouldBindJSON(&users); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.User().DeleteMany(c.Request.Context(), users)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting users", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, "Deleted successfully")
}

// GetById godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user/{id} [GET]
// @Summary		get an user
// @Description	This api get an user
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		id path int true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetById(c *gin.Context) {

	id := c.Param("id")

	userId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		handleResponse(c, h.Log, "error while converting userId", http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Service.User().GetById(c.Request.Context(), userId)
	if err != nil {
		handleResponse(c, h.Log, "error while getting an user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Got successfully", http.StatusOK, user)
}

// GetList godoc
// @Security ApiKeyAuth
// @Router		/api/v1/users [GET]
// @Summary		get  users
// @Description	This api get all avtive users
// @Tags		user
// @Accept		json
// @Produce		json
// @Param   	sort query string false "sort"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetList(c *gin.Context) {
	sort := c.Query("sort")
	sortClauses := ParseSortParam(sort)

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Service.User().GetList(c.Request.Context(), models.GetListRequest{
		Sort:  sortClauses,
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all active users", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "Got successful", http.StatusOK, resp)
}
